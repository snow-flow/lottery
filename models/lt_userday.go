package models

type LtUserday struct {
	Id         int `xorm:"not null pk autoincr INT"`
	Uid        int `xorm:"not null default 0 comment('用户ID') unique(uid_day) INT"`
	Day        int `xorm:"not null default 0 comment('日期，如：20180725') unique(uid_day) INT"`
	Num        int `xorm:"not null default 0 comment('次数') INT"`
	SysCreated int `xorm:"not null default 0 comment('创建时间') INT"`
	SysUpdated int `xorm:"not null default 0 comment('修改时间') INT"`
}
