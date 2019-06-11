package JYBaseDB

import (
	"carp.cn/whale/db"
	"carp.cn/whale/zaplogger"
)

//This file is generate by scripts,don't edit it

//sys_sms_tpl
//短信模板配置表

// +gen
type SysSmsTpl struct {
	Id      int    `db:"id" json:"id"`             //
	SendKey string `db:"send_key" json:"send_key"` // 短信模板KEY
	Title   string `db:"title" json:"title"`       // 短信模板标题
	SmsBody string `db:"sms_body" json:"sms_body"` // 短信模板内容
	CTime   int64  `db:"c_time" json:"c_time"`     // 创建时间
	UTime   int64  `db:"u_time" json:"u_time"`     // 最后修改时间
	OpId    int    `db:"op_id" json:"op_id"`       // 最后操作者ID
}

var DefaultSysSmsTpl = SysSmsTpl{}

type sysSmsTplCache struct {
	objMap  map[int]*SysSmsTpl
	objList []*SysSmsTpl
}

var SysSmsTplCache = &sysSmsTplCache{}

func (c *sysSmsTplCache) LoadAll() {
	sql := "select `id`,`send_key`,`title`,`sms_body`,`c_time`,`u_time`,`op_id` from sys_sms_tpl"
	c.objList = make([]*SysSmsTpl, 0)
	err := db.JYBaseDB.Select(&c.objList, sql)
	if err != nil {
		zaplogger.Fatal(err.Error())
	}
	c.objMap = make(map[int]*SysSmsTpl)
	for _, v := range c.objList {
		c.objMap[v.Id] = v
	}
}

func (c *sysSmsTplCache) All() []*SysSmsTpl {
	return c.objList
}

func (c *sysSmsTplCache) Count() int {
	return len(c.objList)
}

func (c *sysSmsTplCache) Get(id int) (*SysSmsTpl, bool) {
	key := id
	v, ok := c.objMap[key]
	return v, ok
}

// 仅限运营后台实时刷新服务器数据用
func (c *sysSmsTplCache) Update(v *SysSmsTpl) {
	key := v.Id
	c.objMap[key] = v
}
