// @APIVersion 1.0.0
// @Title API Interface
// @Description beego 实现的API方法
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"golang_api/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/book_info",
			beego.NSInclude(
				&controllers.BookController{},
			),
		),
	)
	beego.AddNamespace(ns)
	// 包含token验证，中间件过滤拦截
	beego.InsertFilter("*/", beego.BeforeRouter, controllers.Filter)
}
