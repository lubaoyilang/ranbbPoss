package main

import (
	"github.com/astaxie/beego"
	_ "ranbbPoss/routers"
	_ "github.com/astaxie/beego/session/mysql"
	"github.com/astaxie/beego/orm"
//	"github.com/astaxie/beego/context"
//	"ranbbPoss/models"
)



func init() {
	beego.DirectoryIndex = true
	beego.SessionOn = true
	beego.SessionCookieLifeTime = 6000000
	beego.StaticDir["/ranbb/admin"] = "views"
	beego.SessionGCMaxLifetime = 60000000
	beego.SessionProvider = "mysql"
	beego.SessionSavePath =  beego.AppConfig.String("mydb")

	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.Debug=true

	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mydb"))
}

func main() {
	beego.Run()
}
