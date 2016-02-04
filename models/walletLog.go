package model
import "github.com/astaxie/beego/orm"

type WalletLog struct {
	Wid        int    `orm:"pk;auto;unique;size(11)"`
	Uid        string ` orm:"index;size(35)"`
	Price      int64  `orm:"size(20)"`
	NowValue   int64  `orm:"column(nowValue);size(11)"`
	Categroy   int    `orm:"size(11)"`
	Createtime int64  `orm:"default(0);size(20)"`
	Memo       string `orm:"type(TEXT)"`
}
func init() {
	orm.RegisterModel(new(WalletLog))
}

func AddLog(log * WalletLog) error {
	o := orm.NewOrm()
	o.Begin()
	_,err := o.Insert(log)
	if err != nil {
		o.Rollback()
		return err
	}else{
		return o.Commit()
	}
	return err
}