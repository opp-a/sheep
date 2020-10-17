package models

import (
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
)

type File struct {
	gorm.Model
	ShopID  string `gorm:"column:shopid;NOT NULL" form:"_" json:"shopid"`
	Name    string `gorm:"Type:varchar(516)" form:"name" json:"name"` // string默认长度为255
	Content string `gorm:"Type:varchar(524288);not null;" form:"content" json:"content"`
}

func GetFile(shopid, content string) *File {
	var icon File
	if err := db.Where("shopid = ? AND content = ?", shopid, content).First(&icon).Error; err != nil {
		beego.Error(err)
		return nil
	} else {
		return &icon
	}
	return nil
}

func DeleteFile(tx *gorm.DB, shopid string) error {
	if err := db.Where("shopid = ?", shopid).Delete(&File{}).Error; err != nil {
		beego.Error(err)
		tx.Rollback()
		return err
	}

	return nil
}
