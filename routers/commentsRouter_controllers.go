package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:ClaseEntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:ClaseEntradaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:ClaseEntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:ClaseEntradaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:ClaseEntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:ClaseEntradaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:ClaseEntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:ClaseEntradaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:ClaseEntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:ClaseEntradaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:ElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:ElementoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:ElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:ElementoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:ElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:ElementoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:ElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:ElementoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:ElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:ElementoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:EntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:EntradaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:EntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:EntradaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:EntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:EntradaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:EntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:EntradaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:EntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:EntradaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:EstadoEntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:EstadoEntradaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:EstadoEntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:EstadoEntradaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:EstadoEntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:EstadoEntradaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:EstadoEntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:EstadoEntradaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:EstadoEntradaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:EstadoEntradaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:TipoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/arka_api_crud/controllers:TipoContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
