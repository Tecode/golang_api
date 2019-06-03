package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["golang_api/controllers:BookController"] = append(beego.GlobalControllerRouter["golang_api/controllers:BookController"],
        beego.ControllerComments{
            Method: "GetBookDetail",
            Router: `/detail/:book_id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["golang_api/controllers:BookController"] = append(beego.GlobalControllerRouter["golang_api/controllers:BookController"],
        beego.ControllerComments{
            Method: "GetNewBook",
            Router: `/new_book`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["golang_api/controllers:BookController"] = append(beego.GlobalControllerRouter["golang_api/controllers:BookController"],
        beego.ControllerComments{
            Method: "GetPopularBook",
            Router: `/popular_book`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["golang_api/controllers:BookController"] = append(beego.GlobalControllerRouter["golang_api/controllers:BookController"],
        beego.ControllerComments{
            Method: "GetRecommend",
            Router: `/recommend`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["golang_api/controllers:BookController"] = append(beego.GlobalControllerRouter["golang_api/controllers:BookController"],
        beego.ControllerComments{
            Method: "Today",
            Router: `/today`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["golang_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["golang_api/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["golang_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["golang_api/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["golang_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["golang_api/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["golang_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["golang_api/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["golang_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["golang_api/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["golang_api/controllers:PublicController"] = append(beego.GlobalControllerRouter["golang_api/controllers:PublicController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/picture/:img`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["golang_api/controllers:UserController"] = append(beego.GlobalControllerRouter["golang_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["golang_api/controllers:UserController"] = append(beego.GlobalControllerRouter["golang_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["golang_api/controllers:UserController"] = append(beego.GlobalControllerRouter["golang_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["golang_api/controllers:UserController"] = append(beego.GlobalControllerRouter["golang_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["golang_api/controllers:UserController"] = append(beego.GlobalControllerRouter["golang_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["golang_api/controllers:UserController"] = append(beego.GlobalControllerRouter["golang_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["golang_api/controllers:UserController"] = append(beego.GlobalControllerRouter["golang_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserInfo",
            Router: `/user_info`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
