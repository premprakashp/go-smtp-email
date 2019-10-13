package main

import (
	"gopkg.in/gomail.v2"
)

func main() {
	d := gomail.NewDialer("smtp.gmail.com", 587, "username", "password")
	// configure TLS insecure if using local tls servers without valid certs for testing
	//d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetHeader("From", "fromemail@example.com")
	m.SetHeader("To", "toemail1@example.com", "toemail2@example.com")
	m.SetHeader("Cc", "toemail3@example.com", "toemail4@example.com")
	m.SetHeader("Subject", "Email Subject text")
	/*
		Use text/plain for text email
	*/
	m.SetBody("text/html", "Hello <b>HTML email</b> body")

	/*
		Attachments if required
	*/
	m.Attach("/path_to_local_file_on_server")

	// Dial address and send email
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}
