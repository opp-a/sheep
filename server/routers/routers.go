package routers

import (
	"sheep/controllers/chome"
	"sheep/controllers/clogin"

	"github.com/astaxie/beego"
)

//后台路由
func back() {
	// 登录
	beego.Router("/api/login", &clogin.CLogin{}, "get,post:Login")

}

//前台路由
func front() {
	// for all
	beego.Router("/*", &chome.CHome{})
}
