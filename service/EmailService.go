package service

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/go-gomail/gomail"
	"go-psmp/config"
	"go-psmp/mapper"
	"go-psmp/pojo/entity"
	"go-psmp/pojo/request"
	"go-psmp/utils/short"
	"log"
	"strings"
)

func SendToMail(subject, body string) error {

	// 主题
	config.Message.SetHeader("Subject", subject)

	// 正文
	config.Message.SetBody("text/html", body)

	d := gomail.NewDialer(config.EmailConf.Host, config.EmailConf.Port, config.EmailConf.User, config.EmailConf.Password)
	// 发送
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := d.DialAndSend(config.Message)
	if err != nil {
		log.Print(err)
	}

	return err
}

func SaveMail(email request.SendEmailRequest) bool {

	//source := email.SystemName
	//if source != "psmp-agent" {
	//	fmt.Println("Send mail error!,source 认证失败")
	//	return false
	//}

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
			<title>gaojing/title>
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

	recordEntity := mapper.FindEmailByMd5Code(insert.Md5Code)
	if recordEntity.ID != 0 {
		return
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
