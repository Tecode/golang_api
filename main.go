package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "golang_apiv2/routers"
)

func main() {
	// 开发环境构建swagger文档
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
