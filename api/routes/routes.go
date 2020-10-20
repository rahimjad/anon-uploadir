package routes

import (
	"fmt"
	"net/http"
	"os"

	"../config"
	"../entities"
	"../s3_uploader"
	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")

	s3Metadata := entities.S3Metadata{
		FileName: header.Filename,
		FileSize: header.Size,
	}

	_, err = s3Metadata.Insert()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	up, err := s3_uploader.UploadFile(s3Metadata.ID, file)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":    fmt.Sprintf("%v", err),
			"uploader": up,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully Uploaded File",
		"fileId":  s3Metadata.ID,
	})
}

func DownloadFile(c *gin.Context) {
	id := c.Param("id")
	s3Metadata := entities.S3Metadata{}
	err := s3Metadata.QueryRow(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "File Not Found!",
		})
		return
	}

	f, err := s3_uploader.DownloadFile(s3Metadata.ID)

	// Delete the file once it has been rendered
	defer os.Remove(f.Name())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	c.File(f.Name())
}

func Run() {
	config := config.New().Router
	router := gin.Default()

	router.POST("/file", UploadFile)
	router.GET("/file/:id", DownloadFile)

	router.Run(fmt.Sprintf(":%d", config.Port))
}
