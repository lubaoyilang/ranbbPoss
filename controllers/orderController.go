package controllers
import (
	"github.com/astaxie/beego"
	"ranbbPoss/models"
)


type OrderController struct {
	beego.Controller
}

func (this * OrderController)GetNewOrders() {
	admin := this.checkSession()
	if admin == nil {
		this.ServeJson()
		return
	}

	orders,err := model.GetNewOrders()
	if err != nil {
		this.ServeJson()
		return
	}
	this.Data["json"]=orders
	this.ServeJson()
}

func (this * OrderController)checkSession() ( *model.Admin) {
	user := this.GetSession("userInfo")
	var admin *model.Admin
	if v,ok := user.(*model.Admin);!ok {
		beego.Debug("OrderController checkSession:",v,user);
		this.Redirect("/ranbb/admin/#/sessions/signin",301)
		return admin
	}else{
		admin = v;
	}

	return admin
}