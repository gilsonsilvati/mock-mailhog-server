package main

import (
	"fmt"
	"log"
	"net/smtp"
)

const (
	from     = "gilson.silva@localhost.com"
	to       = "recebedor@localhost.com"
	host     = "localhost"
	port     = "1025"
	subject  = "Teste de envio de email"
	body     = "Mensagem de teste"
)

func main() {
	toList := []string{to}
	message := buildMessage(from, toList, subject, body)
	auth := smtp.PlainAuth("", "", "", host)

	if err := sendEmail(host+":"+port, auth, from, toList, message); err != nil {
		log.Fatal(err)
	}

	fmt.Println(":::: Email enviado com sucesso! ::::")
}

func buildMessage(from string, to []string, subject, body string) []byte {
	return []byte(
		"From: " + from + "\r\n" +
			"To: " + to[0] + "\r\n" +
			"Subject: " + subject + "\r\n\r\n" +
			body,
	)
}

func sendEmail(addr string, auth smtp.Auth, from string, to []string, msg []byte) error {
	return smtp.SendMail(addr, auth, from, to, msg)
}
