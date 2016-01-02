package model

type Goods struct {
	Goodid                int    `xorm:"not null pk autoincr unique INT(11)"`
	Shopid                int    `xorm:"not null index INT(11)"`
	Shopname              string `xorm:"VARCHAR(35)"`
	State                 int    `xorm:"default 1 TINYINT(1)"`
	Price                 int64  `xorm:"BIGINT(20)"`
	Requirelevel          int    `xorm:"TINYINT(2)"`
	Shoprequire           []byte `xorm:"BLOB"`
	Imageurl              []byte `xorm:"BLOB"`
	Brokerage             int64  `xorm:"BIGINT(20)"`
	Createtime            int64  `xorm:"default 0 BIGINT(20)"`
	Updatetime            int64  `xorm:"default 0 BIGINT(20)"`
	Quantity              int    `xorm:"INT(11)"`
	Limitpurchasequantity int    `xorm:"default 1 INT(11)"`
	Memo                  []byte `xorm:"BLOB"`
}
