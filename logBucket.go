package main

import (
	"errors"
	"regexp"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/viper"
)

type LogBucket struct {
	service *s3.S3
	bucket  string
	region  string
}

func NewLogBucket(svc *s3.S3) *LogBucket {
	l := LogBucket{
		service: svc,
		bucket:  viper.GetString("bucket"),
		region:  viper.GetString("region"),
	}
	return &l
}

func (l *LogBucket) GetFirstPropertyName() (string, error) {
	logFiles, _ := regexp.Compile(`([a-zA-Z0-9]+)\.([0-9]{4}-[0-9]{2}-[0-9]{2}-[0-9]{2})\.([a-z0-9]+)\.gz`)
	resp, _ := l.service.ListObjects(&s3.ListObjectsInput{Bucket: &l.bucket})
	if len(resp.Contents) == 0 {
		return "", errors.New("No files in the bucket")
	}
	for _, key := range resp.Contents {
		lfs := logFiles.FindStringSubmatch(*key.Key)
		if len(lfs) > 0 {
			return lfs[1], nil
		}
	}
	return "", errors.New("No log files found in the bucket")
}
