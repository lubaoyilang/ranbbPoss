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
		return
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

func (this * ShopController)UpdateShop() {
	admin := this.checkSession();
	if admin == nil{
		this.Data["json"]=Response{-1,"登陆超时,请重新登陆"}
		this.ServeJson()
		return
	}
	_,_,err := this.GetFile("file")
	if err != nil {
		beego.Error(err.Error())
	}
	this.SaveToFile("file","file.png")
	this.ServeJson()
}

func (this * ShopController)GetShopList() {
	admin := this.checkSession();
	if admin == nil{
		this.Data["json"]=Response{-1,"登陆超时,请重新登陆"}
		this.ServeJson()
		return
	}
	page,_ := this.GetInt("page")
	size,_ := this.GetInt("size")

	shops,err := model.GetShopByPage(page,size)
	if err != nil {
		this.Data["json"]=Response{-1,"查询失败"}
		this.ServeJson()
		return;
	}
	this.Data["json"]=Response{Code:1,Data:shops}
	this.ServeJson()
	return
}