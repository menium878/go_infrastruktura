package routes

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

func ImageUpload(c *gin.Context) {
	// ? dowiedz się jak będzie działało jeśli wrzucisz w cloud i potem zapis tego zdjęcia
	// !popisać timeoutty contexty
	// !Dopisać wywołanie pytanie o maila i podpisz zdjęcie nim potem wyciągniemy sobie z DB mail po nazwie zdjęcia
	// ?Przemyśleć czy może logowanie i rejestracja wtedy punkt wyżej działa inaczej
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
	// TODO: send to api that we wanna start the program when we get the picture
}

func SendEmailHandler(c *gin.Context) {
	// Parse the request body
	var req struct {
		Recipient string `json:"recipient"`
		Text      string `json:"text"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, "Invalid request body")
		return
	}

	// Set up the email message
	m := gomail.NewMessage()
	m.SetHeader("From", "sender@example.com")
	m.SetHeader("To", req.Recipient)
	m.SetHeader("Subject", "Test email")
	m.SetBody("text/plain", req.Text)

	// Set up the SMTP server details
	d := gomail.NewDialer("smtp.gmail.com", 587, "sender@example.com", "password")

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		c.String(http.StatusInternalServerError, "Failed to send email")
		return
	}

	c.String(http.StatusOK, "Email sent successfully")
}
