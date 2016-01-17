package model
import "github.com/astaxie/beego/orm"

type UserToken struct {
	SessionKey    string `orm:"pk;unique;size(64)"`
	SessionData   string `xorm:"VARCHAR(35)"`
	SessionExpiry int64  `xorm:"BIGINT(20)"`
}
func init() {
	orm.RegisterModel(new(UserToken))
}