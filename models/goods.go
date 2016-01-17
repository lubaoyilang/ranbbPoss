package model
import "github.com/astaxie/beego/orm"

type Goods struct {
	Goodid                int    `orm:"pk;auto;unique;size(11)"`
	Shopid                int    //`xorm:"not null index INT(11)"`
	Shopname              string //`xorm:"VARCHAR(35)"`
	State                 int    //`xorm:"default 1 TINYINT(1)"`
	Price                 int64  //`xorm:"BIGINT(20)"`
	Requirelevel          int    //`xorm:"TINYINT(2)"`
	Shoprequire           string //`xorm:"BLOB"`
	Imageurl              string //`xorm:"BLOB"`
	Brokerage             int64  //`xorm:"BIGINT(20)"`
	Createtime            int64  //`xorm:"default 0 BIGINT(20)"`
	Updatetime            int64  //`xorm:"default 0 BIGINT(20)"`
	Quantity              int    //`xorm:"INT(11)"`
	Limitpurchasequantity int    //`xorm:"default 1 INT(11)"`
	Memo                  string //`xorm:"BLOB"`
}

func init() {
	orm.RegisterModel(new(Goods))
}