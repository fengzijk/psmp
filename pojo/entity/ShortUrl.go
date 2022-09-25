package entity

import "gorm.io/gorm"

type ShortURL struct {
	gorm.Model
	ShortUrl string
	LongUrl  string
}
