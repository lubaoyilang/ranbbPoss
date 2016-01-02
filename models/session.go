package model

type Session struct {
	SessionKey    string `xorm:"not null pk unique VARCHAR(64)"`
	SessionData   string `xorm:"VARCHAR(255)"`
	SessionExpiry int64  `xorm:"BIGINT(20)"`
}
