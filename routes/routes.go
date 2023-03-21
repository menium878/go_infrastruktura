package routes

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ImageUpload(c *gin.Context) {
	//popisać timeoutty contexty
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

}
