package main

import (
	_ "go-postgresql-demo/docs"
	_ "go-postgresql-demo/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/api"] = "swagger/api"
	}
	beego.Run()
}
