package models

import (
	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
	"gopkg.in/pg.v3"
)

var db *pg.DB

func init() {
	db = pg.Connect(&pg.Options{
		User:     beego.AppConfig.String("postgre::user"),
		Database: beego.AppConfig.String("postgre::db"),
	})

}

func DB() *pg.DB {
	return db
}
