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
	nameSpace := beego.NewNamespace("/api", beego.NSNamespace("/v1",
		beego.NSInclude(&controllers.UserRelatedController{}),
		beego.NSInclude(&controllers.CommonController{}),
		beego.NSInclude(&controllers.QualityWorkController{}),
	),
		beego.NSNamespace("/v1/user", beego.NSInclude(&controllers.UserController{})),
		beego.NSNamespace("/v1/mock", beego.NSInclude(&controllers.MockController{})),
	)

	// 包含token验证，中间件过滤拦截
	//nameSpace.Filter("before", utils.Interceptor)
	// 包含token验证，中间件过滤拦截
	beego.InsertFilter("/api/v1/user/*", beego.BeforeRouter, utils.Interceptor)
	//beego.InsertFilter("/api/v1/mock/*", beego.BeforeRouter, utils.Interceptor)

	// http://localhost:8080/contact
	// 模板引擎页面
	beego.Router("/", &controllers.MainController{})
	beego.CtrlGet("/contact", (*controllers.MainController).Contact)
	beego.CtrlGet("/about", (*controllers.MainController).Contact)
	beego.CtrlGet("/news", (*controllers.MainController).News)
	beego.CtrlGet("/news-detail", (*controllers.MainController).NewsDetail)
	beego.CtrlGet("/service", (*controllers.MainController).Services)
	beego.CtrlGet("/service-detail", (*controllers.MainController).ServicesDetail)
	beego.CtrlGet("/chats", (*controllers.MainController).Chats)
	beego.CtrlGet("/sign-in", (*controllers.MainController).SignIn)
	beego.CtrlGet("/sign-up", (*controllers.MainController).SignUp)
	beego.CtrlGet("/forgot-password", (*controllers.MainController).ForgotPassword)
	beego.CtrlGet("/quality-work", (*controllers.QualityWorkController).QualityWork)
	beego.Router("/ws", &controllers.WebSocketController{})
	beego.Get("/.well-known/pki-validation/1BA546375A5C77B851DB153B5516819B.txt", controllers.CertificateText)
	beego.Any("*", controllers.MockAnyMethod)
	beego.AddNamespace(nameSpace)
}
