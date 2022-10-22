package service

import (
	"crypto/tls"
	"fmt"
	"github.com/go-gomail/gomail"
	"go-psmp/config"
	"go-psmp/mapper"
	"go-psmp/pojo/entity"
	"go-psmp/pojo/model/page"
	"go-psmp/pojo/request"
	"go-psmp/utils/short"
	"log"
	"strings"
)

type EmailService struct {
}

func (emailService *EmailService) SendToMail(recordEntity entity.EmailRecordEntity) error {

	var toUserListStr []string
	var CcUserListStr []string

	// 发件人
	// 第三个参数为发件人别名，如"李大锤"，可以为空（此时则为邮箱名称）
	config.Message.SetAddressHeader("From", config.EmailConf.User, recordEntity.FromName)

	if len(recordEntity.EmailTo) == 0 {
		return fmt.Errorf("收件人不能为空")
	}

	for _, tmp := range strings.Split(recordEntity.EmailTo, ";") {
		toUserListStr = append(toUserListStr, strings.TrimSpace(tmp))
	}

	// 收件人可以有多个，故用此方式
	config.Message.SetHeader("To", toUserListStr...)

	//抄送列表
	if len(recordEntity.EmailCc) != 0 {
		for _, tmp := range strings.Split(recordEntity.EmailCc, ";") {
			CcUserListStr = append(CcUserListStr, strings.TrimSpace(tmp))
		}
		config.Message.SetHeader("Cc", CcUserListStr...)
	}

	// 主题
	config.Message.SetHeader("Subject", recordEntity.Subject)

	// 正文
	config.Message.SetBody("text/html", recordEntity.Body)

	d := gomail.NewDialer(config.EmailConf.Host, config.EmailConf.Port, config.EmailConf.User, config.EmailConf.Password)

	// 发送
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := d.DialAndSend(config.Message)
	if err != nil {
		log.Print(err)
	}

	return err
}

func (emailService *EmailService) SaveMail(email request.SendEmailRequest) bool {

	if len(email.ToUser) == 0 {
		fmt.Println("Send mail error!,接收人为空")

		return false
	}
	subject := email.Subject
	if strings.TrimSpace(subject) == "" {
		fmt.Println("Send mail error!标题为空")
		return false
	}

	if strings.TrimSpace(email.Body) == "" {
		fmt.Println("Send mail error!邮件内容为空")
		return false
	}

	sendUserName := email.FromName //发送邮件的人名称
	log.Println("send email")

	emailService.saveEmailRecord(sendUserName, email.ToUser, email.CcUser, subject, email.Body, "HTML")

	return true
}

func (emailService *EmailService) UpdateEmailSendSuccess(ids []int64) {
	mapper.UpdateEmailSendSuccess(ids)
}

func (emailService *EmailService) saveEmailRecord(sendUserName, emailTo, emailCc, subject, body, templateFlag string) {

	// 构造发送邮件记录
	insert := entity.EmailRecordEntity{
		Md5Code:       short.Get16MD5Encode(sendUserName + emailTo + emailCc + subject + body + templateFlag),
		FromName:      sendUserName,
		EmailTo:       emailTo,
		Subject:       subject,
		Body:          body,
		EmailCc:       emailCc,
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

//	func getEmailToString(to []string) string {
//		var bt bytes.Buffer
//		for _, s := range to {
//			bt.WriteString(s)
//			bt.WriteString(";")
//		}
//		return bt.String()
//
// }
func (emailService *EmailService) ListPageEmailByAdmin(status string, pageNum, pageSize int) *page.PagerModel {

	unSendList, err := mapper.ListPageEmailByAdmin(status, pageNum, pageSize)
	if err != nil {
		return page.CreatePager(pageNum, pageSize, 0, entity.EmailRecordEntity{})
	}
	return unSendList
}
