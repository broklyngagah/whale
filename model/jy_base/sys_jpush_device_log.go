package JYBaseDB

import (
	"carp.cn/whale/db"
	"carp.cn/whale/zaplogger"
)

//This file is generate by scripts,don't edit it

//sys_jpush_device_log
//用户设备的日志记录（用来app单点登录用）

// +gen
type SysJpushDeviceLog struct {
	Id           int    `db:"id" json:"id"`                       //
	Uid          int    `db:"uid" json:"uid"`                     // 用户id
	JpushDevice  string `db:"jpush_device" json:"jpush_device"`   // 极光设备唯一标识
	XgpushDevice string `db:"xgpush_device" json:"xgpush_device"` // 信鸽推送token
	Platform     int8   `db:"platform" json:"platform"`           // 1表示ios 2.表示android
	Imei         string `db:"imei" json:"imei"`                   // 设备imei唯一标识
	Source       int8   `db:"source" json:"source"`               // 渠道来源
	AppModel     string `db:"app_model" json:"app_model"`         // 手机型号
	Status       int8   `db:"status" json:"status"`               // 登录状态1表示登录，2表示退出登录
	CTime        int64  `db:"c_time" json:"c_time"`               // 创建时间
	UTime        int64  `db:"u_time" json:"u_time"`               // 更新时间
}

var DefaultSysJpushDeviceLog = SysJpushDeviceLog{}

type sysJpushDeviceLogCache struct {
	objMap  map[int]*SysJpushDeviceLog
	objList []*SysJpushDeviceLog
}

var SysJpushDeviceLogCache = &sysJpushDeviceLogCache{}

func (c *sysJpushDeviceLogCache) LoadAll() {
	sql := "select `id`,`uid`,`jpush_device`,`xgpush_device`,`platform`,`imei`,`source`,`app_model`,`status`,`c_time`,`u_time` from sys_jpush_device_log"
	c.objList = make([]*SysJpushDeviceLog, 0)
	err := db.JYBaseDB.Select(&c.objList, sql)
	if err != nil {
		zaplogger.Fatal(err.Error())
	}
	c.objMap = make(map[int]*SysJpushDeviceLog)
	for _, v := range c.objList {
		c.objMap[v.Id] = v
	}
}

func (c *sysJpushDeviceLogCache) All() []*SysJpushDeviceLog {
	return c.objList
}

func (c *sysJpushDeviceLogCache) Count() int {
	return len(c.objList)
}

func (c *sysJpushDeviceLogCache) Get(id int) (*SysJpushDeviceLog, bool) {
	key := id
	v, ok := c.objMap[key]
	return v, ok
}

// 仅限运营后台实时刷新服务器数据用
func (c *sysJpushDeviceLogCache) Update(v *SysJpushDeviceLog) {
	key := v.Id
	c.objMap[key] = v
}
