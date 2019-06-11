package JYBaseDB

import (
	"carp.cn/whale/db"
	"carp.cn/whale/zaplogger"
)

//This file is generate by scripts,don't edit it

//sys_auth_rule
//系统后台 - 菜单表

// +gen
type SysAuthRule struct {
	Id        int    `db:"id" json:"id"`               //
	Name      string `db:"name" json:"name"`           // 路径名称
	Param     string `db:"param" json:"param"`         // 其他参数
	Title     string `db:"title" json:"title"`         // 菜单名称
	Type      int8   `db:"type" json:"type"`           // 类型
	Status    int8   `db:"status" json:"status"`       // 状态
	Css       string `db:"css" json:"css"`             // 样式
	Condition string `db:"condition" json:"condition"` //
	Pid       int    `db:"pid" json:"pid"`             // 父栏目ID
	Sort      int    `db:"sort" json:"sort"`           // 排序
	CTime     int64  `db:"c_time" json:"c_time"`       // 添加时间
	UTime     int64  `db:"u_time" json:"u_time"`       // 更新时间
}

var DefaultSysAuthRule = SysAuthRule{}

type sysAuthRuleCache struct {
	objMap  map[int]*SysAuthRule
	objList []*SysAuthRule
}

var SysAuthRuleCache = &sysAuthRuleCache{}

func (c *sysAuthRuleCache) LoadAll() {
	sql := "select `id`,`name`,`param`,`title`,`type`,`status`,`css`,`condition`,`pid`,`sort`,`c_time`,`u_time` from sys_auth_rule"
	c.objList = make([]*SysAuthRule, 0)
	err := db.JYBaseDB.Select(&c.objList, sql)
	if err != nil {
		zaplogger.Fatal(err.Error())
	}
	c.objMap = make(map[int]*SysAuthRule)
	for _, v := range c.objList {
		c.objMap[v.Id] = v
	}
}

func (c *sysAuthRuleCache) All() []*SysAuthRule {
	return c.objList
}

func (c *sysAuthRuleCache) Count() int {
	return len(c.objList)
}

func (c *sysAuthRuleCache) Get(id int) (*SysAuthRule, bool) {
	key := id
	v, ok := c.objMap[key]
	return v, ok
}

// 仅限运营后台实时刷新服务器数据用
func (c *sysAuthRuleCache) Update(v *SysAuthRule) {
	key := v.Id
	c.objMap[key] = v
}
