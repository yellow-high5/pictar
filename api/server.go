package server

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yellow-high5/pictar/helper"
)

func Boot() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pictar server startup!!",
		})
	})

	router.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("image")

		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}

		if !helper.Contains([]string{".png", ".jpg", ".jpeg", ".bmp", ".gif", ".tiff", ".tif"}, strings.ToLower(filepath.Ext(file.Filename))) {
			c.String(http.StatusBadRequest, fmt.Sprint("cannot read as image file"))
			return
		}

		log.Println(file.Filename)
		filename := filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully.", file.Filename))
	})
	router.Run()

}
