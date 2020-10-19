package s3_uploader

import (
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
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

	fmt.Printf("%v", sess)
	return sess
}

func UploadFile(filename string, file io.Reader) (*s3manager.UploadOutput, string, error) {
	sess := ConnectAws()
	conf := config.New().AWS
	uploader := s3manager.NewUploader(sess)

	//upload to the s3 bucket
	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(conf.BucketName),
		ACL:    aws.String("private"),
		Key:    aws.String(filename),
		Body:   file,
	})

	filePath := buildFilePath(conf.BucketName, conf.Region, filename)

	return up, filePath, err
}

func buildFilePath(bucket string, region string, filename string) string {
	return fmt.Sprintf("https://%s.s3-%s.amazonaws.com/%s", bucket, region, filename)
}
