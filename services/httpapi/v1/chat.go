package v1

import (
	"carp.cn/whale/model/jy_log"
	"carp.cn/whale/model/jy_member"
	"carp.cn/whale/utils"
	"sort"
	"carp.cn/whale/zaplogger"
	"go.uber.org/zap"
	"carp.cn/whale/pkg/cerr"
)

var DefaultChat = NewChat()

type Chat struct {
}

func NewChat() *Chat {
	return &Chat{}
}

func (c *Chat) TwentyHourlist(req map[string]interface{}) map[string]interface{} {

	cs, err := JYLogDB.ChatCountOp.QueryByClause(map[string]interface{}{}, 10, 0, []string{"views desc"}, []string{})
	cerr.CheckError(err, cerr.ERR_DB_QUERY)

	var res []map[string]interface{}

	for _, c := range cs {
		user, err := JYMemberDB.UserOp.Get(c.Uid)
		if err != nil {
			zaplogger.Error("get user error", zap.Error(err), zap.Int("uid:", c.Uid))
			user = &JYMemberDB.User{
				Id: c.Uid,
			}
		}

		remark := ""
		if uInfo, err := JYMemberDB.UserInfoOp.GetByMap(map[string]interface{}{"uid": c.Uid}); err == nil {
			remark = uInfo.Remark
		}

		info := utils.Interface2Map(c)
		info["headimg"] = user.Headimg
		info["nickname"] = user.Nickname
		info["remark"] = remark
		res = append(res, info)
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i]["sum_views"].(int) > res[j]["sum_views"].(int)
	})

	return map[string]interface{}{
		"data": res,
	}
}

//func (c *Chat) NowList(req map[string]interface{}) map[string]interface{} {
//	var page int
//	iPage, ok := req["page"]
//	if !ok {
//		page = 1
//	}
//	page = int(iPage.(float64))
//	limit := (page - 1) * 10
//	logs, err := JYLogDB.ChatMessageLogOp.QueryByClause(map[string]interface{}{}, limit, 0, []string{"u_time desc"}, []string{})
//	cerr.CheckError(err, cerr.ERR_DB_QUERY)
//
//	var res []map[string]interface{}
//
//	for _, log := range logs {
//
//		info := utils.Interface2Map(log)
//
//		u, err := JYMemberDB.UserOp.Get(log.Uid)
//		if err != nil  {
//			zaplogger.Error("get user error", zap.Error(err), zap.Int("uid:", log.Uid))
//			u = &JYMemberDB.User{
//				Id: log.Uid,
//			}
//		}
//
//		// 过滤关闭的房间
//		r, err := JYMemberDB.UserRoomOp.Get(log.RoomId)
//		if err !=nil {
//			zaplogger.Error("get user error", zap.Error(err), zap.Int("uid:", log.Uid))
//			continue
//		}
//		if r.Status != com.RoomStatusOpen {
//			continue
//		}
//
//		uFollRelation, err := JYMemberDB.UserFollowRelationOp.GetByMap(map[string]interface{}{"uid": log.Uid})
//		if err != nil {
//			uFollRelation = &JYMemberDB.UserFollowRelation{Uid: log.Uid}
//		}
//
//		info["headimg"] = u.Headimg
//		info["nickname"] = u.Nickname
//		info["fans_num"] = uFollRelation.FansNum
//
//		reply := map[string]interface{}{}
//
//		if log.ReplyId != 0 {
//			cm, err := JYLogDB.ChatMessageOp.Get(log.ReplyId)
//			if err != nil {
//				zaplogger.Error("get user error", zap.Error(err), zap.Int("uid:", log.ReplyId))
//				rUser, rErr := JYMemberDB.UserOp.Get(log.Uid)
//				if rErr != nil {
//					zaplogger.Error("get user error", zap.Error(err), zap.Int("uid:", log.Uid))
//					reply["replyNickname"] = rUser.Nickname
//				} else {
//					reply["replyNickname"] = ""
//				}
//				reply["replyContent"] = cm.Content
//				reply["replyMsgType"] = cm.MsgType
//			}
//		}
//		info["replyData"] = reply
//
//		chatCount, err := JYLogDB.ChatCountOp.GetByMap(map[string]interface{}{"uid": log.Uid})
//		if err != nil {
//			chatCount = &JYLogDB.ChatCount{Views: 0}
//		}
//
//		info["views"] = chatCount.Views
//		info["time"] = log.CTime
//		res = append(res, info)
//	}
//
//	sort.Slice(res, func(i, j int) bool {
//		return res[i]["time"].(int64) > res[j]["time"].(int64)
//	})
//
//	for k := range res {
//		t := res[k]["time"].(int64)
//		td := time.Unix(t, 0)
//		prefix := utils.FormatDateForChat(t)
//		res[k]["time_text"] = prefix + " " + td.Format("15:04")
//	}
//
//	return map[string]interface{}{
//		"data": res,
//	}
//}
