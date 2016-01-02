package model

type TaobaoAccount struct {
	Tid           int    `xorm:"not null pk autoincr unique INT(11)"`
	Uid           string `xorm:"not null index VARCHAR(35)"`
	Taobaoaccount string `xorm:"not null unique VARCHAR(35)"`
	Createtime    int64  `xorm:"default 0 BIGINT(20)"`
	Updatetime    int64  `xorm:"default 0 BIGINT(20)"`
	Memo          []byte `xorm:"BLOB"`
}
