package cshopcrud

import (
	"encoding/json"
	"sheep/controllers/cbase"
	"sheep/models"

	"github.com/astaxie/beego"
)

type CShopCRUD struct {
	cbase.BaseController
}

func (this *CShopCRUD) UCreateShop() {
	if this.UserName != beego.AppConfig.String("Admin") {
		beego.Error("this user be not permitted!")
		this.ResponseJson(416, "you are not permitted!", nil)
		this.StopRun()
	}
	shopinfo := models.Shop{}
	shopparam := struct {
		ShopID   string   `json:"shopid"`
		Name     string   `json:"name"` // string默认长度为255
		Icons    []string `json:"icons"`
		Pricein  uint     `json:"pricein"`
		Priceout uint     `json:"priceout"`
		Num      uint     `json:"num"`
		Desc     string   `json:"desc"`
	}{}

	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &shopparam); err != nil {
		beego.Error(err)
		this.ResponseJson(400, err.Error(), nil)
		this.StopRun()
	}
	shopinfo.ShopID = shopparam.ShopID
	shopinfo.Name = shopparam.Name
	shopinfo.Icons = make([]models.File, 0)
	for _, icon := range shopparam.Icons {
		file := models.File{}
		file.Content = icon
		shopinfo.Icons = append(shopinfo.Icons, file)
	}
	shopinfo.Pricein = shopparam.Pricein
	shopinfo.Priceout = shopparam.Priceout
	shopinfo.Num = int(shopparam.Num)
	shopinfo.Desc = shopparam.Desc
	if shopinfo.ShopID == "" {
		if err := models.AddShop([]models.Shop{shopinfo}); err != nil {
			beego.Error(err)
			this.ResponseJson(500, err.Error(), nil)
			this.StopRun()
		}
	} else {
		if err := models.UpdateShop(shopinfo.ShopID, shopinfo); err != nil {
			beego.Error(err)
			this.ResponseJson(500, err.Error(), nil)
			this.StopRun()
		}
	}
	this.ResponseJson(200, "", nil)
}

func (this *CShopCRUD) DeleteShop() {
	if this.UserName != beego.AppConfig.String("Admin") {
		beego.Error("this user be not permitted!")
		this.ResponseJson(416, "you are not permitted!", nil)
		this.StopRun()
	}
	shopid := struct {
		ShopID string `json:"shopid"`
	}{}

	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &shopid); err != nil {
		beego.Error(err)
		this.ResponseJson(400, err.Error(), nil)
		this.StopRun()
	}

	beego.Debug(shopid)

	if shopid.ShopID == "" {
		this.ResponseJson(200, "", nil)
		this.StopRun()
	} else {
		if err := models.DeleteShop([]string{shopid.ShopID}); err != nil {
			beego.Error(err)
			this.ResponseJson(500, err.Error(), nil)
			this.StopRun()
		}
	}
	this.ResponseJson(200, "", nil)
}

func (this *CShopCRUD) QueryShops() {
	pageinfo := struct {
		PageIndex uint `json:"page"`
		PageSize  uint `json:"pagenumber"`
	}{}

	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &pageinfo); err != nil {
		beego.Error(err)
		this.ResponseJson(400, err.Error(), nil)
		this.StopRun()
	}

	if rinfos, err := models.ListShop(this.UserName, pageinfo.PageIndex, pageinfo.PageSize); err != nil {
		beego.Error(err)
		this.ResponseJson(500, err.Error(), nil)
		this.StopRun()
	} else {
		this.ResponseJson(200, "", rinfos)
	}
}
