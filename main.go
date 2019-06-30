package main

import (
	"bufio"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/viper"

	"fmt"
	"log"
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
	svc := s3.New(sess)
	l := NewLogBucket(svc)
	property, err := l.GetFirstPropertyName()
	if err != nil {
		log.Fatalf("Unable to identify the property to manage: %v", err)
	}
	fmt.Printf(property)
	p := NewLogProperty(svc, property)
	timeSlot, err := p.GetFirstTimeSlotName()
	if err != nil {
		log.Fatalf("Unable to identify the first time slot for the %s property to manage: %v", property, err)
	}
	fmt.Printf(timeSlot)

	bm, err := NewBucketManager(sess)
	w := bufio.NewWriter(os.Stdout)
	lfs, err := p.ListFilesInTimeSlot(timeSlot)
	if err != nil {
		log.Fatalf("Impossible to read the list of files in the %s timeslot: %v", timeSlot, err)
	}
	for _, lf := range lfs {
		content, err := bm.ReadLogFile(lf)
		if err != nil {
			log.Fatalf("unable to read the content of the %s file: %v", lf, err)
		}
		w.Write(content)
	}
	w.Flush()
}
