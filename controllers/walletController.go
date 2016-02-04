package controllers
import (
"github.com/astaxie/beego"
"ranbbPoss/models"
	"errors"
	"time"
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

	page,err := this.GetInt("page",0)
	if err != nil {
		page = 0;
	}

	enchashments,err :=  model.GetUserEnchashmentList(page)
	if err != nil {
		this.Data["json"]=Response{-1,"获取提现列表失败"}
		this.ServeJson()
		return
	}
	this.Data["json"]=Response{1,enchashments}
	this.ServeJson()
	return
}
//允许取现
func (this * WalletController)Enchash() {
	admin := new(model.Admin)
	if v := this.checkSession();v == nil {
		return
	}else{
		admin = v
	}

	enchashId,err := this.GetInt("enchashId",0)
	if err != nil ||enchashId <= 0{
		this.Data["json"]=Response{-1,"错误的请求,刷新页面重试"}
		this.ServeJson()
		return
	}

	enchashInfo,err := model.GetUserEnchashmentById(enchashId)
	if err != nil {
		this.Data["json"]=Response{-1,"未发现相应记录"}
		this.ServeJson()
		return
	}
	if enchashInfo.State == 1 {
		this.Data["json"]=Response{-1,"已经取现成功,请刷新页面"}
		this.ServeJson()
		return
	}

	if enchashInfo.Amount > enchashInfo.Asset+enchashInfo.Rate+enchashInfo.Income {
		this.Data["json"]=Response{-1,"金额不足,不能取现"}
		this.ServeJson()
		return
	}

	err = model.UpdateEnchashmentState(1,enchashId,admin.Name)
	if err != nil {
		this.Data["json"]=Response{-1,"取现失败,更新状态失败"}
		this.ServeJson()
		return
	}

	err = changeUserAmount(enchashInfo.Amount,enchashInfo.UID)
	if err != nil {
		this.Data["json"]=Response{-1,"取现失败,更新状态失败"}
		this.ServeJson()
		return
	}
	err = addLogToWalletLog(enchashInfo.Amount,enchashInfo.UID)
	if err != nil {
		this.Data["json"]=Response{-1,"取现失败,更新状态失败"}
		this.ServeJson()
		return
	}

	this.Data["json"]=Response{1,"取现成功"}
	this.ServeJson()
	return
}

func (this * WalletController)CanNotEnchash() {
	admin := new(model.Admin)
	if v := this.checkSession();v == nil {
		return
	}else{
		admin = v
	}
	enchashId,err := this.GetInt("enchashId",0)
	if err != nil ||enchashId <= 0{
		this.Data["json"]=Response{-1,"错误的请求,刷新页面重试"}
		this.ServeJson()
		return
	}
	err = model.UpdateEnchashmentState(2,enchashId,admin.Name)
	if err != nil {
		this.Data["json"]=Response{-1,"取现失败,更新状态失败"}
		this.ServeJson()
		return
	}

	this.Data["json"]=Response{1,"已禁止用户取现"}
	this.ServeJson()
	return
}


func changeUserAmount(amount int64, UID string) error{
	user,err := model.GetUserById(UID)
	if err != nil {
		return err
	}
	user.Total -= amount
	user.Asset -= amount
	if user.Asset < 0 {
		user.Income += user.Asset
		user.Asset = 0
		if user.Income < 0 {
			user.Rate += user.Income
			user.Income = 0
			if user.Rate < 0 {
				return errors.New("用户资金不够")
			}
		}
	}
	err = model.UpdateUser(user)
	return err
}

func addLogToWalletLog(amount int64, UID string) error {
	log := 	&model.WalletLog{
		Uid       :UID,
		Price      :amount,
		Categroy   :2,
		Createtime :time.Now().Unix(),
		Memo       :"取现"}
	return model.AddLog(log)
}