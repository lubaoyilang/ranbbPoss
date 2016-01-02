package model

type WalletLog struct {
	Wid        int    `xorm:"not null pk autoincr unique INT(11)"`
	Uid        string `xorm:"not null index VARCHAR(35)"`
	Price      int64  `xorm:"BIGINT(20)"`
	Categroy   int    `xorm:"INT(11)"`
	Createtime int64  `xorm:"default 0 BIGINT(20)"`
	Memo       []byte `xorm:"BLOB"`
}
