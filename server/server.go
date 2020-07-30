package server

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Boot() {
	router := gin.Default()

	router.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("image")

		if err != nil {
			c.String(400, fmt.Sprintf("Cannot get form err: %s", err.Error()))
			return
		}

		const originName = "ORIGINAL.png" // TODO: Read config file

		if err := c.SaveUploadedFile(file, originName); err != nil {
			c.String(400, fmt.Sprintf("Cannot upload server filesystem err: %s", err.Error()))
			return
		}
		defer os.Remove(originName)

		ProcessImage(originName)

		if err := UploadS3(originName, time.Now().String()+".png"); err != nil {
			c.String(400, fmt.Sprintf("Cannot upload S3 err: %s", err.Error()))
		}

		c.String(201, fmt.Sprintf("File %s uploaded successfully.", file.Filename))
	})

	router.Run()

}
