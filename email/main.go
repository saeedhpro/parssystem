package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

func main() {
	from := "saeedhpro@gmail.com"

	user := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWORD")

	to := []string{
		"saeedheydari@example.com",
	}

	addr := fmt.Sprintf("%s:%s", os.Getenv("MAIL_HOST"), os.Getenv("MAIL_PORT"))
	host := os.Getenv("MAIL_HOST")

	msg := []byte("")

	auth := smtp.PlainAuth("", user, password, host)

	err := smtp.SendMail(addr, auth, from, to, msg)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Email sent successfully")
}
