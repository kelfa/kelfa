package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("aws-cloudfront-logCompactor")
	viper.AddConfigPath("$HOME/.kelfa")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// you need to store your AWS credentials in ~/.aws/credentials

	region := viper.GetString("region")
	sess, _ := session.NewSession(&aws.Config{Region: &region})
	bm, err := NewBucketManager(sess)

	property, err := bm.GetFirstPropertyName()
	if err != nil {
		log.Fatalf("Unable to identify the property to manage: %v", err)
	}
	timeSlot, err := bm.GetFirstTimeSlotName(property)
	if err != nil {
		log.Fatalf("Unable to identify the first time slot for the %s property to manage: %v", property, err)
	}

	lfs, err := bm.ListFilesInDay(property, timeSlot[:len(timeSlot)-3])
	if err != nil {
		log.Fatalf("Impossible to read the list of files in the %s timeslot: %v", timeSlot, err)
	}
	err = bm.MergeFiles(lfs, viper.GetString("archive_folder")+"/"+timeSlot[:len(timeSlot)-3]+".log")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File %s created\n", viper.GetString("archive_folder")+"/"+timeSlot[:len(timeSlot)-3]+".log")
	for _, lf := range lfs {
		err := bm.MoveFile(lf, viper.GetString("compacted_folder")+"/"+lf)
		if err != nil {
			log.Fatalf("unable to move the %s file: %v", lf, err)
		}
		fmt.Printf("File %s moved to %s folder\n", lf, viper.GetString("compacted_folder"))
	}
}
