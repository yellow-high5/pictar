package server

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadS3(targetFilePath string, saveName string) error {
	// TODO: Read config file

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile:           "",                         // Write config file
		SharedConfigState: session.SharedConfigEnable, // Write config file
	}))

	file, err := os.Open(targetFilePath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	bucketName := "" // Write config file
	objectKey := saveName

	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
