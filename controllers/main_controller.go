package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"golang_apiv2/utils"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	id, err := utils.TokenToUserId(c.Ctx.GetCookie("token"))
	if err != nil {
		logs.Error(err)
	}
	logs.Info(id)
	c.Data["Active"] = "product"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}

// CertificateText 提供Https文件进行校验
func CertificateText(ctx *context.Context) {
	ctx.Output.Download("./resource/certificate/1BA546375A5C77B851DB153B5516819B.txt")
}

// Contact 联系我们
func (c *MainController) Contact() {
	c.Data["Active"] = "contact"
	c.TplName = "contact.html"
}

// News 新闻资讯
func (c *MainController) News() {
	c.Data["Active"] = "news"
	c.TplName = "news.html"
}

// NewsDetail 新闻资讯详情
func (c *MainController) NewsDetail() {
	c.Data["Active"] = "news"
	c.TplName = "news-detail.html"
}

// Services 服务
func (c *MainController) Services() {
	c.Data["Active"] = "services"
	c.TplName = "services.html"
}

// ServicesDetail 服务详情
func (c *MainController) ServicesDetail() {
	c.Data["Active"] = "services"
	c.TplName = "service-detail.html"
}

// About 关于我们
func (c *MainController) About() {
	c.TplName = "about.html"
}

// Chats 在线聊天
func (c *MainController) Chats() {
	c.Data["Active"] = "chats"
	c.TplName = "chats.html"
}

// SignIn 登录
func (c *MainController) SignIn() {
	c.Ctx.SetCookie("token", "Hour测试cookie set token", 30)
	//c.Redirect()
	c.TplName = "sign-in.html"
}

// SignUp 注册
func (c *MainController) SignUp() {
	c.TplName = "sign-up.html"
}

// ForgotPassword 忘记密码
func (c *MainController) ForgotPassword() {
	c.Data["Type"] = c.Ctx.Input.Query("type")
	fmt.Println(c.Ctx.Input.Query("type"))
	c.TplName = "forgot-password.html"
}
