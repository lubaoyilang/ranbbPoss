package main

import (
	"github.com/astaxie/beego"
	_ "ranbbPoss/routers"
	_ "github.com/astaxie/beego/session/mysql"
	"github.com/astaxie/beego/orm"
)



func init() {
	beego.DirectoryIndex = true
	beego.SessionOn = true
	beego.SessionCookieLifeTime = 6000000
	beego.StaticDir["/ranbb/admin"] = "views"
	beego.SessionGCMaxLifetime = 60000000
	beego.SessionProvider = "mysql"
	beego.SessionSavePath =  "cloudbridge:Cbcnspsp06@tcp(115.29.164.59:3306)/storedb?charset=utf8"
//	beego.SessionDomain="http://127.0.0.1:8082/ranbb/admin/#/sessions/singin"

	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.Debug=true

	orm.RegisterDataBase("default", "mysql", "cloudbridge:Cbcnspsp06@tcp(115.29.164.59:3306)/storedb?charset=utf8")
}

func main() {
	beego.Run()
}
