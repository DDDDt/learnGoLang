package main

import (
	"fmt"
	"net/smtp"
)

func SendToMail() error {

	auth := smtp.PlainAuth("", "", "", "smtp.163.com")
	toUser := []string{""}
	content := []byte("To:  \r\nFrom: >\r\nSubject: 使用golang 发送邮件\r\n\r\n\r\n<html><h1>邮件发送</h1></html>")
	err := smtp.SendMail("smtp.163.com:25", auth, "", toUser, content)
	return err
}

func main() {
	err := SendToMail()
	fmt.Println(err)
}
