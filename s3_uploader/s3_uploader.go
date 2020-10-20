package s3_uploader

import (
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"../config"
)

func ConnectAws() *session.Session {
	conf := config.New().AWS

	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(conf.Region),
			Credentials: credentials.NewStaticCredentials(
				conf.AccessKeyID,
				conf.SecretAccessKey,
				"", // a token will be created when the session it's used.
			),
		})

	if err != nil {
		panic(err)
	}

	return sess
}

// UploadFile takes a key and file and uploads to s3 - this is an unoptimised implememtation
// that will upload the entire file, we could update this to upload in chunks
func UploadFile(key string, file io.Reader) (*s3manager.UploadOutput, error) {
	sess := ConnectAws()
	conf := config.New().AWS
	uploader := s3manager.NewUploader(sess)

	//upload to the s3 bucket
	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(conf.BucketName),
		ACL:    aws.String("private"),
		Key:    aws.String(key),
		Body:   file,
	})

	return up, err
}

// DownloadFile takes a key and downloads the file from s3 - this is an unoptimised implememtation
// that will dowload the entire file, we could update this to dowload in chunks
func DownloadFile(filename string) (*os.File, error) {
	sess := ConnectAws()
	conf := config.New().AWS
	downloader := s3manager.NewDownloader(sess)

	f, err := os.Create(fmt.Sprintf("temp/%s", filename))

	if err != nil {
		return nil, fmt.Errorf("Failed to create file %s, %v", filename, err)
	}

	// Write the contents of S3 Object to the file
	_, err = downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(conf.BucketName),
		Key:    aws.String(filename),
	})

	return f, err
}

func buildFilePath(bucket string, region string, filename string) string {
	return fmt.Sprintf("https://%s.s3-%s.amazonaws.com/%s", bucket, region, filename)
}
