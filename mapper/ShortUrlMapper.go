package mapper

import (
	"fmt"
	"go-psmp/pojo/entity"
	"go-psmp/utils"
	"gorm.io/gorm"
)

func InsertShortUrl(param entity.ShortURLEntity) error {
	deres := db.Create(&entity.ShortURLEntity{ID: utils.NextId(),
		Md5Code:    param.Md5Code,
		LongParam:  param.LongParam,
		ShortParam: param.ShortParam,
		BizType:    param.BizType})
	err := deres.Error
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return err
	}
	return err
}

func SelectShortUrlInfoById(id int) (entity.ShortURLEntity, error) {
	var shortURLEntity entity.ShortURLEntity
	dbRes := db.Model(&entity.ShortURLEntity{}).Where("Id = ?", id).First(&shortURLEntity)
	err := dbRes.Error
	if err != nil {
		return shortURLEntity, err
	}
	fmt.Println(shortURLEntity)
	return shortURLEntity, nil
}

func SelectShortUrlInfoByParam(param string, paramType string) entity.ShortURLEntity {
	var shortURL entity.ShortURLEntity
	var dbRes *gorm.DB
	if paramType == "url" {
		dbRes = db.Model(&entity.ShortURLEntity{}).Where("short_url=? and biz_type=?", param, paramType).First(&shortURL)
	} else {
		dbRes = db.Model(&entity.ShortURLEntity{}).Where("short_param=? and biz_type=?", param, paramType).First(&shortURL)
	}

	err := dbRes.Error
	if err != nil {
		return shortURL
	}
	fmt.Println(shortURL)
	return shortURL
}

func SelectShortUrlInfoByMd5Code(md5code string) entity.ShortURLEntity {
	var shortURL entity.ShortURLEntity
	dbRes := db.Model(&entity.ShortURLEntity{}).Where("md5_code=? ", md5code).First(&shortURL)

	err := dbRes.Error
	if err != nil {
		return shortURL
	}
	fmt.Println(shortURL)
	return shortURL
}

func SelectShortUrlInfoByShortParam(shortParam string) entity.ShortURLEntity {
	var shortURL entity.ShortURLEntity
	dbRes := db.Model(&entity.ShortURLEntity{}).Where("short_param=? ", shortParam).First(&shortURL)

	err := dbRes.Error
	if err != nil {
		return shortURL
	}
	fmt.Println(shortURL)
	return shortURL
}
