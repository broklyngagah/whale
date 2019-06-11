package JYBaseDB

import (
	"carp.cn/whale/db"
	"carp.cn/whale/zaplogger"
)

//This file is generate by scripts,don't edit it

//sys_announce
//公告表

// +gen
type SysAnnounce struct {
	Id          int64  `db:"id" json:"id"`                   //
	Title       string `db:"title" json:"title"`             // 标题
	Type        int8   `db:"type" json:"type"`               // 分类（1系统公告，2自选公告，3协议）
	EventKey    string `db:"event_key" json:"event_key"`     // 事件key
	Keywords    string `db:"keywords" json:"keywords"`       // 关键词
	Description string `db:"description" json:"description"` // 描述
	Thumb       string `db:"thumb" json:"thumb"`             // 图片上传
	Symbol      string `db:"symbol" json:"symbol"`           // 自选公告相关股票
	Content     string `db:"content" json:"content"`         // 文章内容
	Sort        int8   `db:"sort" json:"sort"`               // 排序
	IsPush      int8   `db:"is_push" json:"is_push"`         // 是否推送1是2否
	AdminId     int    `db:"admin_id" json:"admin_id"`       // 发布者uid
	Status      int8   `db:"status" json:"status"`           // 审核状态: 1.  已审核 2. 未审核;  9 删除;
	CTime       int64  `db:"c_time" json:"c_time"`           //
	UTime       int64  `db:"u_time" json:"u_time"`           //
}

var DefaultSysAnnounce = SysAnnounce{}

type sysAnnounceCache struct {
	objMap  map[int64]*SysAnnounce
	objList []*SysAnnounce
}

var SysAnnounceCache = &sysAnnounceCache{}

func (c *sysAnnounceCache) LoadAll() {
	sql := "select `id`,`title`,`type`,`event_key`,`keywords`,`description`,`thumb`,`symbol`,`content`,`sort`,`is_push`,`admin_id`,`status`,`c_time`,`u_time` from sys_announce"
	c.objList = make([]*SysAnnounce, 0)
	err := db.JYBaseDB.Select(&c.objList, sql)
	if err != nil {
		zaplogger.Fatal(err.Error())
	}
	c.objMap = make(map[int64]*SysAnnounce)
	for _, v := range c.objList {
		c.objMap[v.Id] = v
	}
}

func (c *sysAnnounceCache) All() []*SysAnnounce {
	return c.objList
}

func (c *sysAnnounceCache) Count() int {
	return len(c.objList)
}

func (c *sysAnnounceCache) Get(id int64) (*SysAnnounce, bool) {
	key := id
	v, ok := c.objMap[key]
	return v, ok
}

// 仅限运营后台实时刷新服务器数据用
func (c *sysAnnounceCache) Update(v *SysAnnounce) {
	key := v.Id
	c.objMap[key] = v
}
