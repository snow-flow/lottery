package models

type LtCode struct {
	Id         int    `xorm:"not null pk autoincr INT"`
	GiftId     int    `xorm:"not null default 0 comment('奖品ID，关联lt_gift表') index INT"`
	Code       string `xorm:"not null default '' comment('虚拟券编码') unique VARCHAR(255)"`
	SysCreated int    `xorm:"not null default 0 comment('创建时间') INT"`
	SysUpdated int    `xorm:"not null default 0 comment('更新时间') INT"`
	SysStatus  int    `xorm:"not null default 0 comment('状态，0正常，1作废，2已发放') SMALLINT"`
}
