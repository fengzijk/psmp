package mapper

import (
	"fmt"
	"go-psmp/pojo/entity"
	"go-psmp/utils"
)

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
		fmt.Printf("insert failed, err:%v\n", err)
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
		fmt.Printf("param failed, err:%v\n", err)

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
			fmt.Printf("param failed, err:%v\n", err)

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
