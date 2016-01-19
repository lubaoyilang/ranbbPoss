package routers

import (
	"ranbbPoss/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/ranbb/admin", &controllers.MainController{})
	beego.Router("/ranbb/checkSession",&controllers.MainController{},"Post:CheckSession")
	beego.Router("/ranbb/login",&controllers.MainController{},"Post:Login")
	beego.Router("/ranbb/logout",&controllers.MainController{},"post:Logout")

	beego.Router("/ranbb/neworders", &controllers.OrderController{},"Post:GetNewOrders")


	beego.Router("/ranbb/newShops",&controllers.ShopController{},"Post:GetNewShops")
	beego.Router("/ranbb/getShopList",&controllers.ShopController{},"post:GetShopList")
	beego.Router("/ranbb/updateShop",&controllers.ShopController{},"post:UpdateShop")


	beego.Router("/ranbb/getAdminList",&controllers.AdminController{},"Post:GetAdminList")
	beego.Router("/ranbb/addAdmin",&controllers.AdminController{},"Post:AddAdmin")
	beego.Router("/ranbb/deleteAdmin",&controllers.AdminController{},"post:DeleteAdmin")
	beego.Router("/ranbb/updateAdmin",&controllers.AdminController{},"post:UpdateAdmin")
}
