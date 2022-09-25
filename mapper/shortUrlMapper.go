package mapper

import (
	"fmt"
	"short-url/pojo/entity"
)

func InsertShortUrl(short entity.ShortURL) error {
	deres := db.Create(&entity.ShortURL{ShortUrl: short.ShortUrl, LongUrl: short.LongUrl})
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

func SelectShortUrlInfoByEntity(param entity.ShortURL) entity.ShortURL {
	var shortURL entity.ShortURL
	dbRes := db.Model(&entity.ShortURL{}).Find(&param).First(&shortURL)
	err := dbRes.Error
	if err != nil {
		return shortURL
	}
	fmt.Println(shortURL)
	return shortURL
}
