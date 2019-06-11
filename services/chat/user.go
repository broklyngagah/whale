package chat

import "sync"

type User struct {
	uid       int    // user id
	nickname  string // 昵称
	headImg   string // 头像
	level     int8   // 等级
	imei      string // IMEI
	subscribe int    // 订阅类型
	follow    int    // 关注类型
	room      *Room  // room
	isLogin   bool   // 是否已经登录
	loginTime int64  // 登录时间
	sync.Mutex
}


func (u *User)GetUid() int {
	u.Lock()
	defer u.Unlock()
	return u.uid
}

func (u *User)SetUid(uid int) {
	u.Lock()
	defer u.Unlock()
	u.uid = uid
}


