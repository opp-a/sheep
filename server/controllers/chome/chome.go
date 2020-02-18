package chome

import (
	"github.com/astaxie/beego"
)

type CHome struct {
	beego.Controller
}

func (this *CHome) Get() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	this.Redirect("/", 302)
}
