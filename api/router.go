package api

import (
	"fmt"
	"net/http"

	"../config"
	"../s3_uploader"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	conf := config.New().AWS
	bucket := conf.BucketName
	sess := c.MustGet("sess").(*session.Session)
	uploader := s3manager.NewUploader(sess)
	file, header, err := c.Request.FormFile("photo")
	filename := header.Filename

	fmt.Printf("%v", conf)
	//upload to the s3 bucket
	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		ACL:    aws.String("public-read"),
		Key:    aws.String(filename),
		Body:   file,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":    err,
			"uploader": up,
		})
		return
	}

	filepath := "https://" + bucket + "." + "s3-" + conf.Region + ".amazonaws.com/" + filename

	c.JSON(http.StatusOK, gin.H{"filepath": filepath})
}

func Run() {
	sess := s3_uploader.ConnectAws()
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set("sess", sess)
		c.Next()
	})

	router.POST("/file", UploadFile)

	router.Run(":8080")
}
