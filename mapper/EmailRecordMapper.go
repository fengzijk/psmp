package mapper

import (
	"go-psmp/config"
	"go-psmp/pojo/entity"
	"go-psmp/pojo/model/page"
	"go-psmp/utils"
	"strings"
)

type EmailRecordMapper struct {
}

func (emailMapper *EmailRecordMapper) InsertEmailRecord(param entity.EmailRecordEntity) error {
	deres := db.Create(&entity.EmailRecordEntity{ID: utils.NextId(),
		Md5Code:       param.Md5Code,
		FromName:      param.FromName,
		EmailTo:       param.EmailTo,
		Subject:       param.Subject,
		Body:          param.Body,
		SendStatus:    param.SendStatus,
		ErrorMsg:      param.ErrorMsg,
		SendFailCount: param.SendFailCount,
		TemplateFlag:  param.TemplateFlag,
	})

	err := deres.Error
	if err != nil {
		config.Log.Info("insert failed, err:")
		return err
	}
	return err
}

// FindUnSendList 查询未发送的邮件列表
func (emailMapper *EmailRecordMapper) FindUnSendList() ([]entity.EmailRecordEntity, error) {
	var emailList []entity.EmailRecordEntity
	dbRes := db.Model(&entity.EmailRecordEntity{}).Where("send_status=?", "WAIT").Find(&emailList)

	err := dbRes.Error
	if err != nil {
		return emailList, err
	}
	return emailList, nil
}

func (emailMapper *EmailRecordMapper) UpdateEmailSendSuccess(ids []int64) {

	if ids == nil {
		return
	}

	if ids[0] == 0 {
		return
	}

	deRes := db.Model(&entity.EmailRecordEntity{}).Where("id In ?", ids).Updates(entity.EmailRecordEntity{SendStatus: "SUCCESS"})
	err := deRes.Error
	if err != nil {
		config.Log.Info("param failed, err:%v\n")

	}
}

func (emailMapper *EmailRecordMapper) UpdateEmailSendFail(failList []entity.EmailRecordEntity) {

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
			config.Log.Info("param failed, ")

		}
	}
}

func (emailMapper *EmailRecordMapper) FindEmailByMd5Code(md5code string) entity.EmailRecordEntity {
	var emailEntity entity.EmailRecordEntity
	dbRes := db.Model(&entity.EmailRecordEntity{}).Where("md5_code=? ", md5code).First(&emailEntity)

	err := dbRes.Error
	if err != nil {
		return emailEntity
	}
	return emailEntity
}

func (emailMapper *EmailRecordMapper) ListPageEmailByAdmin(status string, pageNum, pageSize int) (*page.PagerModel, error) {
	var emailList []entity.EmailRecordEntity
	var count int64
	if len(status) == 0 {
		db.Model(&entity.EmailRecordEntity{}).Offset(-1).Limit(-1).Count(&count)
		db.Model(&entity.EmailRecordEntity{}).Scopes(Paginate(pageNum, pageSize)).Order("created_at desc").Find(&emailList).Limit(pageSize)

	} else {
		db.Model(&entity.EmailRecordEntity{}).Where("send_status=?", strings.ToUpper(status)).Offset(-1).Limit(-1).Count(&count)
		db.Model(&entity.EmailRecordEntity{}).Scopes(Paginate(pageNum, pageSize)).Where("send_status=?", strings.ToUpper(status)).Order("created_at  desc").Find(&emailList).Limit(pageSize)
	}

	pager := page.CreatePager(pageNum, pageSize, int(count), emailList)
	return pager, nil
}
