package main

import (
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
	utils.RunTimedTask()
	beego.SetStaticPath("/resource", "resource")
	beego.ErrorController(&controllers.ErrorControllerController{})
	// 现在IP的请求速率
	//beego.InsertFilter("*", beego.BeforeStatic, utils.RequestInterceptor)
	beego.InsertFilter("*", beego.BeforeRouter, utils.RequestInterceptor)
	beego.Run()
}
