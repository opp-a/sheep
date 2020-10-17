package cordercrud

import (
	"encoding/json"
	"sheep/controllers/cbase"
	"sheep/models"

	"github.com/astaxie/beego"
)

type COrderCRUD struct {
	cbase.BaseController
}

func (this *COrderCRUD) GetCar() {
	carinfo, err := models.GetShoppingCar(this.UserName)
	if err != nil {
		beego.Error(err)
		this.ResponseJson(500, err.Error(), nil)
		this.StopRun()
	}
	this.ResponseJson(200, "", carinfo)
}

func (this *COrderCRUD) QueryHistoryOrders() {
	pageinfo := struct {
		PageIndex uint `json:"page"`
		PageSize  uint `json:"pagenumber"`
	}{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &pageinfo); err != nil {
		beego.Error(err)
		this.ResponseJson(400, err.Error(), nil)
		this.StopRun()
	}

	if rinfos, err := models.ListHistoryOrder(this.UserName, pageinfo.PageIndex, pageinfo.PageSize); err != nil {
		beego.Error(err)
		this.ResponseJson(500, err.Error(), nil)
		this.StopRun()
	} else {
		this.ResponseJson(200, "", rinfos)
	}
}

func (this *COrderCRUD) PayOrder() {
	if err := models.PayOrModShoppingCar(this.UserName, true, this.Ctx.Input.RequestBody); err != nil {
		beego.Error(err)
		this.ResponseJson(500, err.Error(), nil)
		this.StopRun()
	}

	this.ResponseJson(200, "", nil)
}

func (this *COrderCRUD) ModifyCar() {
	if err := models.PayOrModShoppingCar(this.UserName, false, this.Ctx.Input.RequestBody); err != nil {
		beego.Error(err)
		this.ResponseJson(500, err.Error(), nil)
		this.StopRun()
	}

	this.ResponseJson(200, "", nil)
}
