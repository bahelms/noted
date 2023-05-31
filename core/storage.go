package core

import (
	"fmt"
	"os"
	"path"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/bahelms/noted/config"
)

func commitFile(filePath string, config config.Config) {
	sess := awsSession(config.AwsProfile)
	client := s3Client(sess)

	ensureBucketExists(client, config.S3BucketName)

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, filename := path.Split(filePath)
	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(config.S3BucketName),
		Key:    aws.String(filename),
		Body:   file,
	})

	if err != nil {
		panic(err)
	}
}

func ensureBucketExists(client *s3.S3, bucketName string) {
	buckets, _ := client.ListBuckets(nil)
	if !contains(buckets.Buckets, bucketName) {
		_, err := client.CreateBucket(&s3.CreateBucketInput{
			Bucket: aws.String(bucketName),
		})
		if err != nil {
			fmt.Println(err)
			panic("Error creating bucket")
		}
	}
}

func awsSession(profile string) *session.Session {
	sess, _ := session.NewSessionWithOptions(session.Options{
		Profile: profile,
		Config: aws.Config{
			Region: aws.String("us-east-1"),
		},
	})
	return sess
}

func s3Client(session *session.Session) *s3.S3 {
	return s3.New(session)
}

func objectsInBucket(client *s3.S3, cfg config.Config) (*s3.ListObjectsV2Output, error) {
	objects, err := client.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(cfg.S3BucketName),
	})
	if err != nil {
		fmt.Printf("Couldn't retrieve bucket items: %v", err)
		return nil, err
	}
	return objects, nil
}

func downloadAllFiles(cfg config.Config) {
	sess := awsSession(cfg.AwsProfile)
	client := s3Client(sess)
	objects, _ := objectsInBucket(client, cfg)

	for _, object := range objects.Contents {
		filepath := cfg.LocalFilePath(*object.Key)
		downloadFile(filepath, *object.Key, cfg.S3BucketName, sess)
	}
}

func downloadFile(filepath string, objectKey string, bucketName string, sess *session.Session) {
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Printf("Error creating file for download: %v", err)
		return
	}
	defer file.Close()

	downloader := s3manager.NewDownloader(sess)
	_, err = downloader.Download(
		file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(objectKey),
		},
	)
	if err != nil {
		fmt.Printf("Error downloading file: %v", err)
	}
}

func deleteExternalFile(cfg config.Config, filename string) {
	sess := awsSession(cfg.AwsProfile)
	client := s3Client(sess)
	_, err := client.DeleteObject(
		&s3.DeleteObjectInput{
			Bucket: aws.String(cfg.S3BucketName),
			Key:    aws.String(filename),
		},
	)
	if err != nil {
		fmt.Printf("Error deleting remote file: %v", err)
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
