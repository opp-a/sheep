package main

import (
	"flag"
	_ "sheep/log"
	_ "sheep/models"
	_ "sheep/routers"

	"github.com/astaxie/beego"
)

//初始化函数
func init() {
	flag.Parse()
	beego.Info("")
	beego.Info("-----------------------------------------Welcome Little Sheep-----------------------------------------")
	beego.Info("")
	//sitemap静态目录
	beego.SetStaticPath("/static", "static")
	beego.Info("==程序启动完毕==")
}

func main() {
	/*
		if Debug {
			beego.BConfig.WebConfig.DirectoryIndex = true
			beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		} else {
			beego.BConfig.Listen.HTTPPort = beego.AppConfig.Int("port")
		}*/

	//定义错误和异常处理控制器
	beego.Run()
}
