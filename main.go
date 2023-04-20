package main

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"text/template"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("ENV Loading Error :", err)
	}

	fmt.Println("RadheShyam")
	subject := "GO-MAIL-TEMPLATE"
	// body := "Bolo SitaRam🙏"
	// html := "<h1>Bolo SitaRam🙏</h1><p>Bolo RadheShyam🙏</p><p>Bolo Bajrangbali ki jayy🙏</p>"
	filePath := "./static/home.html"
	to := []string{"rrrrahulmondar@gmail.com", "rahulmondar@gmail.com"}

	// sendSimpleMail(subject, body, to)
	sendSimpleMailWithHTML(subject, filePath, to)
}

// func sendSimpleMail(subject string, body string, to []string) {
// 	smtpAuth := smtp.PlainAuth(
// 		"",
// 		os.Getenv("SMTP_MAIL"),
// 		os.Getenv("SMTP_PASSWORD"),
// 		os.Getenv("SMTP_HOST"),
// 	)

// 	msg := "Subject: " + subject + "\n" + body

// 	err := smtp.SendMail(
// 		os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"),
// 		smtpAuth,
// 		os.Getenv("SMTP_MAIL"),
// 		to,
// 		[]byte(msg),
// 	)

// 	if err != nil {
// 		fmt.Println("Mail Sending Err:", err)
// 	} else {
// 		fmt.Println("Mail Sent successfully")
// 	}
// }

func sendSimpleMailWithHTML(subject string, filePath string, to []string) {
	// Get HTML
	var body bytes.Buffer
	file, fileErr := template.ParseFiles(filePath)
	if fileErr != nil {
		log.Fatal("Template Parsing Error :", fileErr)
	}

	templateExecutionErr := file.Execute(&body, struct{ Greeting string }{Greeting: "RADHE🙏SHYAM"})
	if templateExecutionErr != nil {
		log.Fatal("Template Parsing Error :", fileErr)
	}

	smtpAuth := smtp.PlainAuth(
		"",
		os.Getenv("SMTP_MAIL"),
		os.Getenv("SMTP_PASSWORD"),
		os.Getenv("SMTP_HOST"),
	)

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"
	msg := "Subject: " + subject + "\n" + headers + "\n\n" + body.String()

	err := smtp.SendMail(
		os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"),
		smtpAuth,
		os.Getenv("SMTP_MAIL"),
		to,
		[]byte(msg),
	)

	if err != nil {
		fmt.Println("Mail Sending Err:", err)
	} else {
		fmt.Println("Mail Sent successfully")
	}
}
