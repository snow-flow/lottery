package models

type LtResult struct {
	Id         int    `xorm:"not null pk autoincr INT"`
	GiftId     int    `xorm:"not null default 0 comment('奖品ID，关联lt_gift表') index INT"`
	GiftName   string `xorm:"not null default '' comment('奖品名称') VARCHAR(255)"`
	GiftType   int    `xorm:"not null default 0 comment('奖品类型，同lt_gift. gtype') INT"`
	Uid        int    `xorm:"not null default 0 comment('用户ID') index INT"`
	Username   string `xorm:"not null default '' comment('用户名') VARCHAR(50)"`
	PrizeCode  int    `xorm:"not null default 0 comment('抽奖编号（4位的随机数）') INT"`
	GiftData   string `xorm:"not null default '' comment('获奖信息') VARCHAR(255)"`
	SysCreated int    `xorm:"not null default 0 comment('创建时间') INT"`
	SysIp      string `xorm:"not null default '' comment('用户抽奖的IP') VARCHAR(50)"`
	SysStatus  int    `xorm:"not null default 0 comment('状态，0 正常，1删除，2作弊') SMALLINT"`
}
