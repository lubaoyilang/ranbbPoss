package controllers
import (
	"github.com/astaxie/beego"
	"ranbbPoss/models"
	"os"
	"encoding/csv"
	"fmt"
	"time"
	"reflect"
	"strconv"
)


type OrderController struct {
	beego.Controller
}

func (this * OrderController)GetNewOrders() {
	admin := this.checkSession()
	if admin == nil {
		this.ServeJson()
		return
	}

	orders,err := model.GetNewOrders()
	if err != nil {
		this.ServeJson()
		return
	}
	this.Data["json"]=orders
	this.ServeJson()
}


func (this * OrderController)ExportOrders()  {
	admin := this.checkSession()
	if admin == nil {
		this.Data["json"] = Response{-1,"还没好呢"}
		this.ServeJson()
		return
	}

	goodId,err := this.GetInt("goodId",0)
	if err != nil || goodId <= 0{
		this.Data["json"] = Response{-1,"参数错误"}
		this.ServeJson()
		return
	}

	orders,err :=  model.GetOrdersByGoodsId(goodId,1)
	if err != nil {
		this.Data["json"] = Response{-1,"没有找到相关订单"}
		this.ServeJson()
		return
	}
	f, err := os.Create("test.xls")
	if err != nil {
		this.Data["json"] = Response{-1,"生成表格失败"}
		this.ServeJson()
		return
	}
	defer f.Close()
	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	w := csv.NewWriter(f)
	w.Write([]string{"订单号","店铺名称","商品名称","分类名称","买家姓名","刷单淘宝账号","状态","价格","佣金","要求等级","刷单要求","更新时间","备注"})
	for _,order := range *orders {
		w.Write(modelToStrings(order))
	}
	w.Flush()
	this.Ctx.Output.Download("test.xls",fmt.Sprintf("订单导出-%s.xls",time.Now().Format("2006-01-02_15:04")))
	os.Remove("test.xls")
}


func (this *OrderController)GetOrders() {
	admin := this.checkSession()
	if admin == nil {
		this.Data["json"] = Response{-1,"还没好呢"}
		this.ServeJson()
		return
	}

	goodId,err := this.GetInt("goodsId",0)
	if err != nil || goodId <= 0{
		this.Data["json"] = Response{-1,"参数错误"}
		this.ServeJson()
		return
	}

	page,err := this.GetInt("page",0)
	if err != nil {
		page = 0
	}

	orders,err :=  model.GetAllOrdersByGoodsIdAndPage(goodId,page)
	if err != nil {
		this.Data["json"] = Response{-1,"查询错误"}
		this.ServeJson()
		return
	}

	this.Data["json"] = Response{1,orders}
	this.ServeJson()

}


func (this * OrderController)checkSession() ( *model.Admin) {
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

func modelToStrings(m interface{}) []string {
	var values []string
	mType := reflect.TypeOf(m)
	mValue := reflect.ValueOf(m)
	for i := 0 ;i < mType.NumField() ;i++  {
		key := mType.Field(i).Name
		value := mValue.Field(i);
		if key == "State" {
			if value.String() == "0" {
				values = append(values,"正在刷单")
			}else if value.String() == "1" {
				values = append(values,"等待审核")
			}else if value.String() == "2" {
				values = append(values,"审核通过")
			}else if value.String() == "3" {
				values = append(values,"审核失败")
			}
		}else if  key == "Updatetime"{
			updatetime := value.String()
			timeint,err := strconv.Atoi(updatetime)
			if err != nil {
				values = append(values,"")
			}else{
				values = append(values,time.Unix(int64(timeint),0).Format("2006-01-02 15:04"))
			}
		}else{
			values = append(values,fmt.Sprintf("%v",value.Interface()))
		}
	}
	beego.Debug(values)
	return values
}


func (this * OrderController)ChangeOrderState()  {
	admin := this.checkSession()
	if admin == nil {
		return
	}

	orderId,err:= this.GetInt("orderId",0)
	if err != nil {
		beego.Debug("输入数据有误 orderId",orderId)
		this.Data["json"] = Response{-1,"数据有误"}
		this.ServeJson()
		return
	}
	orderState,err := this.GetInt("orderState",0)
	if err != nil {
		beego.Debug("输入数据有误 orderState",orderState)
		this.Data["json"] = Response{-1,"数据有误"}
		this.ServeJson()
		return
	}

	order,err := model.GetOrderByOrderId(orderId)
	if err != nil {
		beego.Debug(err.Error())
		this.Data["json"] = Response{-1,"不存在的订单"}
		this.ServeJson()
		return
	}
	order.State = orderState
	err = model.UpdateOrderState(order)
	if err != nil {
		beego.Debug(err.Error())
		this.Data["json"] = Response{-1,"更新订单状态失败"}
		this.ServeJson()
		return
	}

	if orderState == 2 {
		err = changeAmount(order.Price,order.Brokerage,order.Uid)
		if err != nil {
			this.Data["json"] = Response{-1,"更新订单状态失败"}
			this.ServeJson()
			return
		}
	}
	this.Data["json"] = Response{1,"更新成功"}
	this.ServeJson()
	return
}

//审核通过 本金增加 佣金增加 总金额增加 log 增加 减少审核中金额

func changeAmount(price,income int64,UID string) error {
	user,err := model.GetUserById(UID)
	if err != nil {
		return err
	}
	user.Asset += price
	user.Income += income
	user.VerifyAmount -= price
	user.Total += income
	err = model.UpdateUser(user)
	if err != nil {
		return err
	}
	log := model.WalletLog{
		Price:price,
		Uid:UID,
		NowValue:user.Asset,
		Createtime:time.Now().Unix(),
		Categroy:0,
		Memo:"本金返还",
	}
	err = model.AddLog(&log);
	if err != nil {
		return err
	}
	log = model.WalletLog{
		Price:income,
		Uid:UID,
		NowValue:user.Income,
		Createtime:time.Now().Unix(),
		Categroy:3,
		Memo:"佣金增加",
	}
	err = model.AddLog(&log);
	if err != nil {
		return err
	}

	log = model.WalletLog{
		Price:0-income,
		Uid:UID,
		NowValue:user.VerifyAmount,
		Createtime:time.Now().Unix(),
		Categroy:4,
		Memo:"审核中金额转入本金",
	}
	err = model.AddLog(&log);
	if err != nil {
		return err
	}
	return nil
}