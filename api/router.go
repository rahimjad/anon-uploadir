package api

import (
	"fmt"
	"net/http"

	"../s3_uploader"
	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")

	up, filepath, err := s3_uploader.UploadFile(header.Filename, file)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":    fmt.Sprintf("%v", err),
			"uploader": up,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"filepath": filepath})
}

func Run() {
	router := gin.Default()

	router.POST("/file", UploadFile)

	router.Run(":8080")
}
