package models

type LtBlackip struct {
	Id         int    `xorm:"not null pk autoincr INT"`
	Ip         string `xorm:"not null default '' comment('IP地址') unique VARCHAR(50)"`
	Blacktime  int    `xorm:"not null default 0 comment('黑名单限制到期时间') INT"`
	SysCreated int    `xorm:"not null default 0 comment('创建时间') INT"`
	SysUpdated int    `xorm:"not null default 0 comment('修改时间') INT"`
}
