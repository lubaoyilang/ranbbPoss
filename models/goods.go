package model
import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

type Goods struct {
	Goodid                int    `orm:"column(goodId);pk;auto;unique;size(11)"`
	Shopid                int    `orm:"column(shopId);size(15);index"`
	Shopname              string `orm:"column(shopName);size(60)"`
	State                 int    `orm:"column(state);default(1);size(1)"`
	Requirelevel          int    `orm:"column(requireLevel);size(2)"`
	Shoprequire           string `orm:"column(shopRequire);size(300)"`
	Imageurl              string `orm:"column(imageUrl);size(300)"`
	Brokerage             int64  `orm:"column(brokerAge);size(20)"`
	Createtime            int64  `orm:"column(createTime);default(0);size(20)"`
	Updatetime            int64  `orm:"column(updateTime);default(0);size(20)"`
	GoodsName			  string `orm:"column(goodsName);size(60)"`
	Memo                  string `orm:"column(memo);size(300)"`
	SettingPrice          int64  `orm:"column(settingPrice);size(20)"`
}

func init() {
	orm.RegisterModel(new(Goods))
}


func GetGoodsByShopId(shopId int) (*[]Goods,error) {
	goods := make([]Goods,0)
	o := orm.NewOrm()
	_,err := o.Raw(`select * from goods where shopId=?`,shopId).QueryRows(&goods)
	return &goods,err
}

func GetGoodsById(goodsId int) (v *Goods,err error) {
	o := orm.NewOrm()
	v = &Goods{Goodid: goodsId}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func UpdateGoods(goods * Goods) (err error) {
	o := orm.NewOrm()
	v := Goods{Goodid: goods.Goodid}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(goods); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}