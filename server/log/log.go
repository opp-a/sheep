package log

import (
	"fmt"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var (
	//develop mode
	Debug = true
)

//日志初始化
func init() {

	os.Chdir("/opt/sheep")
	//beego.LoadAppConfig("ini", "conf/app.conf")
	Debug = beego.AppConfig.String("runmode") == "dev"

	//创建日志目录
	if _, err := os.Stat("logs"); err != nil {
		os.Mkdir("logs", os.ModePerm)
	}

	logs.Async(1e3) //异步输出允许设置缓冲 chan 的大小
	//初始化日志各种配置
	maxsize := 1 << 24
	LogsConf := fmt.Sprintf(`{"filename":"logs/sus.log","level":7,"maxlines":0,"maxsize":%v,"daily":true,"maxdays":7,"color":true}`, maxsize)
	logs.SetLogger(logs.AdapterFile, LogsConf)
	if Debug {
		logs.SetLogger("console")
		beego.Info("日志配置信息：" + LogsConf)
	} else {
		//是否异步输出日志
		logs.Async(1e3)
	}
	logs.EnableFuncCallDepth(true) //是否显示文件和行号
}
