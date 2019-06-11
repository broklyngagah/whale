package v1

import "carp.cn/whale/model/jy_member"

var defaultDailyServer = NewDailyServer()

// 每日任务处理
type DailyServer struct{}

func NewDailyServer() *DailyServer {
	return &DailyServer{}
}

// TODO : 每日登陆奖励
func (d *DailyServer) DailyLoginReward(user *JYMemberDB.User) {}
