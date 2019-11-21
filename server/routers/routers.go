package routers

import (
	"sheep/controllers/admincontrollers"
	"sheep/controllers/homecontrollers"

	"github.com/astaxie/beego"
)

//前台路由
func front() {
	beego.Router("/", &homecontrollers.IndexController{})
}

//后台路由
func back() {
	beego.Router("/admin/login", &admincontrollers.AdminController{}, "get,post:Login")
}
