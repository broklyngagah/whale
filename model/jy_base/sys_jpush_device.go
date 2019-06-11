package JYBaseDB

import (
	"carp.cn/whale/db"
	"carp.cn/whale/zaplogger"
)

//This file is generate by scripts,don't edit it

//sys_jpush_device
//用户设备唯一标识

// +gen
type SysJpushDevice struct {
	Id           int    `db:"id" json:"id"`                       //
	Uid          int    `db:"uid" json:"uid"`                     // 用户id
	JpushDevice  string `db:"jpush_device" json:"jpush_device"`   // 极光注册ID
	XgpushDevice string `db:"xgpush_device" json:"xgpush_device"` // 信鸽推送token
	Imei         string `db:"imei" json:"imei"`                   // 设备imei唯一标识
	Platform     int8   `db:"platform" json:"platform"`           // 1表示ios 2.表示android（最后一次登录的设备来源）
	CTime        int64  `db:"c_time" json:"c_time"`               // 创建时间
	UTime        int64  `db:"u_time" json:"u_time"`               // 更新时间
}

var DefaultSysJpushDevice = SysJpushDevice{}

type sysJpushDeviceCache struct {
	objMap  map[int]*SysJpushDevice
	objList []*SysJpushDevice
}

var SysJpushDeviceCache = &sysJpushDeviceCache{}

func (c *sysJpushDeviceCache) LoadAll() {
	sql := "select `id`,`uid`,`jpush_device`,`xgpush_device`,`imei`,`platform`,`c_time`,`u_time` from sys_jpush_device"
	c.objList = make([]*SysJpushDevice, 0)
	err := db.JYBaseDB.Select(&c.objList, sql)
	if err != nil {
		zaplogger.Fatal(err.Error())
	}
	c.objMap = make(map[int]*SysJpushDevice)
	for _, v := range c.objList {
		c.objMap[v.Id] = v
	}
}

func (c *sysJpushDeviceCache) All() []*SysJpushDevice {
	return c.objList
}

func (c *sysJpushDeviceCache) Count() int {
	return len(c.objList)
}

func (c *sysJpushDeviceCache) Get(id int) (*SysJpushDevice, bool) {
	key := id
	v, ok := c.objMap[key]
	return v, ok
}

// 仅限运营后台实时刷新服务器数据用
func (c *sysJpushDeviceCache) Update(v *SysJpushDevice) {
	key := v.Id
	c.objMap[key] = v
}
