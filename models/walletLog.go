package model
import "github.com/astaxie/beego/orm"

type WalletLog struct {
	Wid        int    `orm:"pk;autoincr;unique;size(11)"`
	Uid        string `xorm:"not null index VARCHAR(35)"`
	Price      int64  `xorm:"BIGINT(20)"`
	Categroy   int    `xorm:"INT(11)"`
	Createtime int64  `xorm:"default 0 BIGINT(20)"`
	Memo       string `xorm:"BLOB"`
}
func init() {
	orm.RegisterModel(new(WalletLog))
}