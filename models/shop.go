package model
import "github.com/astaxie/beego/orm"

type Shop struct {
	Shopid       int    `orm:"column(shopId);unique;auto;pk"`
	Shopname     string `orm:"column(shopName);size(60)"`
	Mobile       string `orm:"column(mobile);unique;size(15);index"`
	Email        string `orm:"column(email);unique;size(25)"`
	Shoptaobaoid string `orm:"column(shopTaoBaoId);unique;size(35)"`
	ImageUrl     string	`orm:"column(imageUrl);size(100)"`
	Createtime   int64  `orm:"column(createTime);default(0);size(10)"`
	Updatetime   int64  `orm:"column(updateTime);default(0);size(10)"`
	Memo         string `orm:"column(memo);size(400)"`
}

func init() {
	orm.RegisterModel(new(Shop))
}

func GetNewShop() (*[]Shop,error) {
	var shops []Shop
	o := orm.NewOrm()
	_,err:=o.Raw(`select * from shop order by createTime desc limit 10 `).QueryRows(&shops)
	return &shops,err;
}

func GetShopByPage(page,size int)(*[]Shop,error) {
	var shops []Shop
	if size == 0 {
		size = 20
	}

	if page == 0 {
		page = 1
	}
	o := orm.NewOrm()
	_,err:=o.Raw(`select * from shop order by createTime desc limit ? offset ? `,size,(page-1)*size).QueryRows(&shops)
	return &shops,err;
}