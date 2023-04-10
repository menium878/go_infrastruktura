package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/menium878/go_infrastruktura/initializers"
	"github.com/menium878/go_infrastruktura/routes"
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
	// Dodanie wielu zdjeć kod
	router.MaxMultipartMemory = 8 << 20 // 8 MiB ogranieczenie
	router.POST("/", routes.ImageUpload)
	// TODO: zrobienie route które wysyła maila z wynikiem
	router.Run()
}
