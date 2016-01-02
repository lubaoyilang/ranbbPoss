package model

type Orders struct {
	Orderid       int    `xorm:"not null pk autoincr unique INT(11)"`
	Goodid        int    `xorm:"not null index INT(11)"`
	Shopid        int    `xorm:"not null index INT(11)"`
	Uid           string `xorm:"not null index VARCHAR(35)"`
	Taobaoaccount string `xorm:"index VARCHAR(35)"`
	State         int    `xorm:"default 0 TINYINT(1)"`
	Shopname      string `xorm:"VARCHAR(35)"`
	Price         int64  `xorm:"BIGINT(20)"`
	Requirelevel  int    `xorm:"TINYINT(2)"`
	Shoprequire   []byte `xorm:"BLOB"`
	Imageurl      []byte `xorm:"BLOB"`
	Brokerage     int64  `xorm:"BIGINT(20)"`
	Createtime    int64  `xorm:"default 0 BIGINT(20)"`
	Updatetime    int64  `xorm:"default 0 BIGINT(20)"`
	Quantity      int    `xorm:"INT(11)"`
	Memo          []byte `xorm:"BLOB"`
}
