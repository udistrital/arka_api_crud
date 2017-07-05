// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/arka_api_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/tipo_contrato",
			beego.NSInclude(
				&controllers.TipoContratoController{},
			),
		),

    beego.NSNamespace("/clase_entrada",
        beego.NSInclude(
            &controllers.ClaseEntradaController{},
        ),
    ),

    beego.NSNamespace("/elemento",
        beego.NSInclude(
            &controllers.ElementoController{},
        ),
    ),

    beego.NSNamespace("/entrada",
        beego.NSInclude(
            &controllers.EntradaController{},
        ),
    ),

		beego.NSNamespace("/estado_entrada",
        beego.NSInclude(
            &controllers.EntradaController{},
        ),
    ),
	)
	beego.AddNamespace(ns)
}
