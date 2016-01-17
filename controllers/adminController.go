package controllers
import (
	"github.com/astaxie/beego"
	"ranbbPoss/models"
	"ranbbService/util"
)

type Response struct {
	Code int
	Data interface{}
}


type AdminController struct {
	beego.Controller
}

func (this * AdminController)GetAdminList() {
	if admin:=this.checkSession();admin == nil {
		this.ServeJson()
		return;
	}
	page,err := this.GetInt("page",1);
	if err != nil {
		page = 1
	}
	size,err := this.GetInt("size",10);
	if err != nil{
		size = 10;
	}
	beego.Debug(page,size)
	admins,err := model.GetAdminListNotRoot(page,size)
	this.Data["json"]=admins
	this.ServeJson()
}

func  (this * AdminController) AddAdmin() {
	if admin:=this.checkSession();admin == nil {
		this.ServeJson()
		return;
	}else{
		if admin.Privilege != "root" {
			this.Data["json"]=Response{Code:-1,Data:"权限不足"}
			this.ServeJson()
			return;
		}
	}

	mobile := this.GetString("Mobile")
	passwd := this.GetString("Password")
	name := this.GetString("Name")
	priv := this.GetString("Privilege")
	admin := model.Admin{
		Mobile:mobile,
		Name:name,
		Privilege:priv,
		Password:passwd}
	beego.Debug(admin)

//	a,err := model.GetAdminByUserName(name)
//	if err != nil&&a != nil {
//		beego.Error(err.Error())
//		this.ServeJson()
//		return;
//	}
	admin.Uid = util.GetGuid()
	admin.Password = util.StringMd5(admin.Password)
	err := model.AddAdmin(&admin)
	if err != nil {
		beego.Error(err.Error())
		this.Data["json"]=Response{Code:-1,Data:"添加失败"}
		this.ServeJson()
		return;
	}
	this.Data["json"] = admin
	this.ServeJson()
}

func (this * AdminController)DeleteAdmin(){
	if admin:=this.checkSession();admin == nil {
		this.ServeJson()
		return;
	}else{
		if admin.Privilege!="root" {
			beego.Error(admin)
			this.Data["json"]=Response{Code:-1,Data:"权限不足"}
			this.ServeJson()
			return;
		}
	}

	mobile := this.GetString("mobile")
	beego.Debug(mobile)
	err := model.DeleteAdmin(mobile)
	if err != nil {
		beego.Error(err.Error())
		this.Data["json"]=Response{Code:-1,Data:"删除失败"}
		this.ServeJson()
		return;
	}
	this.Data["json"]=Response{Code:1,Data:"删除成功"}
	this.ServeJson()
}

func (this * AdminController)checkSession() ( *model.Admin) {

	user := this.GetSession("userInfo")
	beego.Info(user)
	var admin *model.Admin
	if v,ok := user.(*model.Admin);!ok {
		beego.Debug("AdminController checkSession:",v,user);
		this.Redirect("/ranbb/admin/#/sessions/signin",301)
		return admin
	}else{
		admin = v;
	}

	return admin
}

func (this *AdminController)UpdateAdmin() {
	if admin:=this.checkSession();admin == nil {
		this.ServeJson()
		return;
	}else{
		if admin.Privilege != "root" {
			this.Data["json"]=Response{Code:-1,Data:"权限不足"}
			this.ServeJson()
			return;
		}
	}
	mobile := this.GetString("Mobile")
	passwd := this.GetString("Password")
	name := this.GetString("Name")
	priv := this.GetString("Privilege")
	state,_ := this.GetInt("State")
	admin := model.Admin{
		Mobile:mobile,
		Name:name,
		Privilege:priv,
		Password:passwd,
		State:state}
	beego.Debug(admin)
	if len(passwd) > 0 {
		admin.Password = util.StringMd5(admin.Password)
	}
	err := model.UpdateAdmin(&admin)
	if err != nil {
		beego.Debug(err)
		this.Data["json"]=Response{Code:-1,Data:"更新失败"}
		this.ServeJson()
		return;
	}
	this.Data["json"]=Response{Code:1,Data:"更新成功"}
	this.ServeJson()
	return;
}