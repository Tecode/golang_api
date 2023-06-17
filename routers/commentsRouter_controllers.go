package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["golang_apiv2/controllers:HaouxuanController"] = append(beego.GlobalControllerRouter["golang_apiv2/controllers:HaouxuanController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["golang_apiv2/controllers:HaouxuanController"] = append(beego.GlobalControllerRouter["golang_apiv2/controllers:HaouxuanController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["golang_apiv2/controllers:HaouxuanController"] = append(beego.GlobalControllerRouter["golang_apiv2/controllers:HaouxuanController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["golang_apiv2/controllers:HaouxuanController"] = append(beego.GlobalControllerRouter["golang_apiv2/controllers:HaouxuanController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["golang_apiv2/controllers:HaouxuanController"] = append(beego.GlobalControllerRouter["golang_apiv2/controllers:HaouxuanController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
