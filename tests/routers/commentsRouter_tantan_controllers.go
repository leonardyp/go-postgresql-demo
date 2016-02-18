package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["tantan/controllers:UserController"] = append(beego.GlobalControllerRouter["tantan/controllers:UserController"],
		beego.ControllerComments{
			"CreateUserAction",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["tantan/controllers:UserController"] = append(beego.GlobalControllerRouter["tantan/controllers:UserController"],
		beego.ControllerComments{
			"AllUserAction",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["tantan/controllers:UserController"] = append(beego.GlobalControllerRouter["tantan/controllers:UserController"],
		beego.ControllerComments{
			"RelationshipAction",
			`/:user_id/relationships/:other_user_id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["tantan/controllers:UserController"] = append(beego.GlobalControllerRouter["tantan/controllers:UserController"],
		beego.ControllerComments{
			"UserAllRelationshipsAction",
			`/:user_id/relationships`,
			[]string{"get"},
			nil})

}
