package routers

import (
	"ranbbPoss/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/ranbb/admin", &controllers.MainController{})
}
