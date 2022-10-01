package service

import (
	"fmt"
	"log"
	"net/smtp"
	"short-url/config"
	"short-url/pojo/request"
	"strings"
)

func SendToMail(sendUserName, to, subject, body, mailType string) {
	var x = config.EmailConf
	fmt.Println(x)
	user := config.EmailConf.User
	password := config.EmailConf.Password
	host := config.EmailConf.Host

	hp := strings.Split(host, ":")
	//fmt.Println(hp)
	auth := smtp.PlainAuth("", user, password, hp[0])
	var contentType string
	if mailType == "html" {
		contentType = "Content-Type: text/" + mailType + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + sendUserName + "<" + user + ">" + "\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, sendTo, msg)
	fmt.Println(err)
}

func PostMail(email request.SendEmailRequest) bool {

	source := email.SystemName
	if source != "monitor" {
		fmt.Println("Send mail error!,source 认证失败")
		return false
	}
	//println(json.Contacts)
	to := email.Contacts
	if to[0] == "" {
		fmt.Println("Send mail error!,发送人为空")

		return false
	}
	subject := email.Subject
	if strings.TrimSpace(subject) == "" {
		fmt.Println("Send mail error!标题为空")
		return false
	}
	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="iso-8859-15">
			<title>yyyyyy/title>
		</head>
		<body>
			` + fmt.Sprintf(email.Content) +
		`</body>
		</html>`

	sendUserName := "告警平台" //发送邮件的人名称
	log.Println("send email")

	for _, s := range to {
		SendToMail(sendUserName, s, subject, body, "html")

		log.Printf("接收人:%s \n 标题: %s \n 内容: %s \n", s, email.Subject, email.Content)

	}

	return true
}
