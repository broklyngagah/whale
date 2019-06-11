package JYBaseDB

import (
	"carp.cn/whale/db"
	"carp.cn/whale/zaplogger"
)

//This file is generate by scripts,don't edit it

//sys_blacklist
//手机号、邮箱、IP黑名单数据表

// +gen
type SysBlacklist struct {
	Id          int    `db:"id" json:"id"`                     //
	Type        int8   `db:"type" json:"type"`                 // 1手机号 2IP地址 3邮箱
	Body        string `db:"body" json:"body"`                 // 手机号或IP地址
	Status      int8   `db:"status" json:"status"`             // 1屏蔽状态 0已解除屏蔽
	Created     int    `db:"created" json:"created"`           // 添加时间
	Operator    string `db:"operator" json:"operator"`         // 添加操作人
	RemovedTime int64  `db:"removed_time" json:"removed_time"` // 解除时间
	Remover     string `db:"remover" json:"remover"`           // 解除操作者
}

var DefaultSysBlacklist = SysBlacklist{}

type sysBlacklistCache struct {
	objMap  map[int]*SysBlacklist
	objList []*SysBlacklist
}

var SysBlacklistCache = &sysBlacklistCache{}

func (c *sysBlacklistCache) LoadAll() {
	sql := "select `id`,`type`,`body`,`status`,`created`,`operator`,`removed_time`,`remover` from sys_blacklist"
	c.objList = make([]*SysBlacklist, 0)
	err := db.JYBaseDB.Select(&c.objList, sql)
	if err != nil {
		zaplogger.Fatal(err.Error())
	}
	c.objMap = make(map[int]*SysBlacklist)
	for _, v := range c.objList {
		c.objMap[v.Id] = v
	}
}

func (c *sysBlacklistCache) All() []*SysBlacklist {
	return c.objList
}

func (c *sysBlacklistCache) Count() int {
	return len(c.objList)
}

func (c *sysBlacklistCache) Get(id int) (*SysBlacklist, bool) {
	key := id
	v, ok := c.objMap[key]
	return v, ok
}

// 仅限运营后台实时刷新服务器数据用
func (c *sysBlacklistCache) Update(v *SysBlacklist) {
	key := v.Id
	c.objMap[key] = v
}
