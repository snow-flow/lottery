package models

type LtUser struct {
	Id         int    `xorm:"not null pk autoincr INT"`
	Username   string `xorm:"not null default '' comment('用户名') VARCHAR(50)"`
	Blacktime  int    `xorm:"not null default 0 comment('黑名单限制到期时间') INT"`
	Realname   string `xorm:"not null default '' comment('联系人') VARCHAR(50)"`
	Mobile     string `xorm:"not null default '' comment('手机号') VARCHAR(50)"`
	Address    string `xorm:"not null default '' comment('联系地址') VARCHAR(255)"`
	SysCreated int    `xorm:"not null default 0 comment('创建时间') INT"`
	SysUpdated int    `xorm:"not null default 0 comment('修改时间') INT"`
	SysIp      string `xorm:"not null default '' comment('IP地址') VARCHAR(50)"`
}
