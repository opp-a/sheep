package models

import (
	"github.com/jinzhu/gorm"
)

type File struct {
	gorm.Model
	ShopID  string `gorm:"column:shopid;NOT NULL" form:"_" json:"shopid"`
	Name    string `gorm:"Type:varchar(516)" form:"name" json:"name"` // string默认长度为255
	Content string `gorm:"Type:varchar(524288);not null;" form:"content" json:"content"`
}
