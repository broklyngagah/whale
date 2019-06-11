package JYBaseDB

import (
	"carp.cn/whale/db"
	"carp.cn/whale/zaplogger"
)

//This file is generate by scripts,don't edit it

//sys_ad
//系统后台 - 广告

// +gen
type SysAd struct {
	Id         int    `db:"id" json:"id"`                   //
	PositionId int    `db:"position_id" json:"position_id"` // 广告位ID
	Title      string `db:"title" json:"title"`             // 广告标题
	Status     int8   `db:"status" json:"status"`           // 1.正常 2.关闭
	Imgurl     string `db:"imgurl" json:"imgurl"`           // 广告图片
	AimUrl     string `db:"aim_url" json:"aim_url"`         // 跳转链接地址
	JumpType   int8   `db:"jump_type" json:"jump_type"`     // 跳转类型 1 点击url跳转 ，2文章帖子跳转
	Uid        int    `db:"uid" json:"uid"`                 // 大V用户的UID
	OtherId    int    `db:"other_id" json:"other_id"`       // 其他id  （文章帖子id）
	STime      int64  `db:"s_time" json:"s_time"`           // 广告开始时间
	ETime      int64  `db:"e_time" json:"e_time"`           // 广告结束时间
	Sort       int8   `db:"sort" json:"sort"`               // 排序(降序)
	CTime      int64  `db:"c_time" json:"c_time"`           // 添加时间
	UTime      int64  `db:"u_time" json:"u_time"`           // 更新时间
}

var DefaultSysAd = SysAd{}

type sysAdCache struct {
	objMap  map[int]*SysAd
	objList []*SysAd
}

var SysAdCache = &sysAdCache{}

func (c *sysAdCache) LoadAll() {
	sql := "select `id`,`position_id`,`title`,`status`,`imgurl`,`aim_url`,`jump_type`,`uid`,`other_id`,`s_time`,`e_time`,`sort`,`c_time`,`u_time` from sys_ad"
	c.objList = make([]*SysAd, 0)
	err := db.JYBaseDB.Select(&c.objList, sql)
	if err != nil {
		zaplogger.Fatal(err.Error())
	}
	c.objMap = make(map[int]*SysAd)
	for _, v := range c.objList {
		c.objMap[v.Id] = v
	}
}

func (c *sysAdCache) All() []*SysAd {
	return c.objList
}

func (c *sysAdCache) Count() int {
	return len(c.objList)
}

func (c *sysAdCache) Get(id int) (*SysAd, bool) {
	key := id
	v, ok := c.objMap[key]
	return v, ok
}

// 仅限运营后台实时刷新服务器数据用
func (c *sysAdCache) Update(v *SysAd) {
	key := v.Id
	c.objMap[key] = v
}
