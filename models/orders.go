package model
import "github.com/astaxie/beego/orm"

type Orders struct {
	Orderid       int    `orm:"column(orderId);pk;auto;unique;size(11)"`
	Goodid        int    `orm:"column(goodId);index;size(11)"`
	CategroyId	  int	 `orm:"column(categroyId);index;size(11)"`
	CategroyName  string `orm:"column(categroyName);size(35)"`
	Shopid        int    `orm:"column(shopId);index;size(11)"`
	Uid           string `orm:"column(UID);index;size(35)"`
	Taobaoaccount string `orm:"column(taobaoAccount);index;size(35)"`
	State         int    `orm:"column(state);default(0);size(1)"`
	Shopname      string `orm:"column(shopName);size(35)"`
	Price         int64  `orm:"column(price);size(20)"`
	Requirelevel  int    `orm:"column(requireLevel);size(2)"`
	Shoprequire   string `orm:"column(shopRequire);size(400)"`
	Imageurl      string `orm:"column(imageUrl);size(100)"`
	Brokerage     int64  `orm:"column(brokerAge);size(20)"`
	Createtime    int64  `orm:"column(createTime);default(0);size(20)"`
	Updatetime    int64  `orm:"column(updateTime);default(0);size(20)"`
	Quantity      int    `orm:"column(quantity);size(11)"`
	Memo          string `orm:"column(memo);size(400)"`
	GoodsName     string `orm:"column(goodsName);size(50)"`
}

func init() {
	orm.RegisterModel(new(Orders))
}

func GetNewOrders()(*[]Orders,error) {
	var orders []Orders
	o := orm.NewOrm()
	_,err := o.Raw(`select * from orders order by updateTime desc limit 10`).QueryRows(&orders)
	return &orders,err
}