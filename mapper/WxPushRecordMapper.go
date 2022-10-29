package mapper

import (
	"go-psmp/config"
	"go-psmp/pojo/entity"
	"go-psmp/utils"
)

type WxPushRecordMapper struct {
}

func InsertWxPushRecord(param entity.WxPushRecordEntity) error {
	deres := db.Create(&entity.WxPushRecordEntity{ID: utils.NextId(),
		Md5Code:       param.Md5Code,
		ToUser:        param.ToUser,
		ToPartyId:     param.ToPartyId,
		AgentId:       param.AgentId,
		Body:          param.Body,
		SendStatus:    param.SendStatus,
		ErrorMsg:      param.ErrorMsg,
		SendFailCount: param.SendFailCount,
	})

	err := deres.Error
	if err != nil {
		config.Log.Info("insert failed, err:")
		return err
	}
	return err
}

// FindWxPushUnSendList 查询未发送的列表
func FindWxPushUnSendList() ([]entity.WxPushRecordEntity, error) {
	var recordEntities []entity.WxPushRecordEntity
	dbRes := db.Model(&entity.WxPushRecordEntity{}).Where("send_status=?", "WAIT").Find(&recordEntities)

	err := dbRes.Error
	if err != nil {
		return recordEntities, err
	}
	return recordEntities, nil
}

func UpdateWxPushSendSuccess(ids []int64) {

	if ids == nil {
		return
	}

	if ids[0] == 0 {
		return
	}

	deRes := db.Model(&entity.WxPushRecordEntity{}).Where("id In ?", ids).Updates(entity.WxPushRecordEntity{SendStatus: "SUCCESS"})
	err := deRes.Error
	if err != nil {
		config.Log.Info("param failed, err:%v\n")

	}
}

func UpdateWxPushSendFail(failList []entity.WxPushRecordEntity) {

	if failList == nil {
		return
	}

	if failList[0].ID == 0 {
		return
	}

	for _, s := range failList {
		deRes := db.Model(&entity.WxPushRecordEntity{}).Where("id=?", s.ID).Updates(s)
		err := deRes.Error
		if err != nil {
			config.Log.Info("param failed, err:%v\n")

		}
	}
}

func FindWxPushByMd5Code(md5code string) entity.WxPushRecordEntity {
	var recordEntity entity.WxPushRecordEntity
	dbRes := db.Model(&entity.WxPushRecordEntity{}).Where("md5_code=? ", md5code).First(&recordEntity)

	err := dbRes.Error
	if err != nil {
		return recordEntity
	}
	return recordEntity
}
