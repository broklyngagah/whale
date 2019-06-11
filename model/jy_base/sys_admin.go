package JYBaseDB

import (
	"carp.cn/whale/db"
	"carp.cn/whale/zaplogger"
)

//This file is generate by scripts,don't edit it

//sys_admin
//系统后台 - 帐户

// +gen
type SysAdmin struct {
	Id            int    `db:"id" json:"id"`                           //
	AdminName     string `db:"admin_name" json:"admin_name"`           // 用户名
	Password      string `db:"password" json:"password"`               // 密码
	Portrait      string `db:"portrait" json:"portrait"`               // 头像
	Loginnum      int    `db:"loginnum" json:"loginnum"`               // 登陆次数
	LastLoginIp   string `db:"last_login_ip" json:"last_login_ip"`     // 最后登录IP
	LastLoginTime int64  `db:"last_login_time" json:"last_login_time"` // 最后登录时间
	RealName      string `db:"real_name" json:"real_name"`             // 真实姓名
	Status        int    `db:"status" json:"status"`                   // 状态：1.正常  2.锁定  3.关闭 4.异常
	Groupid       int    `db:"groupid" json:"groupid"`                 // 用户角色id
	ErrorNum      int8   `db:"error_num" json:"error_num"`             // 错误次数大于3次，帐户锁定
	LockTime      int64  `db:"lock_time" json:"lock_time"`             // 锁定时间
}

var DefaultSysAdmin = SysAdmin{}

type sysAdminCache struct {
	objMap  map[int]*SysAdmin
	objList []*SysAdmin
}

var SysAdminCache = &sysAdminCache{}

func (c *sysAdminCache) LoadAll() {
	sql := "select `id`,`admin_name`,`password`,`portrait`,`loginnum`,`last_login_ip`,`last_login_time`,`real_name`,`status`,`groupid`,`error_num`,`lock_time` from sys_admin"
	c.objList = make([]*SysAdmin, 0)
	err := db.JYBaseDB.Select(&c.objList, sql)
	if err != nil {
		zaplogger.Fatal(err.Error())
	}
	c.objMap = make(map[int]*SysAdmin)
	for _, v := range c.objList {
		c.objMap[v.Id] = v
	}
}

func (c *sysAdminCache) All() []*SysAdmin {
	return c.objList
}

func (c *sysAdminCache) Count() int {
	return len(c.objList)
}

func (c *sysAdminCache) Get(id int) (*SysAdmin, bool) {
	key := id
	v, ok := c.objMap[key]
	return v, ok
}

// 仅限运营后台实时刷新服务器数据用
func (c *sysAdminCache) Update(v *SysAdmin) {
	key := v.Id
	c.objMap[key] = v
}
