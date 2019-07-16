package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/spf13/viper"

	"go.kelfa.io/aws-cloudfront-logCompactor/pkg/timeSafety"
)

func main() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("aws-cloudfront-logCompactor")
	viper.AddConfigPath("$HOME/.kelfa")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}
	// you need to store your AWS credentials in ~/.aws/credentials

	region := viper.GetString("region")
	sess, _ := session.NewSession(&aws.Config{Region: &region})
	bm, err := NewBucketManager(sess)
	if err != nil {
		log.Fatalf("unable to create a bucket manager instance: %v", err)
	}

	property, err := bm.GetFirstPropertyName()
	if err != nil {
		log.Fatalf("unable to identify the property to manage: %v", err)
	}
	timeSlot, err := bm.GetFirstTimeSlotName(property)
	if err != nil {
		log.Fatalf("unable to identify the first time slot for the %s property to manage: %v", property, err)
	}
	safe, err := timeSafety.IsTimeSlotDaySafe(timeSlot)
	if err != nil {
		log.Fatalf("unable to determine if the %s timeslot is safe: %v", timeSlot, err)
	}
	if !safe {
		fmt.Printf("the %s timeslot is not safe to be processed\n", timeSlot)
		return
	}
	lfs, err := bm.ListFilesInDay(property, timeSlot[:len(timeSlot)-3])
	if err != nil {
		log.Fatalf("impossible to read the list of files in the %s timeslot: %v", timeSlot, err)
	}
	err = bm.MergeFiles(lfs, viper.GetString("archive_folder")+"/"+timeSlot[:len(timeSlot)-3]+".log")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("file %s created\n", viper.GetString("archive_folder")+"/"+timeSlot[:len(timeSlot)-3]+".log")
	for _, lf := range lfs {
		err := bm.MoveFile(lf, viper.GetString("compacted_folder")+"/"+lf)
		if err != nil {
			log.Fatalf("unable to move the %s file: %v", lf, err)
		}
		fmt.Printf("file %s moved to %s folder\n", lf, viper.GetString("compacted_folder"))
	}
}
