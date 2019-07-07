package main

import (
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"sort"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/spf13/viper"

	"go.kelfa.io/elf"
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

// TODO: Make it AWS S3 API limit safe
func (b *BucketManager) GetFirstPropertyName() (string, error) {
	svc := s3.New(b.session)
	logFiles, _ := regexp.Compile(`([a-zA-Z0-9]+)\.([0-9]{4}-[0-9]{2}-[0-9]{2}-[0-9]{2})\.([a-z0-9]+)\.gz`)
	resp, _ := svc.ListObjects(&s3.ListObjectsInput{Bucket: &b.bucket})
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

// TODO: Make it AWS S3 API limit safe
func (b *BucketManager) GetFirstTimeSlotName(property string) (string, error) {
	svc := s3.New(b.session)
	logFiles, _ := regexp.Compile(`([a-zA-Z0-9]+)\.([0-9]{4}-[0-9]{2}-[0-9]{2}-[0-9]{2})\.([a-z0-9]+)\.gz`)
	resp, _ := svc.ListObjects(&s3.ListObjectsInput{
		Bucket: &b.bucket,
		Prefix: aws.String(fmt.Sprintf("%s.", property)),
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
func (b *BucketManager) ListFilesInTimeSlot(property string, ts string) ([]string, error) {
	svc := s3.New(b.session)
	logFiles, _ := regexp.Compile(`([a-zA-Z0-9]+)\.([0-9]{4}-[0-9]{2}-[0-9]{2}-[0-9]{2})\.([a-z0-9]+)\.gz`)
	resp, _ := svc.ListObjects(&s3.ListObjectsInput{
		Bucket: &b.bucket,
		Prefix: aws.String(fmt.Sprintf("%s.%s.", property, ts)),
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
func (b *BucketManager) ListFilesInDay(property string, day string) ([]string, error) {
	svc := s3.New(b.session)
	logFiles, _ := regexp.Compile(`([a-zA-Z0-9]+)\.([0-9]{4}-[0-9]{2}-[0-9]{2}-[0-9]{2})\.([a-z0-9]+)\.gz`)
	resp, _ := svc.ListObjects(&s3.ListObjectsInput{
		Bucket: &b.bucket,
		Prefix: aws.String(fmt.Sprintf("%s.%s-", property, day)),
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

func (b *BucketManager) ReadLogFile(filename string) ([]byte, error) {
	downloader := s3manager.NewDownloader(b.session)

	file, err := os.Create(path.Join(b.tempDir, filename))
	if err != nil {
		return nil, err
	}
	defer os.Remove(file.Name())

	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: &b.bucket,
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

func (b *BucketManager) MergeFiles(fromFiles []string, toFile string) (err error) {
	var fields []string
	var allEntries []map[string]string
	for _, lf := range fromFiles {
		content, err := b.ReadLogFile(lf)
		if err != nil {
			return fmt.Errorf("unable to read the content of the %s file: %v", lf, err)
		}
		elfReader := elf.NewReader(bytes.NewReader(content))
		entries, err := elfReader.ReadAll()
		allEntries = append(allEntries, entries...)
		if err != nil {
			return err
		}
		fields = elfReader.Fields
	}
	sort.Slice(allEntries,
		func(i, j int) bool {
			return allEntries[i]["time"] < allEntries[j]["time"]
		})
	r, w := io.Pipe()
	ew := elf.NewWriter(w, fields)
	go func() {
		err = ew.WriteAll(allEntries)
		if err != nil {
			log.Fatalf("%s", err)
		}
		ew.Flush()
		w.Close()
	}()

	// Upload input parameters
	upParams := &s3manager.UploadInput{
		Body:   aws.ReadSeekCloser(r),
		Bucket: &b.bucket,
		Key:    &toFile,
	}

	// Perform an upload.
	uploader := s3manager.NewUploader(b.session)
	_, err = uploader.Upload(upParams)
	if err != nil {
		return err
	}
	return nil
}

func (b *BucketManager) MoveFile(src string, dst string) (err error) {
	svc := s3.New(b.session)
	copyInput := &s3.CopyObjectInput{
		Bucket:     &b.bucket,
		CopySource: aws.String(fmt.Sprintf("%s/%s", b.bucket, src)),
		Key:        &dst,
	}

	_, err = svc.CopyObject(copyInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeObjectNotInActiveTierError:
				return fmt.Errorf(s3.ErrCodeObjectNotInActiveTierError, aerr.Error())
			default:
				return aerr
			}
		} else {
			return err
		}
		return err
	}

	deleteInput := &s3.DeleteObjectInput{
		Bucket: &b.bucket,
		Key:    &src,
	}
	_, err = svc.DeleteObject(deleteInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				return aerr
			}
		} else {
			return err
		}
		return
	}

	return nil
}
