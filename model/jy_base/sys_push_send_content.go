package JYBaseDB

import (
	"carp.cn/whale/db"
	"carp.cn/whale/zaplogger"
)

//This file is generate by scripts,don't edit it

//sys_push_send_content
//推送文章内容表

// +gen
type SysPushSendContent struct {
	Id          int64  `db:"id" json:"id"`                     // ID
	Title       string `db:"title" json:"title"`               // 推送标题
	Content     string `db:"content" json:"content"`           // 推送内容
	Desc        string `db:"desc" json:"desc"`                 // 消息描述
	DeviceType  int8   `db:"device_type" json:"device_type"`   // 1ios推送 2安卓已推送   ，3表示全推送
	ContentType int8   `db:"content_type" json:"content_type"` // 1重要消息(全推) 2优质文章(全推) 3订阅更新(个推，订阅者) 4自选公告(个推，用户自选)
	PushType    int8   `db:"push_type" json:"push_type"`       // 标识推送通知跳转到的页面
	OtherId     int    `db:"other_id" json:"other_id"`         // 相关id (文章id，公告id)
	Url         string `db:"url" json:"url"`                   // 点击跳转的链接
	Status      int8   `db:"status" json:"status"`             // 审核状态:1表示审核通过，2未审核通过
	IsNow       int8   `db:"is_now" json:"is_now"`             // 是否立即推送 1立即推送，2定时推送
	IsPush      int8   `db:"is_push" json:"is_push"`           // 推送状态:1表示已经推送，2表示未推送
	Uid         int    `db:"uid" json:"uid"`                   // 管理员时后台id，用户uid
	IsAdmin     int8   `db:"is_admin" json:"is_admin"`         // 是否管理员 1是管理员，2不是管理员
	PushTime    int64  `db:"push_time" json:"push_time"`       // 定时推送时间
	OptTime     int64  `db:"opt_time" json:"opt_time"`         // 实际推送时间
	CTime       int64  `db:"c_time" json:"c_time"`             // 创建时间
	UTime       int64  `db:"u_time" json:"u_time"`             // 更新时间
}

var DefaultSysPushSendContent = SysPushSendContent{}

type sysPushSendContentCache struct {
	objMap  map[int64]*SysPushSendContent
	objList []*SysPushSendContent
}

var SysPushSendContentCache = &sysPushSendContentCache{}

func (c *sysPushSendContentCache) LoadAll() {
	sql := "select `id`,`title`,`content`,`desc`,`device_type`,`content_type`,`push_type`,`other_id`,`url`,`status`,`is_now`,`is_push`,`uid`,`is_admin`,`push_time`,`opt_time`,`c_time`,`u_time` from sys_push_send_content"
	c.objList = make([]*SysPushSendContent, 0)
	err := db.JYBaseDB.Select(&c.objList, sql)
	if err != nil {
		zaplogger.Fatal(err.Error())
	}
	c.objMap = make(map[int64]*SysPushSendContent)
	for _, v := range c.objList {
		c.objMap[v.Id] = v
	}
}

func (c *sysPushSendContentCache) All() []*SysPushSendContent {
	return c.objList
}

func (c *sysPushSendContentCache) Count() int {
	return len(c.objList)
}

func (c *sysPushSendContentCache) Get(id int64) (*SysPushSendContent, bool) {
	key := id
	v, ok := c.objMap[key]
	return v, ok
}

// 仅限运营后台实时刷新服务器数据用
func (c *sysPushSendContentCache) Update(v *SysPushSendContent) {
	key := v.Id
	c.objMap[key] = v
}
