package controllers

import (
	"flag"
	beego "github.com/beego/beego/v2/server/web"
	"golang_apiv2/websocket"
)

type WebSocketController struct {
	beego.Controller
}

var hub *websocket.Hub

func (c *WebSocketController) Get() {
	websocket.ServeWebsocket(hub, c.Ctx.ResponseWriter, c.Ctx.Request)
	c.TplName = "websocket.html"
}

func init() {
	flag.Parse()
	hub = websocket.NewHub()
	go hub.Run()
}
