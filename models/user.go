package model
import "github.com/astaxie/beego/orm"

type User struct {
	Uid           string `orm:"pk;unique;size(35)"`
	Mobile        string `xorm:"not null unique VARCHAR(15)"`
	Password      string `xorm:"not null VARCHAR(35)"`
	Realname      string `xorm:"not null VARCHAR(35)"`
	Idcard        string `xorm:"not null unique VARCHAR(20)"`
	Alipayaccount string `xorm:"not null unique VARCHAR(35)"`
	Alipayname    string `xorm:"not null VARCHAR(35)"`
	Active        int    `xorm:"default 1 TINYINT(1)"`
	Asset         int64  `xorm:"default 0 BIGINT(10)"`
	Rate          int64  `xorm:"default 0 BIGINT(10)"`
	Income        int64  `xorm:"default 0 BIGINT(10)"`
	Total         int64  `xorm:"default 0 BIGINT(10)"`
	Createtime    int64  `xorm:"default 0 BIGINT(10)"`
	Updatetime    int64  `xorm:"default 0 BIGINT(10)"`
}

func init() {
	orm.RegisterModel(new(User))
}