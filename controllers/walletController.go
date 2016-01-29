package controllers
import (
"github.com/astaxie/beego"
"ranbbPoss/models"
)

type WalletController struct  {
	beego.Controller
}

func (this * WalletController)checkSession() ( *model.Admin) {
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


func (this * WalletController)GetAllEnchashments()  {
	if v := this.checkSession();v == nil {
		return
	}

	enchashments,err :=  model.GetAllUnDealReq()
	if err != nil {
		this.Data["json"]=Response{-1,"获取提现列表失败"}
		this.ServeJson()
		return
	}
	this.Data["json"]=Response{1,enchashments}
	this.ServeJson()
	return
}

