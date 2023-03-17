package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/menium878/go_infrastruktura/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	router := gin.Default()
	router.Static(os.Getenv("static"), "."+os.Getenv("static"))
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/", func(c *gin.Context) {
		// multiple files
		form, err := c.MultipartForm()
		if err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"error": "Failed to upload image(s)",
			})
			return
		}

		files := form.File["image"]
		if len(files) == 0 {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"error": "No files uploaded",
			})
			return
		}

		for i, file := range files {
			dst := os.Getenv("dir") + file.Filename
			err = c.SaveUploadedFile(file, dst)
			if err != nil {
				c.HTML(http.StatusOK, "index.html", gin.H{
					"error": "Failed to upload image(s)",
				})
				return
			}
			c.HTML(http.StatusOK, "index.html", gin.H{
				"image" + strconv.Itoa(i): "/" + os.Getenv("dir") + file.Filename,
			})
		}

	})

	router.Run()
}
