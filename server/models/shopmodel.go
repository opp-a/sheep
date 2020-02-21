package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego"
	"github.com/segmentio/ksuid"
)

type Shop struct {
	ShopID    string    `gorm:"primary_key" form:"shopid" json:"shopid"`
	Name      string    `gorm:"Type:varchar(128);NOT NULL" form:"name" json:"name"` // string默认长度为255
	Icons     []File    `gorm:"ForeignKey:ShopID" form:"icons" json:"icons"`
	Pricein   uint      `gorm:"not null;default:'0'" form:"pricein" json:"pricein"`
	Priceout  uint      `gorm:"not null;default:'0'" form:"priceout" json:"priceout"`
	Num       uint      `gorm:"not null;default:'0'" form:"num" json:"num"`
	Desc      string    `gorm:"Type:varchar(256);not null;default:'its great'" form:"desc" json:"desc"`
	CreatedAt time.Time `json:"addtime"`
	UpdatedAt time.Time `json:"updatetime"`
}

func (s *Shop) AfterFind() (err error) {
	if len(s.Icons) == 0 {
		s.Icons = make([]File, 0)
	}
	return
}

func AddShop(shops []Shop) error {
	if len(shops) <= 0 {
		return errors.New("no shop need to insert!")
	}
	var err error
	for _, shop := range shops {
		uid, _ := ksuid.NewRandomWithTime(time.Now())
		shop.ShopID = uid.String()

		if err = db.Save(&shop).Error; err != nil {
			beego.Error("add shop fail! err:", err)
		}
	}
	return err
}

func DeleteShop(shopids []string) error {
	if len(shopids) <= 0 {
		return nil
	}
	if err := db.Delete(&Shop{}, "shop_id in (?)", shopids).Error; err != nil {
		beego.Error("delete shop fail! err:", err)
		return err
	}
	return nil
}

func UpdateShop(shopid string, shop Shop) error {
	if shopid == "" {
		return nil
	}

	if err := db.Update(&shop, "shop_id = ?", shopid).Error; err != nil {
		beego.Error("update shop fail! err:", err)
		return err
	}
	return nil
}

func ListShop(pageindex, pagesize uint) (interface{}, error) {
	rinfos := struct {
		Total int    `json:"total"`
		Infos []Shop `json:"infos"`
	}{0, make([]Shop, 0)}
	if pageindex == 0 {
		pageindex = 1
	}
	index := (pageindex - 1) * pagesize
	if err := db.Order("created_at desc").
		Limit(pagesize).Offset(index).
		Find(&rinfos.Infos).Error; err != nil {
		beego.Error("list shop fail! err:", err)
		return rinfos, err
	}
	if err := db.Model(&Shop{}).Count(&rinfos.Total).Error; err != nil {
		beego.Error("list shop fail! err:", err)
		return rinfos, err
	}

	return rinfos, nil
}
