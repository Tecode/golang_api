package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.vip"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}

// CertificateText 提供Https文件进行校验
func CertificateText(ctx *context.Context) {
	ctx.Output.Download("./resource/certificate/1BA546375A5C77B851DB153B5516819B.txt")
}

// Contact 联系我们
func (c *MainController) Contact() {
	c.TplName = "contact.html"
}

// News 新闻资讯
func (c *MainController) News() {
	c.TplName = "news.html"
}

// NewsDetail 新闻资讯详情
func (c *MainController) NewsDetail() {
	c.TplName = "news-detail.html"
}
