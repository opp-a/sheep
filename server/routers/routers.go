package routers

import (
	"sheep/controllers/chome"
	"sheep/controllers/clogin"
	"sheep/controllers/cordercrud"
	"sheep/controllers/cshopcrud"

	"github.com/astaxie/beego"
)

//前台路由
func front() {
	// for all
	beego.Router("/home/*", &chome.CHome{})
	beego.Router("/ticket", &chome.CHome{})
}

//后台路由
func back() {
	// 登录
	beego.Router("/api/login", &clogin.CLogin{}, "get,post:Login")
	beego.Router("/api/register", &clogin.CLogin{}, "post:Register")

	// shop crud
	beego.Router("/api/shop/ucreate", &cshopcrud.CShopCRUD{}, "post:UCreateShop")
	beego.Router("/api/shop/delete", &cshopcrud.CShopCRUD{}, "delete:DeleteShop")
	beego.Router("/api/shop/list", &cshopcrud.CShopCRUD{}, "post:QueryShops")

	// order
	beego.Router("/api/car/get", &cordercrud.COrderCRUD{}, "get:GetCar")
	beego.Router("/api/car/pay", &cordercrud.COrderCRUD{}, "post:PayOrder")
	beego.Router("/api/car/modify", &cordercrud.COrderCRUD{}, "post:ModifyCar")
	beego.Router("/api/historyorder/list", &cordercrud.COrderCRUD{}, "post:QueryHistoryOrders")

}
