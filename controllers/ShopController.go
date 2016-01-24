package controllers
import (
	"ranbbPoss/models"
	"github.com/astaxie/beego"
	"fmt"
	"time"
	"strconv"
	"strings"
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

	beego.Debug(this.Ctx.Input.Params)

	shopId := this.GetString("Shopid")
	shopName := this.GetString("Shopname")
	shopPhone := this.GetString("Mobile")
	shopEmail := this.GetString("Email")
	shopTaobaoAccount := this.GetString("Shoptaobaoid")
	memo := this.GetString("Memo")


	beego.Debug(shopId,shopName,shopPhone,shopEmail,shopTaobaoAccount,memo)
	id,err := strconv.Atoi(shopId)
	if err != nil {
		beego.Error(shopId,err.Error())
		this.Data["json"]=Response{-1,"不存在的商户"}
		this.ServeJson()
		return
	}
	shop,err := model.GetShopById(id)
	if err != nil || shop == nil{
		beego.Error(shop,err.Error(),shopId)
		this.Data["json"]=Response{-1,"不存在的商户"}
		this.ServeJson()
		return
	}

	if len(shopName) > 0 {
		shop.Shopname = shopName
	}
	if len(shopPhone) > 0 {
		shop.Mobile = shopPhone
	}
	if len(shopEmail) > 0 {
		shop.Email = shopEmail
	}
	if len(shopTaobaoAccount) > 0 {
		shop.Shoptaobaoid = shopTaobaoAccount
	}
	if len(memo) > 0 {
		shop.Memo = memo
	}

	shop.Updatetime = time.Now().Unix()



	_,fileHeader,err := this.GetFile("file")
	if err != nil {
		beego.Error(err.Error())
	}else {

		filetype := fileHeader.Header.Get("Content-Type")

		var imgPath string
		if strings.Contains(filetype,"image") {
			imgPath = fmt.Sprintf("./views/img/shop-%s-%d.%s",shopId,shop.Updatetime,filetype[6:])
			err = this.SaveToFile("file",imgPath)
			if err == nil {
				shop.ImageUrl = beego.AppConfig.String("localHost")+imgPath[8:]
			}
		}else{
			beego.Error("上传的不是图片文件")
			this.Data["json"]=Response{-1,"上传的不是图片文件"}
			this.ServeJson()
			return
		}
	}

	err = model.UpdateShopById(shop)
	if err != nil {
		beego.Error(err.Error())
	}
	this.Data["json"]=Response{1,"更新成功"}
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

func (this * ShopController)AddShop() {
	admin := this.checkSession();
	if admin == nil{
		this.Data["json"]=Response{-1,"登陆超时,请重新登陆"}
		this.ServeJson()
		return
	}

	beego.Debug(this.Ctx.Input.Params)

	shopName := this.GetString("Shopname")
	shopPhone := this.GetString("Mobile")
	shopEmail := this.GetString("Email")
	shopTaobaoAccount := this.GetString("Shoptaobaoid")
	memo := this.GetString("Memo")

	shop,err := model.GetShopByName(shopName)
	if err == nil&&shop!=nil {
		this.Data["json"]=Response{Code:-1,Data:"店铺名称已存在"}
		this.ServeJson()
		return
	}
	shop,err = model.GetShopByMobile(shopPhone)
	if err == nil&&shop!=nil {
		this.Data["json"]=Response{Code:-1,Data:"店铺手机号码已存在"}
		this.ServeJson()
		return
	}
	shop,err = model.GetShopByEmail(shopEmail)
	if err == nil&&shop!=nil {
		this.Data["json"]=Response{Code:-1,Data:"店铺邮箱已存在"}
		this.ServeJson()
		return
	}
	shop,err = model.GetShopByShoptaobaoid(shopTaobaoAccount)
	if err == nil&&shop!=nil {
		this.Data["json"]=Response{Code:-1,Data:"店铺淘宝账号已存在"}
		this.ServeJson()
		return
	}

	var newshop model.Shop
	if len(shopName) > 0 {
		newshop.Shopname = shopName
	}
	if len(shopPhone) > 0 {
		newshop.Mobile = shopPhone
	}
	if len(shopEmail) > 0 {
		newshop.Email = shopEmail
	}
	if len(shopTaobaoAccount) > 0 {
		newshop.Shoptaobaoid = shopTaobaoAccount
	}
	if len(memo) > 0 {
		newshop.Memo = memo
	}

	newshop.Updatetime = time.Now().Unix()
	newshop.Createtime = time.Now().Unix()

	shopId,err := model.AddShop(&newshop)
	if err != nil && shopId <= 0 {
		this.Data["json"]=Response{Code:-1,Data:"添加店铺失败"}
		this.ServeJson()
		return
	}

	_,fileHeader,err := this.GetFile("file")
	if err != nil {
		beego.Error(err.Error())
	}else {
		filetype := fileHeader.Header.Get("Content-Type")
		beego.Info(filetype)
		var imgPath string
		if strings.Contains(filetype,"image") {
			imgPath = fmt.Sprintf("./views/img/shop-%d-%d.%s",shopId,newshop.Updatetime,filetype[6:])
			err = this.SaveToFile("file",imgPath)
			if err == nil {
				newshop.ImageUrl = beego.AppConfig.String("localHost")+imgPath[8:]
			}
		}else{
			beego.Error("上传的不是图片文件")
			this.Data["json"]=Response{-1,"上传的不是图片文件"}
			this.ServeJson()
			return
		}
	}

	err = model.UpdateShopById(&newshop)
	if err != nil {
		beego.Error(err.Error())
	}

	this.Data["json"]=Response{Code:1,Data:"添加成功"}
	this.ServeJson()
	return
}