package mapper

import (
	"fmt"
	"go-psmp/pojo/entity"
	"go-psmp/utils"
)

func InsertEmailRecord(param entity.EmailRecordEntity) error {
	deres := db.Create(&entity.EmailRecordEntity{ID: utils.NextId(),
		Md5Code:       param.Md5Code,
		EmailFrom:     param.EmailFrom,
		EmailTo:       param.EmailTo,
		Subject:       param.Subject,
		Content:       param.Content,
		SendStatus:    param.SendStatus,
		ErrorMsg:      param.ErrorMsg,
		SendFailCount: param.SendFailCount,
		TemplateFlag:  param.TemplateFlag,
	})

	err := deres.Error
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return err
	}
	return err
}

// FindUnSendList 查询未发送的邮件列表
func FindUnSendList() ([]entity.EmailRecordEntity, error) {
	var emailList []entity.EmailRecordEntity
	dbRes := db.Model(&entity.EmailRecordEntity{}).Where("send_status=?", "WAIT").Find(&emailList)

	err := dbRes.Error
	if err != nil {
		return emailList, err
	}
	return emailList, nil
}

func UpdateEmailSendSuccess(ids []int64) {

	if ids == nil {
		return
	}

	if ids[0] == 0 {
		return
	}

	deRes := db.Model(&entity.EmailRecordEntity{}).Where("id In ?", ids).Updates(entity.EmailRecordEntity{SendStatus: "SUCCESS"})
	err := deRes.Error
	if err != nil {
		fmt.Printf("param failed, err:%v\n", err)

	}
}

func UpdateEmailSendFail(failList []entity.EmailRecordEntity) {

	if failList == nil {
		return
	}

	if failList[0].ID == 0 {
		return
	}

	for _, s := range failList {
		deRes := db.Model(&entity.EmailRecordEntity{}).Where("id=?", s.ID).Updates(s)
		err := deRes.Error
		if err != nil {
			fmt.Printf("param failed, err:%v\n", err)

		}
	}
}

func FindEmailByMd5Code(md5code string) entity.EmailRecordEntity {
	var emailEntity entity.EmailRecordEntity
	dbRes := db.Model(&entity.EmailRecordEntity{}).Where("md5_code=? ", md5code).First(&emailEntity)

	err := dbRes.Error
	if err != nil {
		return emailEntity
	}
	return emailEntity
}
