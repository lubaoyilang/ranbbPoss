package main

import (
	"github.com/astaxie/beego"
	_ "ranbbPoss/routers"
	_ "github.com/astaxie/beego/session/mysql"
)



func init() {
	beego.DirectoryIndex = true
	beego.SessionOn = true
	beego.SessionAutoSetCookie = false
	beego.SessionCookieLifeTime = 60 * 5
	beego.StaticDir["/ranbb/admin"] = "views"
	beego.SessionGCMaxLifetime = 60 * 5
	beego.SessionProvider = "mysql"
	beego.SessionSavePath =  "cloudbridge:Cbcnspsp06@tcp(115.29.164.59:3306)/storedb?charset=utf8"
}

func main() {
	beego.Run()
}


