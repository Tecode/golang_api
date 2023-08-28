package main

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"golang_apiv2/controllers"
	_ "golang_apiv2/routers"
	"golang_apiv2/utils"
)

func main() {
	// 开发环境构建swagger文档
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		// 后台管理
		//beego.BConfig.Listen.EnableAdmin = true
		//beego.BConfig.Listen.AdminAddr = "localhost"
		//beego.BConfig.Listen.AdminPort = 8088
	}
	// 日志模块
	if beego.BConfig.RunMode == "prod" {
		logs.Async()
		err01 := logs.SetLogger(
			logs.AdapterMultiFile,
			`{"filename":"logs/golang_api.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`,
		)
		if err01 != nil {
			logs.Error(err01.Error())
		}
		emailConfig := fmt.Sprintf(
			`{"username":"%s","password":"%s","host":"%s:587","sendTos":["283731869@qq.com"]}`,
			utils.GetAppConfigValue("emailaccount"),
			utils.GetAppConfigValue("emailpassword"),
			utils.GetAppConfigValue("emailhost"),
		)
		err02 := logs.SetLogger(logs.AdapterMail, emailConfig)
		if err02 != nil {
			logs.Error(err02.Error())
		}

	}
	utils.RunTimedTask()
	beego.SetStaticPath("/resource", "resource")
	beego.ErrorController(&controllers.ErrorControllerController{})
	// 现在IP的请求速率
	//beego.InsertFilter("*", beego.BeforeStatic, utils.RequestInterceptor)
	beego.InsertFilter("*", beego.BeforeRouter, utils.RequestInterceptor)
	beego.Run()
}
