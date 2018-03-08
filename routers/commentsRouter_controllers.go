package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:ParticipanteSesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:ParticipanteSesionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:ParticipanteSesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:ParticipanteSesionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:ParticipanteSesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:ParticipanteSesionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:ParticipanteSesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:ParticipanteSesionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:ParticipanteSesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:ParticipanteSesionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:RelacionSesionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:RelacionSesionesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:RelacionSesionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:RelacionSesionesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:RelacionSesionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:RelacionSesionesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:RelacionSesionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:RelacionSesionesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:RelacionSesionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:RelacionSesionesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:RolParticipanteSesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:RolParticipanteSesionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:RolParticipanteSesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:RolParticipanteSesionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:RolParticipanteSesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:RolParticipanteSesionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:RolParticipanteSesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:RolParticipanteSesionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:RolParticipanteSesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:RolParticipanteSesionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:SesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:SesionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:SesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:SesionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:SesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:SesionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:SesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:SesionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:SesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:SesionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:TipoSesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:TipoSesionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:TipoSesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:TipoSesionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:TipoSesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:TipoSesionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:TipoSesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:TipoSesionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:TipoSesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sesionesCrud/controllers:TipoSesionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
