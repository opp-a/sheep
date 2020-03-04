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

// 注册
func (this *CLogin) Register() {
	if this.UserName != beego.AppConfig.String("Admin") {
		beego.Error("this user be not permitted!")
		this.ResponseJson(416, "you are not permitted!", nil)
		this.StopRun()
	}

	type FrontUser struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	var user FrontUser
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &user); err != nil {
		beego.Error(err)
		this.ResponseJson(500, "服务器未收到注册信息", nil)
		this.StopRun()
	}

	if err := models.AddUser(user.Name, user.Password); err != nil {
		beego.Error(err)
		this.ResponseJson(500, err.Error(), nil)
		this.StopRun()
	}
	this.ResponseJson(200, "", nil)
}
