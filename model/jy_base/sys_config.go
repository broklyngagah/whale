package JYBaseDB

import (
	"carp.cn/whale/db"
	"carp.cn/whale/zaplogger"
)

//This file is generate by scripts,don't edit it

//sys_config
//系统后台 - 配置

// +gen
type SysConfig struct {
	Id     int    `db:"id" json:"id"`         // 配置ID
	Name   string `db:"name" json:"name"`     // 配置名称
	Type   int8   `db:"type" json:"type"`     // 配置类型
	Title  string `db:"title" json:"title"`   // 配置说明
	Group  int8   `db:"group" json:"group"`   // 配置分组
	Extra  string `db:"extra" json:"extra"`   // 配置值
	Remark string `db:"remark" json:"remark"` // 配置说明
	CTime  int64  `db:"c_time" json:"c_time"` // 创建时间
	UTime  int64  `db:"u_time" json:"u_time"` // 更新时间
	Status int8   `db:"status" json:"status"` // 状态
	Value  string `db:"value" json:"value"`   // 配置值
	Sort   int8   `db:"sort" json:"sort"`     // 排序
}

var DefaultSysConfig = SysConfig{}

type sysConfigCache struct {
	objMap  map[int]*SysConfig
	objList []*SysConfig
}

var SysConfigCache = &sysConfigCache{}

func (c *sysConfigCache) LoadAll() {
	sql := "select `id`,`name`,`type`,`title`,`group`,`extra`,`remark`,`c_time`,`u_time`,`status`,`value`,`sort` from sys_config"
	c.objList = make([]*SysConfig, 0)
	err := db.JYBaseDB.Select(&c.objList, sql)
	if err != nil {
		zaplogger.Fatal(err.Error())
	}
	c.objMap = make(map[int]*SysConfig)
	for _, v := range c.objList {
		c.objMap[v.Id] = v
	}
}

func (c *sysConfigCache) All() []*SysConfig {
	return c.objList
}

func (c *sysConfigCache) Count() int {
	return len(c.objList)
}

func (c *sysConfigCache) Get(id int) (*SysConfig, bool) {
	key := id
	v, ok := c.objMap[key]
	return v, ok
}

// 仅限运营后台实时刷新服务器数据用
func (c *sysConfigCache) Update(v *SysConfig) {
	key := v.Id
	c.objMap[key] = v
}
