package models

import (
	"encoding/json"
	"time"

	"github.com/astaxie/beego"
	"github.com/segmentio/ksuid"
)

const (
	OrderTypeForPay = iota
	OrderTypeForHistory
)

type Order struct {
	OrderID    string    `gorm:"primary_key" form:"orderid" json:"orderid"`
	UserName   string    `gorm:"Type:varchar(255);not null;" form:"username" json:"username"`
	OrderType  uint      `form:"ordertype" json:"ordertype"`
	ShopIDs    string    `gorm:"Type:varchar(2048);not null;default:'[]'" form:"shopids" json:"shopids"`
	PriceTotal uint      `gorm:"not null;default:'0'" form:"pricetotal" json:"pricetotal"`
	CreatedAt  time.Time `json:"addtime"`
	UpdatedAt  time.Time `json:"updatetime"`
}

func GetShoppingCar(username string) (interface{}, error) {
	type FrontOrderShop struct {
		ShopID   string `json:"shopid"`
		Name     string `json:"name"`
		Cover    string `json:"cover"`
		Priceout uint   `json:"priceout"`
		Num      uint   `json:"num"`
	}
	type FrontOrder struct {
		OrderID    string           `json:"orderid"`
		Myshops    []FrontOrderShop `json:"myshops"`
		PriceTotal uint             `json:"pricetotal"`
		CreatedAt  time.Time        `json:"addtime"`
	}
	var rcar FrontOrder

	var order Order
	if db.Model(&Order{}).
		Where("order_type = ? AND user_name = ?", OrderTypeForPay, username).
		First(&order).
		RecordNotFound() {
		uid, _ := ksuid.NewRandomWithTime(time.Now())
		if err := db.FirstOrCreate(&order, Order{
			OrderID:   uid.String(),
			OrderType: OrderTypeForPay,
			UserName:  username}).Error; err != nil {
			beego.Error("get and init shopping car fail! err:", err)
			return rcar, err
		}
	}
	rcar.OrderID = order.OrderID
	rcar.CreatedAt = order.CreatedAt
	rcar.PriceTotal = 0
	rcar.Myshops = make([]FrontOrderShop, 0)

	// fill my shops
	type CarShop struct {
		ShopID string `json:"shopid"`
		Num    uint   `json:"num"`
	}
	carshops := make([]CarShop, 0)
	if err := json.Unmarshal([]byte(order.ShopIDs), &carshops); err != nil {
		beego.Error("Unmarshal shopping car shop fail! err:", err)
		return rcar, err
	}
	for _, carshop := range carshops {
		// get shop
		var shop Shop
		if err := db.First(&shop, "shop_id = ?", carshop.ShopID).Error; err != nil {
			beego.Error("get shop fail! err:", err)
			return rcar, err
		}
		var cover string
		if len(shop.Icons) > 0 {
			cover = shop.Icons[0].Content
		}
		rcar.PriceTotal = rcar.PriceTotal + carshop.Num*shop.Priceout
		rcar.Myshops = append(rcar.Myshops, FrontOrderShop{
			ShopID:   shop.ShopID,
			Name:     shop.Name,
			Cover:    cover,
			Priceout: shop.Priceout,
			Num:      carshop.Num})
	}

	return rcar, nil
}

func ListHistoryOrder(username string, pageindex, pagesize uint) (interface{}, error) {
	type FrontOrderShop struct {
		ShopID   string `json:"shopid"`
		Name     string `json:"name"`
		Cover    string `json:"cover"`
		Priceout uint   `json:"priceout"`
		Num      uint   `json:"num"`
	}
	type FrontOrder struct {
		OrderID    string           `json:"orderid"`
		Myshops    []FrontOrderShop `json:"myshops"`
		PriceTotal uint             `json:"pricetotal"`
		CreatedAt  time.Time        `json:"addtime"`
	}

	rinfos := struct {
		Total int          `json:"total"`
		Infos []FrontOrder `json:"infos"`
	}{0, make([]FrontOrder, 0)}
	if pageindex == 0 {
		pageindex = 1
	}
	if pagesize == 0 {
		pagesize = 20
	}
	index := (pageindex - 1) * pagesize

	var orders []Order = make([]Order, 0)
	if username == beego.AppConfig.String("Admin") {
		if err := db.Where("order_type = ?", OrderTypeForHistory).
			Order("created_at desc").
			Limit(pagesize).Offset(index).
			Find(&orders).Error; err != nil {
			beego.Error("list history order fail! err:", err)
			return rinfos, err
		}
	} else {
		if err := db.Where("order_type = ? AND user_name = ?", OrderTypeForHistory, username).
			Order("created_at desc").
			Limit(pagesize).Offset(index).
			Find(&orders).Error; err != nil {
			beego.Error("list history order fail! err:", err)
			return rinfos, err
		}
	}

	for _, order := range orders {
		frontorder := FrontOrder{
			OrderID:    order.OrderID,
			Myshops:    make([]FrontOrderShop, 0),
			PriceTotal: 0,
			CreatedAt:  order.CreatedAt}

		// fill my shops
		type HistoryOrderShop struct {
			ShopID string `json:"shopid"`
			Num    uint   `json:"num"`
		}
		ordershops := make([]HistoryOrderShop, 0)
		if err := json.Unmarshal([]byte(order.ShopIDs), &ordershops); err != nil {
			beego.Error("Unmarshal histroy shops fail! err:", err)
			return rinfos, err
		}
		for _, ordershop := range ordershops {
			// get shop
			var shop Shop
			if err := db.First(&shop, "shop_id = ?", ordershop.ShopID).Error; err != nil {
				beego.Error("get shop fail! err:", err)
				return rinfos, err
			}
			var cover string
			if len(shop.Icons) > 0 {
				cover = shop.Icons[0].Content
			}
			frontorder.PriceTotal = frontorder.PriceTotal + ordershop.Num*shop.Priceout
			frontorder.Myshops = append(frontorder.Myshops, FrontOrderShop{
				ShopID:   shop.ShopID,
				Name:     shop.Name,
				Cover:    cover,
				Priceout: shop.Priceout,
				Num:      ordershop.Num,
			})

			rinfos.Infos = append(rinfos.Infos, frontorder)
		}
	}

	return rinfos, nil
}

func PayOrModShoppingCar(username string, pay bool, carparams []byte) error {
	type FrontOrderShop struct {
		ShopID   string `json:"shopid"`
		Name     string `json:"name"`
		Cover    string `json:"cover"`
		Priceout uint   `json:"priceout"`
		Num      uint   `json:"num"`
	}
	type FrontOrder struct {
		OrderID    string           `json:"orderid"`
		Myshops    []FrontOrderShop `json:"myshops"`
		PriceTotal uint             `json:"pricetotal"`
		CreatedAt  time.Time        `json:"addtime"`
	}
	var car FrontOrder
	if err := json.Unmarshal(carparams, &car); err != nil {
		beego.Error(err)
		return err
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	var order Order
	if pay {
		for _, carshop := range car.Myshops {
			if err := ConsumeShops(tx, carshop.ShopID, int(carshop.Num)); err != nil {
				beego.Error(err)
				return err
			}
		}
		order.OrderType = OrderTypeForHistory
	} else {
		order.OrderType = OrderTypeForPay
	}
	order.OrderID = car.OrderID
	order.PriceTotal = car.PriceTotal
	order.UserName = username
	order.ShopIDs = ""

	type CarShop struct {
		ShopID string `json:"shopid"`
		Num    uint   `json:"num"`
	}
	var carshops []CarShop = make([]CarShop, 0)
	for _, shop := range car.Myshops {
		carshops = append(carshops, CarShop{ShopID: shop.ShopID, Num: shop.Num})
	}
	shopids, err := json.Marshal(carshops)
	if err != nil {
		beego.Error(err)
		return err
	}
	order.ShopIDs = string(shopids)

	if err = tx.Save(&order).Error; err != nil {
		if pay {
			beego.Error("pay fail! err:", err)
		} else {
			beego.Error("modify car fail! err:", err)
		}
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
