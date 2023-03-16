package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/menium878/go_infrastruktura/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}
func main() {
	router := gin.Default()
	router.Static("/testowyfolder", "./testowyfolder") // do zapamiętania
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/", func(c *gin.Context) {
		// single file
		file, err := c.FormFile("image")
		if err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"error": "Failed to upload image",
			})
		}
		//log.Println(file.Filename)
		dst := os.Getenv("dir") + file.Filename
		// Upload the file to specific dst.
		err = c.SaveUploadedFile(file, dst)
		if err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"error": "Failed to upload image",
			})
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"image": "/" + os.Getenv("dir") + file.Filename,
			//"title": "COS" + file.Filename,
		})
	})

	router.Run()
}
