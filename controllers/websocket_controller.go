package controllers

import (
	"flag"
	beego "github.com/beego/beego/v2/server/web"
	"golang_apiv2/websocket"
	"sync"
)

type WebSocketController struct {
	beego.Controller
}

type SocketSingleton struct {
	hub *websocket.Hub
}

var (
	instance *SocketSingleton
	once     sync.Once
)

// GetSocket 获取socket实例
func GetSocket() *SocketSingleton {
	return instance
}

var hub *websocket.Hub

func (c *WebSocketController) Get() {
	websocket.ServeWebsocket(hub, c.Ctx.ResponseWriter, c.Ctx.Request)
	c.TplName = "websocket.html"
}

func init() {
	flag.Parse()
	hub = websocket.NewHub()
	// 单例模式
	once.Do(func() {
		instance = &SocketSingleton{hub: hub}
	})
	go hub.Run()
}
