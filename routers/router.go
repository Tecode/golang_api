package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"golang_apiv2/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
