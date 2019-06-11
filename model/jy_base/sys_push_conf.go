package JYBaseDB

import (
	"carp.cn/whale/db"
	"carp.cn/whale/zaplogger"
)

//This file is generate by scripts,don't edit it

//sys_push_conf
//推送通道配置表

// +gen
type SysPushConf struct {
	Id       int    `db:"id" json:"id"`             //
	Title    string `db:"title" json:"title"`       // 推送通道名称
	Keywords string `db:"keywords" json:"keywords"` // 关键字
	Status   int8   `db:"status" json:"status"`     // 是否启用 1启用 9关闭
	CTime    int64  `db:"c_time" json:"c_time"`     // 创建时间
	UTime    int64  `db:"u_time" json:"u_time"`     // 更新时间
}

var DefaultSysPushConf = SysPushConf{}

type sysPushConfCache struct {
	objMap  map[int]*SysPushConf
	objList []*SysPushConf
}

var SysPushConfCache = &sysPushConfCache{}

func (c *sysPushConfCache) LoadAll() {
	sql := "select `id`,`title`,`keywords`,`status`,`c_time`,`u_time` from sys_push_conf"
	c.objList = make([]*SysPushConf, 0)
	err := db.JYBaseDB.Select(&c.objList, sql)
	if err != nil {
		zaplogger.Fatal(err.Error())
	}
	c.objMap = make(map[int]*SysPushConf)
	for _, v := range c.objList {
		c.objMap[v.Id] = v
	}
}

func (c *sysPushConfCache) All() []*SysPushConf {
	return c.objList
}

func (c *sysPushConfCache) Count() int {
	return len(c.objList)
}

func (c *sysPushConfCache) Get(id int) (*SysPushConf, bool) {
	key := id
	v, ok := c.objMap[key]
	return v, ok
}

// 仅限运营后台实时刷新服务器数据用
func (c *sysPushConfCache) Update(v *SysPushConf) {
	key := v.Id
	c.objMap[key] = v
}
