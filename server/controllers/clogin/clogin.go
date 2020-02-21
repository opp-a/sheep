package clogin

import (
	"encoding/json"
	"sheep/controllers/cbase"
	"sheep/models"
	"sheep/service/admin/user"

	"github.com/astaxie/beego"
)

type CLogin struct {
	cbase.BaseController
}

// login html
func (this *CLogin) Login() {
	userinfo := models.User{}

	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &userinfo); err != nil {
		beego.Error(err)
		this.ResponseJson(401, "服务器未收到登录信息", nil)
		this.StopRun()
	}

	beego.Info(userinfo)
	// Authenticate
	err := user.Authenticate(userinfo)
	if err != nil {
		beego.Error(err)
		this.ResponseJson(401, "账户认证失败", nil)
		this.StopRun()
	}

	// Get Token
	token, err := this.GetToken(userinfo.Name, this.Ctx.Input.Domain())
	if err != nil {
		beego.Error(err)
		this.ResponseJson(401, "获取签名失败", nil)
		this.StopRun()
	}
	this.ResponseJson(200, "", map[string]interface{}{"token": "Bearer " + token, "name": userinfo.Name})
}
