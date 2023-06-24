package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"golang_apiv2/controllers"
	"golang_apiv2/utils"
)

func init() {
	// http://localhost:8080/v1/haoxuan
	// 登录
	//beego.CtrlPost("/api/login", (*controllers.UserRelatedController).UserLogin)
	// 注解路由
	nameSpace := beego.NewNamespace("/api",
		beego.NSNamespace("/v1",
			beego.NSInclude(&controllers.UserRelatedController{}),
			beego.NSInclude(&controllers.CommonController{}),
		),
		beego.NSNamespace("/v1/user", beego.NSInclude(&controllers.HaouxuanController{})),
	)

	// 包含token验证，中间件过滤拦截
	//nameSpace.Filter("before", utils.Interceptor)
	// 包含token验证，中间件过滤拦截
	beego.InsertFilter("/api/v1/user/*", beego.BeforeRouter, utils.Interceptor)

	// http://localhost:8080/contact
	// 模板引擎页面
	beego.Router("/", &controllers.MainController{})
	beego.CtrlGet("/contact", (*controllers.MainController).Contact)
	beego.CtrlGet("/news", (*controllers.MainController).News)
	beego.CtrlGet("/news-detail", (*controllers.MainController).NewsDetail)
	beego.Get("/.well-known/pki-validation/1BA546375A5C77B851DB153B5516819B.txt", controllers.CertificateText)
	beego.AddNamespace(nameSpace)
}
