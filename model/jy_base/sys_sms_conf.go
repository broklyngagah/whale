package JYBaseDB

import (
	"carp.cn/whale/db"
	"carp.cn/whale/zaplogger"
)

//This file is generate by scripts,don't edit it

//sys_sms_conf
//短信通道配置表

// +gen
type SysSmsConf struct {
	Id       int    `db:"id" json:"id"`             //
	Title    string `db:"title" json:"title"`       // 短信通道名称
	Type     int8   `db:"type" json:"type"`         // 1短信 2语音
	Uname    string `db:"uname" json:"uname"`       // 账户名
	Upwd     string `db:"upwd" json:"upwd"`         // 密码
	Level    int8   `db:"level" json:"level"`       // 优先级(升序)
	MaxNum   int    `db:"max_num" json:"max_num"`   // 最高处理限额
	Keywords string `db:"keywords" json:"keywords"` // 通道简写关键词
	Other    string `db:"other" json:"other"`       // 其他配置信息(一维数组的json格式)
	Status   int8   `db:"status" json:"status"`     // 是否启用 1启用 9关闭
	STime    int64  `db:"s_time" json:"s_time"`     // 通道开始时间
	ETime    int64  `db:"e_time" json:"e_time"`     // 通道结束时间
	UTime    int64  `db:"u_time" json:"u_time"`     // 更新时间
}

var DefaultSysSmsConf = SysSmsConf{}

type sysSmsConfCache struct {
	objMap  map[int]*SysSmsConf
	objList []*SysSmsConf
}

var SysSmsConfCache = &sysSmsConfCache{}

func (c *sysSmsConfCache) LoadAll() {
	sql := "select `id`,`title`,`type`,`uname`,`upwd`,`level`,`max_num`,`keywords`,`other`,`status`,`s_time`,`e_time`,`u_time` from sys_sms_conf"
	c.objList = make([]*SysSmsConf, 0)
	err := db.JYBaseDB.Select(&c.objList, sql)
	if err != nil {
		zaplogger.Fatal(err.Error())
	}
	c.objMap = make(map[int]*SysSmsConf)
	for _, v := range c.objList {
		c.objMap[v.Id] = v
	}
}

func (c *sysSmsConfCache) All() []*SysSmsConf {
	return c.objList
}

func (c *sysSmsConfCache) Count() int {
	return len(c.objList)
}

func (c *sysSmsConfCache) Get(id int) (*SysSmsConf, bool) {
	key := id
	v, ok := c.objMap[key]
	return v, ok
}

// 仅限运营后台实时刷新服务器数据用
func (c *sysSmsConfCache) Update(v *SysSmsConf) {
	key := v.Id
	c.objMap[key] = v
}
