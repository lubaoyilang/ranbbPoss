package routers

import (
	"ranbbPoss/controllers"
	"github.com/astaxie/beego"
//	"github.com/astaxie/beego/context"
)

func init() {
//    beego.Router("/ranbb/admin", &controllers.MainController{})
	beego.Router("/ranbb/checkSession",&controllers.MainController{},"Post:CheckSession")
	beego.Router("/ranbb/login",&controllers.MainController{},"Post:Login")
	beego.Router("/ranbb/logout",&controllers.MainController{},"post:Logout")

	beego.Router("/ranbb/neworders", &controllers.OrderController{},"Post:GetNewOrders")


	beego.Router("/ranbb/newShops",&controllers.ShopController{},"Post:GetNewShops")
	beego.Router("/ranbb/getShopList",&controllers.ShopController{},"post:GetShopList")
	beego.Router("/ranbb/updateShop",&controllers.ShopController{},"post:UpdateShop")
	beego.Router("/ranbb/addShop",&controllers.ShopController{},"post:AddShop")


	beego.Router("/ranbb/getAdminList",&controllers.AdminController{},"Post:GetAdminList")
	beego.Router("/ranbb/addAdmin",&controllers.AdminController{},"Post:AddAdmin")
	beego.Router("/ranbb/deleteAdmin",&controllers.AdminController{},"post:DeleteAdmin")
	beego.Router("/ranbb/updateAdmin",&controllers.AdminController{},"post:UpdateAdmin")


	beego.Router("/ranbb/getGoodsList",&controllers.GoodsController{},"post:GetGoodsList")
	beego.Router("/ranbb/updateGoods",&controllers.GoodsController{},"post:UpdateGoods")
	beego.Router("/ranbb/addGoods",&controllers.GoodsController{},"post:AddGoods")

	beego.Router("/ranbb/getCategorys",&controllers.GoodsCategoryController{},"post:GetGoodsCategorys")
	beego.Router("/ranbb/updateCategory",&controllers.GoodsCategoryController{},"post:UpdateGoodsCategory")
	beego.Router("/ranbb/addCategory",&controllers.GoodsCategoryController{},"post:AddGoodsCategory")
	beego.Router("/ranbb/deleteCategory",&controllers.GoodsCategoryController{},"post:DeleteGoodsCategory")
	beego.Router("/ranbb/editCategory",&controllers.GoodsCategoryController{},"post:UpdateGoodsCategory")

//	var FilterUser = func(ctx *context.Context) {
//		_, ok := ctx.Input.Session("uid").(int)
//		if !ok && ctx.Request.RequestURI != "/ranbb/admin/#/sessions/signin" {
//			ctx.Redirect(302, "/ranbb/admin/#/sessions/signin")
//		}
//	}
//
//	beego.InsertFilter("/*",beego.BeforeRouter,FilterUser)
}
