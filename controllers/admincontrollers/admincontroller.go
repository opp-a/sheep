package admincontrollers

import (
	"encoding/json"
	"sheep/controllers/commoncontrollers"
	"sheep/models"
	"sheep/service/admin/user"

	"github.com/astaxie/beego"
)

type AdminController struct {
	commoncontrollers.BaseController
}

// login html
func (this *AdminController) Login() {
	userinfo := models.User{}

	if err := this.ParseForm(&userinfo); err != nil {
		if err = json.Unmarshal(this.Ctx.Input.RequestBody, &userinfo); err != nil {
			beego.Error(err)
			this.ResponseAuth(1001, err.Error(), nil)
			this.StopRun()
		}
	}

	// Authenticate
	err := user.Authenticate(userinfo)
	if err != nil {
		beego.Error(err)
		this.ResponseAuth(1001, err.Error(), nil)
		this.StopRun()
	}

	// Get Token
	token, err := this.GetToken(userinfo.Name, this.Ctx.Input.Domain())
	if err != nil {
		beego.Error(err)
		this.ResponseAuth(1001, err.Error(), nil)
		this.StopRun()
	}
	this.ResponseJson(true, "", map[string]interface{}{"token": "Bearer " + token})
	this.StopRun()
}
