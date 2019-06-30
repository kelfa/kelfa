package main

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/spf13/viper"
)

type BucketManager struct {
	session *session.Session
	bucket  string
	region  string
	tempDir string
}

func NewBucketManager(session *session.Session) (*BucketManager, error) {
	dir, err := ioutil.TempDir(os.TempDir(), "kelfa")
	if err != nil {
		return nil, err
	}
	l := BucketManager{
		session: session,
		bucket:  viper.GetString("bucket"),
		region:  viper.GetString("region"),
		tempDir: dir,
	}
	return &l, nil
}

func (l *BucketManager) ReadLogFile(filename string) ([]byte, error) {
	downloader := s3manager.NewDownloader(l.session)

	file, err := os.Create(path.Join(l.tempDir, filename))
	if err != nil {
		return nil, err
	}
	defer os.Remove(file.Name())

	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: &l.bucket,
			Key:    &filename,
		})

	if err != nil {
		return nil, fmt.Errorf("Unable to download item %q, %v", filename, err)
	}

	_, err = file.Seek(0, 0)

	gzr, err := gzip.NewReader(file)
	if err != nil {
		return nil, err
	}
	defer gzr.Close()

	s, err := ioutil.ReadAll(gzr)
	if err != nil {
		return nil, err
	}
	return s, nil
}
