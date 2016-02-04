package controllers

import (
	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson"
	"ranbbService/util"
	"ranbbPoss/models"
)

type MainController struct {
	beego.Controller
}


func (this *MainController) Get() {
	this.Redirect("/ranbb/admin",302)
}


func (this * MainController) CheckSession() {
	admin := this.GetSession("userInfo")
	if admin == nil {
		beego.Debug("MainController-CheckSession",admin)
//		this.Redirect("/ranbb/admin/#/sessions/signin",301)
		this.Data["json"]=&Response{Code:-1,Data:nil}
		this.ServeJson()
		return;
	}
	beego.Debug("MainController-CheckSession",admin.(*model.Admin))
	this.Data["json"] = &Response{Code:1,Data:admin}
	this.ServeJson()
}

func (this * MainController) Login() {
	username := this.GetString("username")
	password := this.GetString("password")
	beego.Debug("login ==>",username,password)

	admin,err := model.GetAdminByUserName(username)
	if err != nil {
		this.Data["json"] = "登录名错误"
		this.ServeJson()
		return
	}
	this.SetSession("userInfo",admin);

	beego.Debug(util.StringMd5(password))

	if admin.Password != util.StringMd5(password) {
		this.Data["json"] = "密码错误"
		this.ServeJson()
		return
	}

	js := simplejson.New()
	js.Set("username",username)
	js.Set("password",password)
	this.Data["json"] = js
	this.ServeJson()
}

func (this * MainController) Logout() {
	this.DestroySession()
	this.Data["json"] = "OK"
	this.ServeJson()
}
