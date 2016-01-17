package controllers
import (
	"ranbbPoss/models"
	"github.com/astaxie/beego"
)


type ShopController struct {
	beego.Controller
}

func (this *ShopController)GetNewShops() {
	admin := this.checkSession();
	if admin == nil{
		this.ServeJson()
	}

	shops,err := model.GetNewShop()
	if err != nil {
		this.ServeJson()
		return;
	}
	this.Data["json"] = shops;
	this.ServeJson()
}

func (this * ShopController)checkSession() ( *model.Admin) {

	user := this.GetSession("userInfo")
	var admin *model.Admin
	if v,ok := user.(*model.Admin);!ok {
		beego.Debug("ShopController  checkSession",v,user);
		this.Redirect("/ranbb/admin/#/sessions/signin",301)
		return admin
	}else{
		admin = v;
	}

	return admin
}