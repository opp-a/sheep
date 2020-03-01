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
	if err := db.Unscoped().Delete(&Shop{}, "shop_id in (?)", shopids).Error; err != nil {
		beego.Error("delete shop fail! err:", err)
		return err
	}
	return nil
}

func UpdateShop(shopid string, shop Shop) error {
	if shopid == "" {
		return nil
	}
	shop.ShopID = shopid

	beego.Debug(shop.Desc)
	if err := db.Save(&shop).Error; err != nil {
		beego.Error("update shop fail! err:", err)
		return err
	}
	return nil
}

func ListShop(pageindex, pagesize uint) (interface{}, error) {
	type FrontShop struct {
		ShopID    string    `json:"shopid"`
		Name      string    `json:"name"` // string默认长度为255
		Icons     []string  `json:"icons"`
		Pricein   uint      `json:"pricein"`
		Priceout  uint      `json:"priceout"`
		Num       uint      `json:"num"`
		Desc      string    `json:"desc"`
		CreatedAt time.Time `json:"addtime"`
	}
	rinfos := struct {
		Total int         `json:"total"`
		Infos []FrontShop `json:"infos"`
	}{0, make([]FrontShop, 0)}
	if pageindex == 0 {
		pageindex = 1
	}
	index := (pageindex - 1) * pagesize

	var shops []Shop = make([]Shop, 0)
	if err := db.Order("created_at desc").
		Limit(pagesize).Offset(index).
		Find(&shops).Error; err != nil {
		beego.Error("list shop fail! err:", err)
		return rinfos, err
	}
	// fill file
	for index, shop := range shops {
		rinfos.Infos = append(rinfos.Infos, FrontShop{
			ShopID:    shop.ShopID,
			Name:      shop.Name,
			Icons:     make([]string, 0),
			Pricein:   shop.Pricein,
			Priceout:  shop.Priceout,
			Num:       shop.Num,
			Desc:      shop.Desc,
			CreatedAt: shop.CreatedAt})

		var files []File = make([]File, 0)
		if err := db.Model(&shop).Related(&files, "Icons").Error; err != nil {
			beego.Error("association file fail! err:", err)
			return rinfos, err
		}
		for _, file := range files {
			rinfos.Infos[index].Icons = append(rinfos.Infos[index].Icons, file.Content)
		}
	}
	if err := db.Model(&Shop{}).Count(&rinfos.Total).Error; err != nil {
		beego.Error("list shop fail! err:", err)
		return rinfos, err
	}

	return rinfos, nil
}
