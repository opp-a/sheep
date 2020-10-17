package chome

import (
	"io/ioutil"

	"github.com/astaxie/beego"
)

type CHome struct {
	beego.Controller
}

func (this *CHome) Get() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	indexhtml, err := ioutil.ReadFile("dist/index.html")
	if err != nil {
		this.Redirect("/", 302)
	} else {
		this.Ctx.WriteString(string(indexhtml))
	}
}
