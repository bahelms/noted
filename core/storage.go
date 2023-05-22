package core

import (
	"fmt"
	"os"
	"path"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func commitFile(filePath string) {
	bucketName := "noted-file-storage"

	sess, _ := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region: aws.String("us-east-1"),
		},
	})
	s3Client := s3.New(sess)

	buckets, _ := s3Client.ListBuckets(nil)
	if !contains(buckets.Buckets, bucketName) {
		_, err := s3Client.CreateBucket(&s3.CreateBucketInput{
			Bucket: aws.String(bucketName),
		})
		if err != nil {
			fmt.Println(err)
			panic("Error creating bucket")
		}
	}

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, filename := path.Split(filePath)
	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
		Body:   file,
	})

	if err != nil {
		panic(err)
	}
}

func contains(s []*s3.Bucket, str string) bool {
	for _, bucket := range s {
		if *bucket.Name == str {
			return true
		}
	}

	return false
}
