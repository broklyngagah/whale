package JYBaseDB

import (
	"carp.cn/whale/db"
	"carp.cn/whale/zaplogger"
)

//This file is generate by scripts,don't edit it

//sys_announce_view
//用户公告阅读日志

// +gen
type SysAnnounceView struct {
	Id         int64  `db:"id" json:"id"`                   //
	Uid        int    `db:"uid" json:"uid"`                 // 用户ID
	Platform   int8   `db:"platform" json:"platform"`       // 平台类型设备:（1：ios；2：android；3：wap；4：PC，5微信游戏,6  ios回馈版）
	Imei       string `db:"imei" json:"imei"`               // 手机IMEI值
	AnnounceId int64  `db:"announce_id" json:"announce_id"` // 浏览公告ID
	CTime      int64  `db:"c_time" json:"c_time"`           // 创建时间
	UTime      int64  `db:"u_time" json:"u_time"`           // 修改时间
}

var DefaultSysAnnounceView = SysAnnounceView{}

type sysAnnounceViewCache struct {
	objMap  map[int64]*SysAnnounceView
	objList []*SysAnnounceView
}

var SysAnnounceViewCache = &sysAnnounceViewCache{}

func (c *sysAnnounceViewCache) LoadAll() {
	sql := "select `id`,`uid`,`platform`,`imei`,`announce_id`,`c_time`,`u_time` from sys_announce_view"
	c.objList = make([]*SysAnnounceView, 0)
	err := db.JYBaseDB.Select(&c.objList, sql)
	if err != nil {
		zaplogger.Fatal(err.Error())
	}
	c.objMap = make(map[int64]*SysAnnounceView)
	for _, v := range c.objList {
		c.objMap[v.Id] = v
	}
}

func (c *sysAnnounceViewCache) All() []*SysAnnounceView {
	return c.objList
}

func (c *sysAnnounceViewCache) Count() int {
	return len(c.objList)
}

func (c *sysAnnounceViewCache) Get(id int64) (*SysAnnounceView, bool) {
	key := id
	v, ok := c.objMap[key]
	return v, ok
}

// 仅限运营后台实时刷新服务器数据用
func (c *sysAnnounceViewCache) Update(v *SysAnnounceView) {
	key := v.Id
	c.objMap[key] = v
}
