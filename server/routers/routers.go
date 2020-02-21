package routers

import (
	"sheep/controllers/chome"
	"sheep/controllers/clogin"
	"sheep/controllers/cshopcrud"

	"github.com/astaxie/beego"
)

//前台路由
func front() {
	// for all
	beego.Router("/home/*", &chome.CHome{})
}

//后台路由
func back() {
	// 登录
	beego.Router("/api/login", &clogin.CLogin{}, "get,post:Login")

	// shop crud
	beego.Router("/api/shop/ucreate", &cshopcrud.CShopCRUD{}, "post:UCreateShop")
	beego.Router("/api/shop/delete", &cshopcrud.CShopCRUD{}, "delete:DeleteShop")
	beego.Router("/api/shop/list", &cshopcrud.CShopCRUD{}, "get:QueryShops")

}
