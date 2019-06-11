package JYBaseDB

import (
	"carp.cn/whale/db"
	"carp.cn/whale/zaplogger"
)

//This file is generate by scripts,don't edit it

//sys_auth_group
//系统后台 - 权限

// +gen
type SysAuthGroup struct {
	Id     int    `db:"id" json:"id"`         //
	Title  string `db:"title" json:"title"`   // 角色标题
	Status int8   `db:"status" json:"status"` // 角色状态: 1.正常  其他.不正常
	Rules  string `db:"rules" json:"rules"`   // 权限ID
	CTime  int64  `db:"c_time" json:"c_time"` //
	UTime  int64  `db:"u_time" json:"u_time"` //
}

var DefaultSysAuthGroup = SysAuthGroup{}

type sysAuthGroupCache struct {
	objMap  map[int]*SysAuthGroup
	objList []*SysAuthGroup
}

var SysAuthGroupCache = &sysAuthGroupCache{}

func (c *sysAuthGroupCache) LoadAll() {
	sql := "select `id`,`title`,`status`,`rules`,`c_time`,`u_time` from sys_auth_group"
	c.objList = make([]*SysAuthGroup, 0)
	err := db.JYBaseDB.Select(&c.objList, sql)
	if err != nil {
		zaplogger.Fatal(err.Error())
	}
	c.objMap = make(map[int]*SysAuthGroup)
	for _, v := range c.objList {
		c.objMap[v.Id] = v
	}
}

func (c *sysAuthGroupCache) All() []*SysAuthGroup {
	return c.objList
}

func (c *sysAuthGroupCache) Count() int {
	return len(c.objList)
}

func (c *sysAuthGroupCache) Get(id int) (*SysAuthGroup, bool) {
	key := id
	v, ok := c.objMap[key]
	return v, ok
}

// 仅限运营后台实时刷新服务器数据用
func (c *sysAuthGroupCache) Update(v *SysAuthGroup) {
	key := v.Id
	c.objMap[key] = v
}
