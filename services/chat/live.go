package chat

import (
	"carp.cn/whale/com"
	"go.uber.org/zap"
	"time"
	"carp.cn/whale/utils"
	"carp.cn/whale/zaplogger"
	"carp.cn/whale/model/jy_member"
	"carp.cn/whale/model/jy_trade"
	"fmt"
	"sort"
	"carp.cn/whale/pkg/cerr"
	"encoding/json"
	"carp.cn/whale/services"
	"gopkg.in/mgo.v2/bson"
)

var DefaultLive = NewLive()

type Live struct {
}

func NewLive() *Live {
	return &Live{}
}

func checkUserRoomAvailable(user *User) bool {
	if user.room != nil && user.room.Info.Status != com.RoomStatusClosed {
		return true
	}
	return false
}

// 发送消息
func (l *Live) Send(c *Client, secret, isInteract, isRed, msgType int8, stocks, content string, replayId int, sign string) {

	if c.User == nil {
		return
	}
	if c.User.GetUid() == 0 {
		return
	}

	if !checkUserRoomAvailable(c.User) {
		return
	}

	isClient := int8(2)
	if c.User.room.Info.Uid == c.User.GetUid() {
		isClient = 1
	}

	t := time.Now().Unix()
	msg := &Message{
		Uid:        c.User.GetUid(),
		IsClient:   isClient,
		Secret:     secret,
		RoomId:     c.User.room.Info.Id,
		Content:    content,
		Stocks:     stocks,
		IsInteract: isInteract,
		IsRed:      isRed,
		MsgType:    msgType,
		ReplyId:    replayId,
		CTime:      t,
		Nickname:   c.User.nickname,
		Level:      c.User.level,
		HeadImg:    c.User.headImg,
		IsSub:      c.User.subscribe,
		Date:       utils.FormatDateForChat(t),
		Time:       utils.FormatTimeForChat(t),
	}

	err := DefaultMsgService.SaveAndGC(msg)

	cerr.CheckError(err, cerr.ERR_DB_INSERT)

	msgSend := DefaultMsgService.GetWipeMsg(msg)
	msgSend["sign"] = sign

	buf, err := json.Marshal(msgSend)
	cerr.CheckError(err, cerr.ERR_JSON_MARSHAL)

	BroadcastChat("UserSend", c.User.room.Info.Id, string(buf))
}

// 订阅
func (l *Live) Subscribe(c *Client) {
	if c.User.room == nil {
		return
	}

	query := map[string]interface{}{
		"uid=":         c.User.GetUid(),
		"passive_uid=": c.User.room.Info.Uid,
		"e_time>=":     time.Now().Unix(),
		"type=":        2,
	}
	subs, err := JYTradeDB.JyUserSubscribeOp.QueryByMapComparison(query)
	cerr.CheckError(err, cerr.ERR_DB_QUERY)
	if subs != nil && len(subs) > 0 {
		c.User.subscribe = int(subs[0].Type)
		BroadcastChat("UserSubscribe", c.User.room.Info.Id, c.User.GetUid(), c.User.nickname)
	}
}

// 关注
func (l *Live) Follow(c *Client) {

	if !checkUserRoomAvailable(c.User) {
		return
	}

	query := map[string]interface{}{
		"uid=":         c.User.GetUid(),
		"passive_uid=": c.User.room.Info.Uid,
		"type=":        2,
	}
	fs, err := JYMemberDB.UserFollowOp.QueryByMapComparison(query)
	cerr.CheckError(err, cerr.ERR_DB_QUERY)
	if fs != nil && len(fs) > 0 {
		c.User.subscribe = int(fs[0].Type)
		BroadcastChat("UserFollow", c.User.room.Info.Id, c.User.GetUid(), c.User.nickname)
	}
}

// 获取互动历史消息
func (l *Live) InteractHistory(c *Client, page int, chatID int) map[string]interface{} {
	return l.History(c, 1, page, chatID)
}

// 获取非互动历史消息
func (l *Live) NoInteractHistory(c *Client, page int, chatID int) map[string]interface{} {
	return l.History(c, 2, page, chatID)
}

// 获取历史消息(先从缓存中查找找不到再去数据库中查找， 将从数据库中查找到的消息保存至缓存中 并设定时间自动清楚)
func (l *Live) History(c *Client, isInteract int8, page int, chatID int) map[string]interface{} {
	res := map[string]interface{}{}

	if page > 3 {
		cerr.RaiseErrorCode(cerr.ERR_PARAM_VALUE, "page > 3")
	}

	limit := 10

	if page > 0 {
		limit = 10 * page
	}

	if !checkUserRoomAvailable(c.User) {
		return res
	}

	msgList, err := DefaultMsgService.GetMsgByCondition(chatID, c.User.room.Info.Id, isInteract, limit)
	cerr.CheckError(err, cerr.ERR_DB_QUERY)

	res["chat_list"] = l.getMsgByBaseMsg(msgList)
	return res
}

func (l *Live) getMsgByBaseMsg(msgs []*Message) []map[string]interface{} {
	var chatList []map[string]interface{}
	for _, msg := range msgs {
		chatList = append(chatList, DefaultMsgService.GetWipeMsg(msg))
	}

	sort.Slice(chatList, func(i, j int) bool {
		return chatList[i]["id"].(int) < chatList[j]["id"].(int)
	})
	return chatList
}

// 进入搜索界面获取最新的互动消息
func (l *Live) InteractSearchLast(c *Client, page int, chatID int) map[string]interface{} {
	return l.SearchLast(c, 1, page, chatID)
}

// 进入搜索界面获取最新的直播消息
func (l *Live) NoInteractSearchLast(c *Client, page int, chatID int) map[string]interface{} {
	return l.SearchLast(c, 2, page, chatID)
}

func (l *Live) SearchLast(c *Client, isInteract int8, page int, chatID int) map[string]interface{} {
	res := map[string]interface{}{}

	if page > 3 {
		cerr.RaiseErrorCode(cerr.ERR_PARAM_VALUE, "page > 3")
	}
	limit := 10
	if page > 0 {
		limit = 10 * page
	}

	if !checkUserRoomAvailable(c.User) {
		return res
	}

	msgList, err := DefaultMsgService.GetMsgByCondition(chatID, c.User.room.Info.Id, isInteract, limit)
	cerr.CheckError(err, cerr.ERR_DB_QUERY)

	res["chat_list"] = l.getMsgByBaseMsg(msgList)

	return res
}

// 根据关键字 时间段 搜索互动历史消息
func (l *Live) InteractSearch(c *Client, key string, begin, end int64) map[string]interface{} {
	return l.Search(c, 1, key, begin, end)
}

// 根据关键字 时间段 搜索直播历史消息
func (l *Live) NoInteractSearch(c *Client, key string, begin, end int64) map[string]interface{} {
	return l.Search(c, 2, key, begin, end)
}

func (l *Live) Search(c *Client, isInteract int8, key string, begin, end int64) map[string]interface{} {
	if begin > end {
		cerr.RaiseErrorCode(cerr.ERR_PARAM_VALUE, "begin time > end time")
	}
	if key == "" {
		cerr.RaiseErrorCode(cerr.ERR_PARAM_VALUE, "key is null")
	}

	res := map[string]interface{}{}

	msgs, err := DefaultMsgService.SelectByClause(bson.M{
		"is_interact": isInteract,
		"room_id":     c.User.room.Info.Id,
		"c_time":      bson.M{"$gt": begin, "$lte": end},
		"content":     bson.M{"$regex": fmt.Sprintf("/%s/", key)},
	}, []string{"-id"}, 10)

	cerr.CheckError(err, cerr.ERR_DB_QUERY)

	res["chat_list"] = l.getMsgByBaseMsg(msgs)

	return res
}

func GetPageSize(i int) int {
	res := i / 20
	if i%20 > 0 {
		res ++
	}
	return res
}

// imei 为设备号 为了区分游客
func Login(c *Client, uid int, roomID int, imei string) bool {

	var user *JYMemberDB.User

	if uid != 0 {
		user = services.DefaultUserService.GetUserByID(uid)
		if user == nil {
			return false
		}
	} else {
		user = &JYMemberDB.User{
			Id:       0,
			Nickname: com.VISITOR,
		}
	}

	var room *Room
	room = DefaultRoomSet.GetRoomById(roomID)
	if room == nil {
		r, err := JYMemberDB.UserRoomOp.Get(roomID)
		if err != nil {
			zaplogger.Error("login JYMemberDB.UserRoomOp.Get error", zap.Error(err))
			return false
		}
		if r == nil {
			zaplogger.Error("not find room", zap.Int(" id:", roomID))
			return false
		}
		room = NewRoom(r)
		DefaultRoomSet.AddRoom(room)
	}

	if room.Info.Status == com.RoomStatusClosed {
		zaplogger.Error("EnterRoom room closed", zap.Int("room id:", room.Info.Id))
		return false
	}

	if !room.AddUser(c) {
		zaplogger.Error("EnterRoom add user fail", zap.String("IMEI duplicate:", imei))
		return false
	}

	// 订阅赋值
	if uid != 0 {
		c.User.subscribe = getSubScribeType(uid, room.Info.Uid)
		c.User.follow = getFollowType(uid, room.Info.Uid)
	}

	c.User.SetUid(user.Id)
	c.User.room = room
	c.User.imei = imei
	c.User.nickname = user.Nickname
	c.User.headImg = user.Headimg
	c.User.level = user.Level
	c.User.isLogin = true
	return true
}

func getSubScribeType(uid, targetID int) int {
	query := map[string]interface{}{
		"uid=":         uid,
		"passive_uid=": targetID,
		"type=":        2,
	}

	subs, err := JYTradeDB.JyUserSubscribeOp.QueryByMapComparison(query)
	if err != nil {
		zaplogger.Error("EnterRoom select Subscribe from db error",
			zap.Error(err),
			zap.Reflect("query:", query))
	}
	if subs != nil && len(subs) > 0 {
		return int(subs[0].Type)
	}
	return 0

}

func getFollowType(uid, targetID int) int {
	query := map[string]interface{}{
		"uid=":         uid,
		"passive_uid=": targetID,
		"type=":        2,
	}
	fs, err := JYMemberDB.UserFollowOp.QueryByMapComparison(query)
	if err != nil {
		zaplogger.Error("EnterRoom select follow from db error",
			zap.Error(err),
			zap.Reflect("query:", query))
	}
	if fs != nil && len(fs) > 0 {
		return int(fs[0].Type)
	}
	return 0
}

// 房间关闭
func CloseRoom(roomId int) error {
	room := DefaultRoomSet.GetRoomById(roomId)
	if room == nil {
		return fmt.Errorf("not find room by id:%d", roomId)
	}
	if room.Info.Status == com.RoomStatusClosed {
		return fmt.Errorf("room has been closed id:%d", roomId)
	}
	room.Info.Status = com.RoomStatusClosed
	err := JYMemberDB.UserRoomOp.Update(room.Info)
	if err != nil {
		zaplogger.Error("CloseRoom JYMemberDB.UserRoomOp.Update error:", zap.Error(err), zap.Reflect(" UserRoom:", room.Info))
		room.Info.Status = com.RoomStatusOpen
		return err
	}

	room.Echo(func(client *Client) {
		err := client.Send(com.RoomClose, map[string]interface{}{})
		if err != nil {
			zaplogger.Error("CloseRoom client.WriteMsg error:", zap.Error(err), zap.Int(" uid", client.User.GetUid()))
		}
	})

	room.Echo(func(client *Client) {
		time.Sleep(time.Millisecond * 100) // 保证客户端能收到连接断开的消息
		client.CloseFlag()
	})
	return nil
}
