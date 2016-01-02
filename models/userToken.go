package model

type UserToken struct {
	SessionKey    string `xorm:"not null pk unique VARCHAR(64)"`
	SessionData   string `xorm:"VARCHAR(35)"`
	SessionExpiry int64  `xorm:"BIGINT(20)"`
}
