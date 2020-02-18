package routers

import (
	"sheep/controllers/clogin"

	"github.com/astaxie/beego"
)

//前台路由
func front() {
	// beego.Router("/", &chome.IndexController{})
}

//后台路由
func back() {
	// 登录
	beego.Router("/api/login", &clogin.CLogin{}, "get,post:Login")

}
