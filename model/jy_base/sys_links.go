package JYBaseDB

import (
	"carp.cn/whale/db"
	"carp.cn/whale/zaplogger"
)

//This file is generate by scripts,don't edit it

//sys_links
//友情链接

// +gen
type SysLinks struct {
	Id     int8   `db:"id" json:"id"`         //
	Name   string `db:"name" json:"name"`     // 名称
	Link   string `db:"link" json:"link"`     // 链接地址
	Ord    int8   `db:"ord" json:"ord"`       // 排序(越大越前面)
	Status int8   `db:"status" json:"status"` // 状态：1为启用 2为禁用
	Note   string `db:"note" json:"note"`     // 备注
	CTime  int64  `db:"c_time" json:"c_time"` // 创建时间
	Type   int8   `db:"type" json:"type"`     // 友情链接显示页面:1为首页，2为说有页面
	Form   int8   `db:"form" json:"form"`     // 链接的类型 1为站外链接 2为站内链接
	UTime  int64  `db:"u_time" json:"u_time"` //
}

var DefaultSysLinks = SysLinks{}

type sysLinksCache struct {
	objMap  map[int8]*SysLinks
	objList []*SysLinks
}

var SysLinksCache = &sysLinksCache{}

func (c *sysLinksCache) LoadAll() {
	sql := "select `id`,`name`,`link`,`ord`,`status`,`note`,`c_time`,`type`,`form`,`u_time` from sys_links"
	c.objList = make([]*SysLinks, 0)
	err := db.JYBaseDB.Select(&c.objList, sql)
	if err != nil {
		zaplogger.Fatal(err.Error())
	}
	c.objMap = make(map[int8]*SysLinks)
	for _, v := range c.objList {
		c.objMap[v.Id] = v
	}
}

func (c *sysLinksCache) All() []*SysLinks {
	return c.objList
}

func (c *sysLinksCache) Count() int {
	return len(c.objList)
}

func (c *sysLinksCache) Get(id int8) (*SysLinks, bool) {
	key := id
	v, ok := c.objMap[key]
	return v, ok
}

// 仅限运营后台实时刷新服务器数据用
func (c *sysLinksCache) Update(v *SysLinks) {
	key := v.Id
	c.objMap[key] = v
}
