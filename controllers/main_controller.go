package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
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

// SignIn 登录
func (c *MainController) SignIn() {
	c.TplName = "sign-in.html"
}

// SignUp 登录
func (c *MainController) SignUp() {
	c.TplName = "sign-up.html"
}
