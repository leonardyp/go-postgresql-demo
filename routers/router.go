// @APIVersion 1.0.0
// @Title beego+postgresql test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact 1679550318@qq.com
// @TermsOfServiceUrl https://github.com/leonardyp
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"go-postgresql-demo/controllers"

	"github.com/astaxie/beego"
)

func init() {

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
