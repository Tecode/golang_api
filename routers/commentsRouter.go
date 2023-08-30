package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["golang_apiv2/controllers:CommonController"] = append(beego.GlobalControllerRouter["golang_apiv2/controllers:CommonController"],
		beego.ControllerComments{
			Method:           "ImageResize",
			Router:           `/picture/:img`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["golang_apiv2/controllers:CommonController"] = append(beego.GlobalControllerRouter["golang_apiv2/controllers:CommonController"],
		beego.ControllerComments{
			Method:           "Upload",
			Router:           `/upload`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["golang_apiv2/controllers:ErrorControllerController"] = append(beego.GlobalControllerRouter["golang_apiv2/controllers:ErrorControllerController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["golang_apiv2/controllers:ErrorControllerController"] = append(beego.GlobalControllerRouter["golang_apiv2/controllers:ErrorControllerController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["golang_apiv2/controllers:ErrorControllerController"] = append(beego.GlobalControllerRouter["golang_apiv2/controllers:ErrorControllerController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["golang_apiv2/controllers:ErrorControllerController"] = append(beego.GlobalControllerRouter["golang_apiv2/controllers:ErrorControllerController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["golang_apiv2/controllers:UserControllerController"] = append(beego.GlobalControllerRouter["golang_apiv2/controllers:UserControllerController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["golang_apiv2/controllers:UserControllerController"] = append(beego.GlobalControllerRouter["golang_apiv2/controllers:UserControllerController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["golang_apiv2/controllers:UserControllerController"] = append(beego.GlobalControllerRouter["golang_apiv2/controllers:UserControllerController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["golang_apiv2/controllers:UserControllerController"] = append(beego.GlobalControllerRouter["golang_apiv2/controllers:UserControllerController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["golang_apiv2/controllers:UserControllerController"] = append(beego.GlobalControllerRouter["golang_apiv2/controllers:UserControllerController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["golang_apiv2/controllers:UserRelatedController"] = append(beego.GlobalControllerRouter["golang_apiv2/controllers:UserRelatedController"],
		beego.ControllerComments{
			Method:           "UserLogin",
			Router:           `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["golang_apiv2/controllers:UserRelatedController"] = append(beego.GlobalControllerRouter["golang_apiv2/controllers:UserRelatedController"],
		beego.ControllerComments{
			Method:           "RecordAccount",
			Router:           `/record-account`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["golang_apiv2/controllers:UserRelatedController"] = append(beego.GlobalControllerRouter["golang_apiv2/controllers:UserRelatedController"],
		beego.ControllerComments{
			Method:           "Register",
			Router:           `/register`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["golang_apiv2/controllers:UserRelatedController"] = append(beego.GlobalControllerRouter["golang_apiv2/controllers:UserRelatedController"],
		beego.ControllerComments{
			Method:           "ResetPassword",
			Router:           `/reset-password`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["golang_apiv2/controllers:UserRelatedController"] = append(beego.GlobalControllerRouter["golang_apiv2/controllers:UserRelatedController"],
		beego.ControllerComments{
			Method:           "SendEmailCode",
			Router:           `/send-code`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
