package JYBaseDB

import (
	"carp.cn/whale/db"
	"carp.cn/whale/zaplogger"
)

//This file is generate by scripts,don't edit it

//sys_ad_position
//系统后台 - 广告位

// +gen
type SysAdPosition struct {
	Id          int    `db:"id" json:"id"`                     // 广告位ID
	PositionKey string `db:"position_key" json:"position_key"` // 广告位唯一KEY,系统自动生成
	ParentId    int    `db:"parent_id" json:"parent_id"`       // 广告位 父ID
	Title       string `db:"title" json:"title"`               // 广告位名称
	Remark      string `db:"remark" json:"remark"`             // 广告位备注
	Width       int8   `db:"width" json:"width"`               // 广告位规格之宽度
	Height      int8   `db:"height" json:"height"`             // 广告位规格之高度
	Arrange     int8   `db:"arrange" json:"arrange"`           // 排序(降序)
	Status      int8   `db:"status" json:"status"`             // 1.正常 2.关闭
}

var DefaultSysAdPosition = SysAdPosition{}

type sysAdPositionCache struct {
	objMap  map[int]*SysAdPosition
	objList []*SysAdPosition
}

var SysAdPositionCache = &sysAdPositionCache{}

func (c *sysAdPositionCache) LoadAll() {
	sql := "select `id`,`position_key`,`parent_id`,`title`,`remark`,`width`,`height`,`arrange`,`status` from sys_ad_position"
	c.objList = make([]*SysAdPosition, 0)
	err := db.JYBaseDB.Select(&c.objList, sql)
	if err != nil {
		zaplogger.Fatal(err.Error())
	}
	c.objMap = make(map[int]*SysAdPosition)
	for _, v := range c.objList {
		c.objMap[v.Id] = v
	}
}

func (c *sysAdPositionCache) All() []*SysAdPosition {
	return c.objList
}

func (c *sysAdPositionCache) Count() int {
	return len(c.objList)
}

func (c *sysAdPositionCache) Get(id int) (*SysAdPosition, bool) {
	key := id
	v, ok := c.objMap[key]
	return v, ok
}

// 仅限运营后台实时刷新服务器数据用
func (c *sysAdPositionCache) Update(v *SysAdPosition) {
	key := v.Id
	c.objMap[key] = v
}
