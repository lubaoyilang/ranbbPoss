package model
import "github.com/astaxie/beego/orm"

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
