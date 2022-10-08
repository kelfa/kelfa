package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/spf13/viper"

	"github.com/kelfa/kelfa/pkg/timeSafety"
)

func main() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("aws-cloudfront-logCompactor")
	viper.AddConfigPath("$HOME/.kelfa")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("fatal error config file: %s\n", err)
	}
	// you need to store your AWS credentials in ~/.aws/credentials

	region := viper.GetString("region")
	sess, _ := session.NewSession(&aws.Config{Region: &region})
	bm, err := NewBucketManager(sess)
	if err != nil {
		log.Fatalf("unable to create a bucket manager instance: %v\n", err)
	}

	property, err := bm.GetFirstPropertyName()
	if err != nil {
		log.Fatalf("unable to identify the property to manage: %v\n", err)
	}
	timeSlot, err := bm.GetFirstTimeSlotName(property)
	if err != nil {
		log.Fatalf("unable to identify the first time slot for the %s property to manage: %v\n", property, err)
	}
	safe, err := timeSafety.IsTimeSlotDaySafe(timeSlot)
	if err != nil {
		log.Fatalf("unable to determine if the %s timeslot is safe: %v\n", timeSlot, err)
	}
	if !safe {
		fmt.Printf("the %s timeslot is not safe to be processed\n", timeSlot)
		return
	}
	lfs, err := bm.ListFilesInDay(property, timeSlot[:len(timeSlot)-3])
	if err != nil {
		log.Fatalf("impossible to read the list of files in the %s timeslot: %v\n", timeSlot, err)
	}
	err = bm.MergeFiles(lfs, viper.GetString("archive_folder")+"/"+timeSlot[:len(timeSlot)-3]+".log")
	if err != nil {
		log.Fatalf("impossible to merge files: %v\n", err)
	}
	fmt.Printf("file %s created\n", viper.GetString("archive_folder")+"/"+timeSlot[:len(timeSlot)-3]+".log")
	for _, lf := range lfs {
		err := bm.MoveFile(lf, viper.GetString("compacted_folder")+"/"+lf)
		if err != nil {
			log.Fatalf("unable to move the %s file: %v\n", lf, err)
		}
		fmt.Printf("file %s moved to %s folder\n", lf, viper.GetString("compacted_folder"))
	}
}
