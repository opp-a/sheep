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
	shopinfo := models.Shop{}

	if err := this.ParseForm(&shopinfo); err != nil {
		beego.Error(err)
		this.ResponseJson(400, err.Error(), nil)
		this.StopRun()
	}

	beego.Debug(shopinfo)

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

	beego.Debug(pageinfo)

	if rinfos, err := models.ListShop(pageinfo.PageIndex, pageinfo.PageSize); err != nil {
		beego.Error(err)
		this.ResponseJson(500, err.Error(), nil)
		this.StopRun()
	} else {
		this.ResponseJson(200, "", rinfos)
	}
}
