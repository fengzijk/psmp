package task

import (
	"github.com/robfig/cron/v3"
	"go-psmp/mapper"
	"go-psmp/pojo/entity"
	"go-psmp/service"
	"log"
)

func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour |
		cron.Dom | cron.Month | cron.Dow | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}

func InitTask() {

	log.Println("[Cron] Starting...")

	c := newWithSeconds()

	spec := "*/15 * * * * *" //每5秒执行一次

	_, _ = c.AddFunc(spec, func() {

		sendEmailTask()
		log.Println("[Cron] Run sendEmailTask...")

	})

	c.Start()
}

func sendEmailTask() {

	// 查询邮件
	unSendList, err := mapper.FindUnSendList()
	if err != nil {
		return
	}

	if len(unSendList) == 0 {
		return
	}

	if unSendList[0].ID == 0 {
		return
	}

	var successIds []int64

	var failList []entity.EmailRecordEntity
	// 发送邮件
	for i := 0; i < len(unSendList); i++ {
		err := service.SendToMail(unSendList[i].EmailFrom, unSendList[i].EmailTo, unSendList[i].Subject, unSendList[i].Content, unSendList[i].TemplateFlag)
		if err == nil {
			successIds = append(successIds, unSendList[i].ID)
			continue
		} else {
			failEntity := entity.EmailRecordEntity{
				ID:            unSendList[i].ID,
				SendFailCount: unSendList[i].SendFailCount + 1,
				ErrorMsg:      err.Error(),
				SendStatus:    "FAIL",
			}

			failList = append(failList, failEntity)
		}
	}

	// 更新邮件
	service.UpdateEmailSendSuccess(successIds)

	// 更新失败
	service.UpdateEmailSendFail(failList)
}
