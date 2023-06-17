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
	nameSpace := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/data", beego.NSInclude(&controllers.UserRelatedController{})),
		beego.NSNamespace("/user", beego.NSInclude(&controllers.HaouxuanController{})),
		// 嵌套 namespace
		//beego.NSNamespace("/admin", beego.NSRouter("/user", uc)),
	)

	// 包含token验证，中间件过滤拦截
	nameSpace.Filter("before", utils.Interceptor)

	// http://localhost:8080/haoxuan
	beego.Router("/", &controllers.MainController{})
	beego.AddNamespace(nameSpace)
}
