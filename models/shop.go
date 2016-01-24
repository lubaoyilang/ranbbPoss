package model
import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

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
func GetShopById(id int) (v *Shop, err error) {
	o := orm.NewOrm()
	v = &Shop{Shopid: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}
func UpdateShopById(m *Shop) (err error) {
	o := orm.NewOrm()
	v := Shop{Shopid: m.Shopid}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

func GetShopByName(name string) (v *Shop, err error) {
	o := orm.NewOrm()
	v = &Shop{Shopname: name}
	if err = o.Read(v,"shopName"); err == nil {
		return v, nil
	}
	return nil, err
}

func GetShopByMobile(mobile string) (v *Shop, err error) {
	o := orm.NewOrm()
	v = &Shop{Mobile: mobile}
	if err = o.Read(v,"mobile"); err == nil {
		return v, nil
	}
	return nil, err
}

func GetShopByEmail(email string) (v *Shop, err error) {
	o := orm.NewOrm()
	v = &Shop{Email: email}
	if err = o.Read(v,"email"); err == nil {
		return v, nil
	}
	return nil, err
}
func GetShopByShoptaobaoid(shoptaobaoid string) (v *Shop, err error) {
	o := orm.NewOrm()
	v = &Shop{Shoptaobaoid: shoptaobaoid}
	if err = o.Read(v,"shopTaoBaoId"); err == nil {
		return v, nil
	}
	return nil, err
}


func AddShop(s * Shop) (int64,error) {
	o := orm.NewOrm()
	o.Begin()
	if id,err := o.Insert(s);err == nil&&id > 0 {
		return id,o.Commit()
	}else {
		o.Rollback()
		return 0,err
	}
}