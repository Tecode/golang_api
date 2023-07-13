package main

import (
	beego "github.com/beego/beego/v2/server/web"
	"golang_apiv2/controllers"
	_ "golang_apiv2/routers"
	"golang_apiv2/websocket"
)

func main() {
	// 开发环境构建swagger文档
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	websocket.Run()
	beego.SetStaticPath("/resource", "resource")
	beego.ErrorController(&controllers.ErrorControllerController{})
	beego.Run()
}
