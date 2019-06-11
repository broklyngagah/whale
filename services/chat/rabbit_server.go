package chat

import (
	"carp.cn/whale/model/jy_member"
	"carp.cn/whale/pkg/cerr"
	"carp.cn/whale/zaplogger"
	"go.uber.org/zap"
	"carp.cn/whale/com"
	"carp.cn/whale/utils"
	"time"
	"encoding/json"
)

type RabbitLive struct {
}

func NewRabbitLive() *RabbitLive {
	return &RabbitLive{}
}

func (r *RabbitLive) UserLogin(roomID int, uid int, nickname string, imei string) {

	room := getRoomByID(roomID)
	if room == nil {
		return
	}

	t := time.Now().Unix()
	room.Echo(func(client *Client) {
		err := client.Send(com.UserLogin, map[string]interface{}{
			"uid":      uid,
			"nickname": nickname,
			"count":    getCountFromCache(roomID),
			"date":     utils.FormatDateForChat(t),
			"time":     utils.FormatTimeForChat(t),
		})
		if err != nil {
			zaplogger.Error("client write message error :", zap.Error(err), zap.Int("uid", client.User.GetUid()))
		}
	})
}

func (r *RabbitLive) UserSubscribe(roomID, uid int, nickname string) {
	room := getRoomByID(roomID)
	if room == nil {
		return
	}
	t := time.Now().Unix()
	room.Echo(func(client *Client) {
		err := client.Send(com.UserSubscribe, map[string]interface{}{
			"nickname": nickname,
			"uid":      uid,
			"date":     utils.FormatDateForChat(t),
			"time":     utils.FormatTimeForChat(t),
		})
		if err != nil {
			zaplogger.Error("client write message error :", zap.Error(err), zap.Int("uid", client.User.GetUid()))
		}
	})
}

func (r *RabbitLive) UserFollow(roomID, uid int, nickname string) {
	room := getRoomByID(roomID)
	if room == nil {
		return
	}
	t := time.Now().Unix()
	room.Echo(func(client *Client) {
		err := client.Send(com.UserFollow, map[string]interface{}{
			"nickname": nickname,
			"uid":      uid,
			"date":     utils.FormatDateForChat(t),
			"time":     utils.FormatTimeForChat(t),
		})
		if err != nil {
			zaplogger.Error("client write message error :", zap.Error(err), zap.Int("uid", client.User.GetUid()))
		}
	})
}

func (r *RabbitLive) UserSend(roomID int, msg string){
	room := getRoomByID(roomID)
	if room == nil {
		return
	}
	data := map[string]interface{}{}
	err := json.Unmarshal([]byte(msg), &data)
	if err != nil {
		zaplogger.Error("UserSend json unmarshal error", zap.Error(err), zap.String(" msg:", msg))
		return
	}

	t := time.Now().Unix()
	data["date"] = utils.FormatDateForChat(t)
	data["time"] = utils.FormatTimeForChat(t)
	room.Echo(func(client *Client) {
		err := client.Send(com.MsgRecChat, data)
		if err != nil {
			zaplogger.Error("client write message error :", zap.Error(err), zap.Int("uid", client.User.GetUid()))
		}
	})
}

func (r *RabbitLive)UserLogout(roomID int, uid int,nickname string){

	room := getRoomByID(roomID)
	if room == nil {
		return
	}

	t := time.Now().Unix()
	room.Echo(func(client *Client) {
		err := client.Send(com.UserLogout, map[string]interface{}{
			"uid":      uid,
			"nickname": nickname,
			"count":    getCountFromCache(roomID),
			"date":     utils.FormatDateForChat(t),
			"time":     utils.FormatTimeForChat(t),
		})
		if err != nil {
			zaplogger.Error("client write message error :", zap.Error(err), zap.Int("uid", client.User.GetUid()))
		}
	})
}

func getCountFromCache(roomID int) int {
	//
	return 0
}

func getRoomByID(roomID int) *Room {
	var room *Room

	room = DefaultRoomSet.GetRoomById(roomID)
	if room == nil {
		r, err := JYMemberDB.UserRoomOp.Get(roomID)
		if err != nil {
			cerr.CheckError(err, cerr.ERR_DB_QUERY)
			return nil
		}
		if r == nil {
			zaplogger.Error("not find room", zap.Int(" id:", roomID))
			return nil
		}
		room = NewRoom(r)
		DefaultRoomSet.AddRoom(room)
	}
	return room
}
