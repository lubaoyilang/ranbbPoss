package model
import "github.com/astaxie/beego/orm"

type TaobaoAccount struct {
	Tid           int    `orm:"pk;auto;unique;size(11)"`
	Uid           string `xorm:"not null index VARCHAR(35)"`
	Taobaoaccount string `xorm:"not null unique VARCHAR(35)"`
	Createtime    int64  `xorm:"default 0 BIGINT(20)"`
	Updatetime    int64  `xorm:"default 0 BIGINT(20)"`
	Memo          string `xorm:"BLOB"`
}
func init() {
	orm.RegisterModel(new(TaobaoAccount))
}