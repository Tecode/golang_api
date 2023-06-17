package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"golang_apiv2/controllers"
	"golang_apiv2/utils"
)

func init() {
	// http://localhost:8080/v1/haoxuan
	// 登录
	beego.CtrlPost("/api/login", (*controllers.UserRelatedController).UserLogin)
	nameSpace := beego.NewNamespace("/api/v1",
		beego.NSRouter("/haoxuan", &controllers.HaouxuanController{}),
		// 嵌套 namespace
		//beego.NSNamespace("/admin", beego.NSRouter("/user", uc)),
	)

	// 包含token验证，中间件过滤拦截
	nameSpace.Filter("before", utils.Interceptor)

	// http://localhost:8080/haoxuan
	beego.Router("/", &controllers.MainController{})
	beego.Router("/:haoxuan", &controllers.HaouxuanController{})
	beego.AddNamespace(nameSpace)
}
