package s3_uploader

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"

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
