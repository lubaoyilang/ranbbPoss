package controllers
import (
	"github.com/astaxie/beego"
	"ranbbPoss/models"
	"time"
	"strconv"
	"strings"
	"fmt"
)


type GoodsController struct {
	beego.Controller
}

func (this * GoodsController)GetGoodsList() {
	admin := this.checkSession();
	if admin == nil{
		this.Data["json"]=Response{-1,"登陆超时,请重新登陆"}
		this.ServeJson()
		return
	}
	shopId,err := this.GetInt("shopId",0)
	if err != nil||shopId <= 0 {
		this.Data["json"]=Response{-1,"错误的店铺"}
		this.ServeJson()
		return
	}
	goodsList,err := model.GetGoodsByShopId(shopId)
	if err != nil {
		this.Data["json"]=Response{-1,"获取列表失败"}
		this.ServeJson()
		return
	}

	this.Data["json"]=Response{1,goodsList}
	this.ServeJson()
	return
}

func (this * GoodsController)AddGoods() {
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

func (this * GoodsController)UpdateGoods() {

	admin := this.checkSession();
	if admin == nil{
		this.Data["json"]=Response{-1,"登陆超时,请重新登陆"}
		this.ServeJson()
		return
	}

	beego.Debug(this.Input())

	Goodid,err:= this.GetInt("Goodid",0)
	if err != nil||Goodid <= 0 {
		beego.Error(Goodid)
		this.Data["json"]=Response{-1,"不存在的商品"}
		this.ServeJson()
		return
	}
	id := Goodid
	GoodsName := this.GetString("GoodsName")
	Requirelevel := this.GetString("Requirelevel")
	Brokerage := this.GetString("Brokerage")
	SettingPrice := this.GetString("SettingPrice")
	State := this.GetString("State")
	memo := this.GetString("Memo")

//	id,err := strconv.Atoi(Goodid)
//	if err != nil {
//		beego.Error(Goodid,err.Error())
//		this.Data["json"]=Response{-1,"不存在的商品"}
//		this.ServeJson()
//		return
//	}

	level,err := strconv.Atoi(Requirelevel)
	if err != nil {
		beego.Error(Requirelevel,err.Error())
		this.Data["json"]=Response{-1,"错误的等级要求"}
		this.ServeJson()
		return
	}
	brokerage,err := strconv.Atoi(Brokerage)
	if err != nil {
		beego.Error(Brokerage,err.Error())
		this.Data["json"]=Response{-1,"错误的佣金数量"}
		this.ServeJson()
		return
	}

	settingPrice,err := strconv.Atoi(SettingPrice)
	if err != nil {
		beego.Error(SettingPrice,err.Error())
		this.Data["json"]=Response{-1,"错误的价格数量"}
		this.ServeJson()
		return
	}

	state,err := strconv.Atoi(State)
	if err != nil {
		beego.Error(State,err.Error())
		this.Data["json"]=Response{-1,"错误的状态"}
		this.ServeJson()
		return
	}


	goods,err := model.GetGoodsById(id)
	if err != nil || goods == nil{
		beego.Error(goods,err.Error(),goods)
		this.Data["json"]=Response{-1,"不存在的商品"}
		this.ServeJson()
		return
	}

	if len(GoodsName) > 0 {
		goods.GoodsName = GoodsName
	}
	if level >= 0 {
		goods.Requirelevel = level
	}
	if brokerage > 0 {
		goods.Brokerage = int64(brokerage)
	}
	if settingPrice > 0 {
		goods.SettingPrice = int64(settingPrice)
	}
	if state ==0||state == 1 {
		goods.State = state
	}
	if len(memo) > 0 {
		goods.Memo = memo
	}

	goods.Updatetime = time.Now().Unix()

	_,fileHeader,err := this.GetFile("file")
	if err != nil {
		beego.Error(err.Error())
	}else {

		filetype := fileHeader.Header.Get("Content-Type")

		var imgPath string
		if strings.Contains(filetype,"image") {
			imgPath = fmt.Sprintf("./views/img/goods-%d-%d.%s",goods.Goodid,goods.Updatetime,filetype[6:])
			err = this.SaveToFile("file",imgPath)
			if err == nil {
				goods.Imageurl = beego.AppConfig.String("localHost")+imgPath[8:]
			}
		}else{
			beego.Error("上传的不是图片文件")
			this.Data["json"]=Response{-1,"上传的不是图片文件"}
			this.ServeJson()
			return
		}
	}

	err = model.UpdateGoods(goods)
	if err != nil {
		beego.Error(err.Error())
	}
	this.Data["json"]=Response{1,"更新成功"}
	this.ServeJson()
}


func (this * GoodsController)checkSession() ( *model.Admin) {

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
