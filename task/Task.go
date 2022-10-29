package task

import (
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"go-psmp/mapper"
	"go-psmp/pojo/entity"
	"go-psmp/service"
	"go-psmp/utils/redis"
	"log"
)

func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour |
		cron.Dom | cron.Month | cron.Dow | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}

const (
	sendEmailLockKey  = "lock:send_email_task:"
	sendWxPushLockKey = "lock:send_wx_task:"
)

var wxPushService = service.ServiceGroup.WxPushService
var emailService = service.ServiceGroup.EmailService

func InitTask() {

	log.Println("[Cron] Starting...")

	c := newWithSeconds()

	//
	spec := viper.GetString("task-cron.send-alarm-email")
	_, _ = c.AddFunc(spec, func() {

		sendEmailTask()
		//log.Println("[Cron] Run sendEmailTask...")

	})

	// Agent 告警定时任务
	AgentHeartbeatSpec := viper.GetString("task-cron.agent-heartbeat-alarm")
	_, _ = c.AddFunc(AgentHeartbeatSpec, func() {
		service.AgentHeartbeatAlarm()
		//	log.Println("[Cron] Run AgentHeartbeatAlarmTask...")

	})

	// Agent 告警定时任务
	WxPushSpec := viper.GetString("task-cron.send-alarm-wx")
	_, _ = c.AddFunc(WxPushSpec, func() {
		SendWxPushTask()
		log.Println("[Cron] Run SendWxPushTask...")

	})
	c.Start()
}

func sendEmailTask() {

	var lockKey = sendEmailLockKey + "sendEmailTask"
	lock := redis.Lock(lockKey, 4)
	if !lock {
		log.Println("发送邮件获取锁失败")
		return
	}

	// 解锁
	defer redis.UnLock(lockKey)

	// 查询邮件
	unSendList, err := emailService.FindUnSendList()
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
		err := emailService.SendToMail(unSendList[i])

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

			if failEntity.SendFailCount > 5 {
				failEntity.SendStatus = "SUCCESS"
				failEntity.ErrorMsg = "Exceeded times"
			}
			failList = append(failList, failEntity)
		}
	}

	// 更新邮件
	emailService.UpdateEmailSendSuccess(successIds)

	// 更新失败
	service.UpdateEmailSendFail(failList)
}

func SendWxPushTask() {

	var lockKey = sendWxPushLockKey + "sendWxPushTask"
	lock := redis.Lock(lockKey, 4)
	if !lock {
		log.Println("发送微信推送获取锁失败")
		return
	}

	// 解锁
	defer redis.UnLock(lockKey)

	corpId := viper.GetString("alarm-weixin.corpId")
	corpSecret := viper.GetString("alarm-weixin.corpSecret")
	// 查询
	unSendList, err := mapper.FindWxPushUnSendList()
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

	var failList []entity.WxPushRecordEntity
	// 发送邮件
	for i := 0; i < len(unSendList); i++ {
		err := wxPushService.SendWxPushMessage(corpId, corpSecret, unSendList[i].ToUser, unSendList[i].ToPartyId, unSendList[i].Body, unSendList[i].AgentId)
		if err == "" {
			successIds = append(successIds, unSendList[i].ID)
			continue
		} else {
			failEntity := entity.WxPushRecordEntity{
				ID:            unSendList[i].ID,
				SendFailCount: unSendList[i].SendFailCount + 1,
				ErrorMsg:      err,
				SendStatus:    "FAIL",
			}

			if failEntity.SendFailCount > 5 {
				failEntity.SendStatus = "SUCCESS"
				failEntity.ErrorMsg = "Exceeded times"
			}
			failList = append(failList, failEntity)
		}
	}

	// 更新邮件
	wxPushService.UpdateWxPushSendSuccess(successIds)

	// 更新失败
	wxPushService.UpdateWxPushSendFail(failList)
}
