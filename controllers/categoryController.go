package controllers
import (
	"github.com/astaxie/beego"
	"ranbbPoss/models"
)

type GoodsCategoryController struct {
	beego.Controller
}


func (this * GoodsCategoryController)GetGoodsCategorys() {
	admin := this.checkSession();
	if admin == nil{
		this.Data["json"]=Response{-1,"登陆超时,请重新登陆"}
		this.ServeJson()
		return
	}
	goodsId,err := this.GetInt("goodsId",0)
	if err != nil ||goodsId <= 0{
		this.Data["json"]=Response{-1,"未找到相应的商品"}
		this.ServeJson()
		return
	}

	goods,err := model.GetGoodsById(goodsId)
	if err != nil || goods == nil{
		beego.Error(err.Error(),goodsId)
		this.Data["json"]=Response{-1,"未找到相应的商品"}
		this.ServeJson()
		return
	}

	categorys,err := model.GetAllCategorysByGoodsId(goodsId)
	if err != nil {
		beego.Error(err.Error(),goodsId)
		this.Data["json"]=Response{-1,"查找商品分类失败"}
		this.ServeJson()
		return
	}
	this.Data["json"] = Response{1,categorys}
	this.ServeJson()
}

func (this * GoodsCategoryController)AddGoodsCategory() {
	admin := this.checkSession();
	if admin == nil{
		this.Data["json"]=Response{-1,"登陆超时,请重新登陆"}
		this.ServeJson()
		return
	}
	var category model.GoodsCategroy
	err := this.ParseForm(&category)
	if err != nil {
		beego.Error(err.Error())
		this.Data["json"] = Response{-1,"输入信息有误"}
		this.ServeJson()
	}

	if category.GoodsId <= 0 && category.ShopId <= 0{
		beego.Error(err.Error())
		this.Data["json"] = Response{-1,"输入信息有误"}
		this.ServeJson()
	}

	_,err = model.AddCategory(&category)
	if err != nil {
		beego.Error(err.Error())
		this.Data["json"] = Response{-1,"添加失败"}
		this.ServeJson()
	}
	beego.Debug(category)
	this.Data["json"] = Response{1,"添加成功"}
	this.ServeJson()
}


func (this * GoodsCategoryController)UpdateGoodsCategory() {
	admin := this.checkSession();
	if admin == nil{
		this.Data["json"]=Response{-1,"登陆超时,请重新登陆"}
		this.ServeJson()
		return
	}

	var category model.GoodsCategroy
	err := this.ParseForm(&category)
	if err != nil {
		beego.Error(err.Error())
		this.Data["json"] = Response{-1,"输入信息有误"}
		this.ServeJson()
	}

	beego.Debug(category)

	categorytmp,err := model.GetCategoryById(category.CategroyId)
	if err != nil ||categorytmp == nil{
		beego.Error(err.Error())
		this.Data["json"] = Response{-1,"不存在的商品分类"}
		this.ServeJson()
	}

	if categorytmp.GoodsId <= 0 && categorytmp.ShopId <= 0{
		this.Data["json"] = Response{-1,"输入信息有误"}
		this.ServeJson()
		return
	}


	categorytmp.Name = category.Name
	categorytmp.EnableTime = category.EnableTime
	categorytmp.Memo = category.Memo
	categorytmp.TotalNum = category.TotalNum
	categorytmp.Price = category.Price
	categorytmp.LimitPurchaseQuantity = 1;

	err = model.UpdateCategory(categorytmp)
	if err != nil {
		beego.Error(err.Error())
		this.Data["json"] = Response{-1,"更新失败"}
		this.ServeJson()
	}
	this.Data["json"] = Response{1,"更新成功"}
	this.ServeJson()
}

func (this * GoodsCategoryController)DeleteGoodsCategory() {
	admin := this.checkSession();
	if admin == nil{
		this.Data["json"]=Response{-1,"登陆超时,请重新登陆"}
		this.ServeJson()
		return
	}

	categoryId,err := this.GetInt("categoryId",0)
	if err != nil || categoryId <= 0{
		this.Data["json"]=Response{-1,"错误的分类"}
		this.ServeJson()
		return
	}

	category,err := model.GetCategoryById(categoryId)
	if err != nil || categoryId <= 0{
		this.Data["json"]=Response{-1,"不存在的分类"}
		this.ServeJson()
		return
	}

	orders,err := model.GetOrderByCategoryId(category.CategroyId)
	if err == nil && len(*orders) > 0 {
		this.Data["json"]=Response{-1,"有未完成的订单,不能删除分类"}
		this.ServeJson()
		return
	}
	err = model.DeleteCateGoryById(category.CategroyId)
	if err != nil {
		this.Data["json"]=Response{-1,"删除分类失败"}
		this.ServeJson()
		return
	}
	this.Data["json"]=Response{1,"删除分类成功"}
	this.ServeJson()
	return
}


func (this * GoodsCategoryController)checkSession() ( *model.Admin) {

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
