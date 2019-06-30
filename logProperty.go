package main

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/viper"
)

type LogProperty struct {
	service  *s3.S3
	bucket   string
	region   string
	property string
}

func NewLogProperty(svc *s3.S3, property string) *LogProperty {
	l := LogProperty{
		service:  svc,
		bucket:   viper.GetString("bucket"),
		region:   viper.GetString("region"),
		property: property,
	}
	return &l
}

// TODO: Make it AWS S3 API limit safe
func (l *LogProperty) GetFirstTimeSlotName() (string, error) {
	logFiles, _ := regexp.Compile(`([a-zA-Z0-9]+)\.([0-9]{4}-[0-9]{2}-[0-9]{2}-[0-9]{2})\.([a-z0-9]+)\.gz`)
	resp, _ := l.service.ListObjects(&s3.ListObjectsInput{
		Bucket: &l.bucket,
		Prefix: aws.String(fmt.Sprintf("%s.", l.property)),
	})
	if len(resp.Contents) == 0 {
		return "", errors.New("No files in the bucket")
	}
	for _, key := range resp.Contents {
		lfs := logFiles.FindStringSubmatch(*key.Key)
		if len(lfs) > 0 {
			return lfs[2], nil
		}
	}
	return "", errors.New("No log files found in the bucket")
}

// TODO: Make it AWS S3 API limit safe
func (l *LogProperty) ListFilesInTimeSlot(ts string) ([]string, error) {
	logFiles, _ := regexp.Compile(`([a-zA-Z0-9]+)\.([0-9]{4}-[0-9]{2}-[0-9]{2}-[0-9]{2})\.([a-z0-9]+)\.gz`)
	resp, _ := l.service.ListObjects(&s3.ListObjectsInput{
		Bucket: &l.bucket,
		Prefix: aws.String(fmt.Sprintf("%s.%s.", l.property, ts)),
	})
	if len(resp.Contents) == 0 {
		return nil, errors.New("No log files matching the requested parameters in the bucket")
	}
	var files []string
	for _, key := range resp.Contents {
		if logFiles.MatchString(*key.Key) {
			files = append(files, *key.Key)
		}
	}
	return files, nil
}

// TODO: Make it AWS S3 API limit safe
func (l *LogProperty) ListFilesInDay(day string) ([]string, error) {
	logFiles, _ := regexp.Compile(`([a-zA-Z0-9]+)\.([0-9]{4}-[0-9]{2}-[0-9]{2}-[0-9]{2})\.([a-z0-9]+)\.gz`)
	resp, _ := l.service.ListObjects(&s3.ListObjectsInput{
		Bucket: &l.bucket,
		Prefix: aws.String(fmt.Sprintf("%s.%s-", l.property, day)),
	})
	if len(resp.Contents) == 0 {
		return nil, errors.New("No log files matching the requested parameters in the bucket")
	}
	var files []string
	for _, key := range resp.Contents {
		if logFiles.MatchString(*key.Key) {
			files = append(files, *key.Key)
		}
	}
	return files, nil
}
