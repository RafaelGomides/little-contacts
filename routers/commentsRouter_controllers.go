package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["little-contacts/controllers:ContactController"] = append(beego.GlobalControllerRouter["little-contacts/controllers:ContactController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["little-contacts/controllers:ContactController"] = append(beego.GlobalControllerRouter["little-contacts/controllers:ContactController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["little-contacts/controllers:ContactController"] = append(beego.GlobalControllerRouter["little-contacts/controllers:ContactController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:ID`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["little-contacts/controllers:ContactController"] = append(beego.GlobalControllerRouter["little-contacts/controllers:ContactController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:ID`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["little-contacts/controllers:ContactController"] = append(beego.GlobalControllerRouter["little-contacts/controllers:ContactController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:ID`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["little-contacts/controllers:ContactController"] = append(beego.GlobalControllerRouter["little-contacts/controllers:ContactController"],
        beego.ControllerComments{
            Method: "SendEmail",
            Router: `/email`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
