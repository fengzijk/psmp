package mapper

import (
	"fmt"
	"gorm.io/gorm"
	"short-url/pojo/entity"
	"short-url/utils"
)

func InsertShortUrl(param entity.ShortURL) error {
	deres := db.Create(&entity.ShortURL{ID: utils.NextId(),
		Md5Code:     param.Md5Code,
		LongParam:   param.LongParam,
		ShortParam:  param.ShortParam,
		RedirectUrl: param.RedirectUrl,
		BizType:     param.BizType})
	err := deres.Error
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return err
	}
	return err
}

func SelectShortUrlInfoById(id int) (entity.ShortURL, error) {
	var shortURL entity.ShortURL
	dbRes := db.Model(&entity.ShortURL{}).Where("Id = ?", id).First(&shortURL)
	err := dbRes.Error
	if err != nil {
		return shortURL, err
	}
	fmt.Println(shortURL)
	return shortURL, nil
}

func SelectShortUrlInfoByParam(param string, paramType string) entity.ShortURL {
	var shortURL entity.ShortURL
	var dbRes *gorm.DB
	if paramType == "url" {
		dbRes = db.Model(&entity.ShortURL{}).Where("short_url=? and biz_type=?", param, paramType).First(&shortURL)
	} else {
		dbRes = db.Model(&entity.ShortURL{}).Where("short_param=? and biz_type=?", param, paramType).First(&shortURL)
	}

	err := dbRes.Error
	if err != nil {
		return shortURL
	}
	fmt.Println(shortURL)
	return shortURL
}

func SelectShortUrlInfoByMd5Code(md5code string) entity.ShortURL {
	var shortURL entity.ShortURL
	dbRes := db.Model(&entity.ShortURL{}).Where("md5_code=? ", md5code).First(&shortURL)

	err := dbRes.Error
	if err != nil {
		return shortURL
	}
	fmt.Println(shortURL)
	return shortURL
}

func SelectShortUrlInfoByShortParam(shortParam string) entity.ShortURL {
	var shortURL entity.ShortURL
	dbRes := db.Model(&entity.ShortURL{}).Where("short_param=? ", shortParam).First(&shortURL)

	err := dbRes.Error
	if err != nil {
		return shortURL
	}
	fmt.Println(shortURL)
	return shortURL
}
