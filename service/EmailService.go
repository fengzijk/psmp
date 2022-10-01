package service

import (
	"bytes"
	"fmt"
	"go-psmp/config"
	"go-psmp/mapper"
	"go-psmp/pojo/entity"
	"go-psmp/pojo/request"
	"go-psmp/utils/short"
	"log"
	"net/smtp"
	"strings"
)

func SendToMail(sendUserName, to, subject, body, mailType string) error {
	user := config.EmailConf.User
	password := config.EmailConf.Password
	host := config.EmailConf.Host

	hp := strings.Split(host, ":")
	//fmt.Println(hp)
	auth := smtp.PlainAuth("", user, password, hp[0])
	var contentType string
	if mailType == "HTML" {
		contentType = "Content-Type: text/" + mailType + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + sendUserName + "<" + user + ">" + "\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, sendTo, msg)
	return err
}

func SaveMail(email request.SendEmailRequest) bool {

	source := email.SystemName
	if source != "monitor" {
		fmt.Println("Send mail error!,source 认证失败")
		return false
	}

	to := email.EmailTo
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

	emailTo := getEmailToString(to)
	saveEmailRecord(sendUserName, emailTo, subject, body, "HTML")

	return true
}

func UpdateEmailSendSuccess(ids []int64) {
	mapper.UpdateEmailSendSuccess(ids)
}

func saveEmailRecord(sendUserName, emailTo, subject, content, templateFlag string) {

	// 构造发送邮件记录
	insert := entity.EmailRecordEntity{
		Md5Code:       short.Get16MD5Encode(sendUserName + emailTo + subject + content + templateFlag),
		EmailFrom:     sendUserName,
		EmailTo:       emailTo,
		Subject:       subject,
		Content:       content,
		SendStatus:    "WAIT",
		SendFailCount: 0,
		TemplateFlag:  templateFlag,
	}
	_ = mapper.InsertEmailRecord(insert)
}

func UpdateEmailSendFail(failList []entity.EmailRecordEntity) {
	mapper.UpdateEmailSendFail(failList)
}

func getEmailToString(to []string) string {
	var bt bytes.Buffer
	for _, s := range to {
		bt.WriteString(s)
		bt.WriteString(";")
	}
	return bt.String()

}
