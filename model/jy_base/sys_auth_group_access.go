package JYBaseDB

import (
	"carp.cn/whale/db"
	"carp.cn/whale/zaplogger"
)

//This file is generate by scripts,don't edit it

//sys_auth_group_access
//系统后台 - 帐户与权限关系表

// +gen
type SysAuthGroupAccess struct {
	Id      int `db:"id" json:"id"`             //
	Uid     int `db:"uid" json:"uid"`           //
	GroupId int `db:"group_id" json:"group_id"` //
}

var DefaultSysAuthGroupAccess = SysAuthGroupAccess{}

type sysAuthGroupAccessCache struct {
	objMap  map[int]*SysAuthGroupAccess
	objList []*SysAuthGroupAccess
}

var SysAuthGroupAccessCache = &sysAuthGroupAccessCache{}

func (c *sysAuthGroupAccessCache) LoadAll() {
	sql := "select `id`,`uid`,`group_id` from sys_auth_group_access"
	c.objList = make([]*SysAuthGroupAccess, 0)
	err := db.JYBaseDB.Select(&c.objList, sql)
	if err != nil {
		zaplogger.Fatal(err.Error())
	}
	c.objMap = make(map[int]*SysAuthGroupAccess)
	for _, v := range c.objList {
		c.objMap[v.Id] = v
	}
}

func (c *sysAuthGroupAccessCache) All() []*SysAuthGroupAccess {
	return c.objList
}

func (c *sysAuthGroupAccessCache) Count() int {
	return len(c.objList)
}

func (c *sysAuthGroupAccessCache) Get(id int) (*SysAuthGroupAccess, bool) {
	key := id
	v, ok := c.objMap[key]
	return v, ok
}

// 仅限运营后台实时刷新服务器数据用
func (c *sysAuthGroupAccessCache) Update(v *SysAuthGroupAccess) {
	key := v.Id
	c.objMap[key] = v
}
