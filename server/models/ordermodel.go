package models

import (
	// "errors"
	"time"
	// "github.com/astaxie/beego"
	// "github.com/segmentio/ksuid"
)

const (
	OrderTypeForPay = iota
	OrderTypeForHistory
)

type Order struct {
	OrderID    string      `gorm:"primary_key" form:"orderid" json:"orderid"`
	OrderType  uint        `form:"name" json:"name"` // string默认长度为255
	Myshops    []OrderShop `gorm:"ForeignKey:ShopID" form:"icons" json:"icons"`
	PriceTotal uint        `gorm:"not null;default:'0'" form:"pricein" json:"pricein"`
	CreatedAt  time.Time   `json:"addtime"`
	UpdatedAt  time.Time   `json:"updatetime"`
}
type OrderShop struct {
	ShopID   string `gorm:"primary_key" form:"shopid" json:"shopid"`
	Name     string `gorm:"Type:varchar(128);NOT NULL" form:"name" json:"name"` // string默认长度为255
	Icons    []File `gorm:"ForeignKey:ShopID" form:"icons" json:"icons"`
	Priceout uint   `gorm:"not null;default:'0'" form:"priceout" json:"priceout"`
	Num      uint   `gorm:"not null;default:'0'" form:"num" json:"num"`
}
