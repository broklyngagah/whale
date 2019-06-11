package JYBaseDB

import (
	"carp.cn/whale/db"
	"carp.cn/whale/zaplogger"
)

//This file is generate by scripts,don't edit it

//sys_app_version
//APP版本控制配置表

// +gen
type SysAppVersion struct {
	Id         int    `db:"id" json:"id"`                   //
	Platform   int8   `db:"platform" json:"platform"`       // 1:ios,2安卓
	Version    string `db:"version" json:"version"`         // 内部版本号
	VerNum     string `db:"ver_num" json:"ver_num"`         // app外部版本号
	ImgUrl     string `db:"img_url" json:"img_url"`         // 升级图片
	UpdateUrl  string `db:"update_url" json:"update_url"`   // 升级地址
	Desc       string `db:"desc" json:"desc"`               // 升级描述
	IsUpdate   int8   `db:"is_update" json:"is_update"`     // 是否提示弹窗 1提示 2不提示
	MustUpdate int8   `db:"must_update" json:"must_update"` // 强制升级  1 强制，2不强制
	Status     int8   `db:"status" json:"status"`           // 审核, 1为审核，2为未审核
	CTime      int64  `db:"c_time" json:"c_time"`           // 创建时间
}

var DefaultSysAppVersion = SysAppVersion{}

type sysAppVersionCache struct {
	objMap  map[int]*SysAppVersion
	objList []*SysAppVersion
}

var SysAppVersionCache = &sysAppVersionCache{}

func (c *sysAppVersionCache) LoadAll() {
	sql := "select `id`,`platform`,`version`,`ver_num`,`img_url`,`update_url`,`desc`,`is_update`,`must_update`,`status`,`c_time` from sys_app_version"
	c.objList = make([]*SysAppVersion, 0)
	err := db.JYBaseDB.Select(&c.objList, sql)
	if err != nil {
		zaplogger.Fatal(err.Error())
	}
	c.objMap = make(map[int]*SysAppVersion)
	for _, v := range c.objList {
		c.objMap[v.Id] = v
	}
}

func (c *sysAppVersionCache) All() []*SysAppVersion {
	return c.objList
}

func (c *sysAppVersionCache) Count() int {
	return len(c.objList)
}

func (c *sysAppVersionCache) Get(id int) (*SysAppVersion, bool) {
	key := id
	v, ok := c.objMap[key]
	return v, ok
}

// 仅限运营后台实时刷新服务器数据用
func (c *sysAppVersionCache) Update(v *SysAppVersion) {
	key := v.Id
	c.objMap[key] = v
}
