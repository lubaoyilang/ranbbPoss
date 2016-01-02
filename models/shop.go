package model

type Shop struct {
	Shopid       int    `xorm:"not null pk autoincr unique INT(11)"`
	Shopname     string `xorm:"VARCHAR(35)"`
	Mobile       string `xorm:"not null unique VARCHAR(15)"`
	Email        string `xorm:"not null unique VARCHAR(25)"`
	Shoptaobaoid string `xorm:"not null unique VARCHAR(35)"`
	Createtime   int64  `xorm:"default 0 BIGINT(10)"`
	Updatetime   int64  `xorm:"default 0 BIGINT(10)"`
	Memo         []byte `xorm:"BLOB"`
}
