package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

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
func ImageUpload(c *gin.Context) {
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

	var result string
	for _, file := range files {
		dst := os.Getenv("dir") + file.Filename
		err = c.SaveUploadedFile(file, dst)
		if err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"error": "Failed to upload image(s)",
			})
			return
		}

		// execute Python script that calls AI model API
		cmd := exec.Command("python", "ai_script.py", dst)
		out, err := cmd.Output()
		if err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"error": "Failed to call AI model",
			})
			return
		}

		// parse JSON response
		var aiResult struct {
			Result string `json:"result"`
		}
		err = json.Unmarshal(out, &aiResult)
		if err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"error": "Failed to parse AI model response",
			})
			return
		}

		result = aiResult.Result
	}

	// send email with result using SendEmailHandler
	req := struct {
		Recipient string `json:"recipient"`
		Text      string `json:"text"`
	}{
		Recipient: "recipient@example.com",
		Text:      result,
	}
	jsonReq, err := json.Marshal(req)
	if err != nil {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"error": "Failed to create email request body",
		})
		return
	}
	resp, err := http.Post("http://127.0.0.1:9000/send-email", "application/json", bytes.NewBuffer(jsonReq))
	if err != nil {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"error": "Failed to send email request",
		})
		return
	}
	defer resp.Body.Close()

	c.HTML(http.StatusOK, "index.html", gin.H{
		"result": result,
	})
}
