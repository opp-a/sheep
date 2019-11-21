package homecontrollers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	this.Redirect("/static/index.html", 302)
}
