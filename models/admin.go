package model

type Admin struct {
	Uid       string `xorm:"not null pk unique VARCHAR(35)"`
	Name      string `xorm:"VARCHAR(35)"`
	Mobile    string `xorm:"not null unique VARCHAR(15)"`
	Password  string `xorm:"not null VARCHAR(35)"`
	Privilege string `xorm:"not null VARCHAR(100)"`
}
