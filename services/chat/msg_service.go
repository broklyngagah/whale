package chat

import (
	"carp.cn/whale/pkg/mongo"
	"gopkg.in/mgo.v2/bson"
	"carp.cn/whale/pkg/cerr"
	"go.uber.org/zap"
	"carp.cn/whale/kit"
	"encoding/json"
	"fmt"
	"gopkg.in/redis.v4"
	"carp.cn/whale/utils"
	"carp.cn/whale/zaplogger"
)

const (
	_db                 = "whale"
	_chat_message_table = "chat_massage"
)

const (
	CHAT_HISTORY_KEY       = "Whale:Chat:History:Room-%d:Interact-%d" // % 房间号  % 互动
	GC_MAX_MASSAGE_COUNT   = 300                                      // 超过500条删除
	OVERFLOW_MASSAGE_COUNT = 100                                      // 默认保存条数
)

var DefaultMsgService = &MsgService{}

type MsgService struct{}

func (s *MsgService) SaveAndGC(msg *Message) error {

	if err := s.insertToDB(msg); err != nil {
		return err
	}

	cerr.CheckErrDoNothing(s.saveMessageToCache(msg), "save msg to redis error", zap.Reflect(" msg:", msg))
	cerr.CheckErrDoNothing(s.gcCacheMessage(msg.RoomId, msg.IsInteract), "gc redis error",
		zap.Int(" room_id:", msg.RoomId), zap.Int8("is_interact:", msg.IsInteract))
	return nil
}
func (op *MsgService) insertToDB(msg *Message) error {
	sess := mongo.Factory.Acquire()
	defer mongo.Factory.Release(sess)
	id, err := mongo.GetNextSequence(sess.DB(_db), _chat_message_table)
	if err != nil {
		return err
	}
	msg.Id = id
	return sess.DB(_db).C(_chat_message_table).Insert(msg)
}

func (s *MsgService) FindByID(id int) (*Message, error) {
	sess := mongo.Factory.Acquire()
	defer mongo.Factory.Release(sess)
	c := sess.DB(_db).C(_chat_message_table)
	var res *Message
	err := c.Find(bson.M{"id": id}).One(&res)
	return res, err
}

func (s *MsgService) GetMsgByCondition(id int, roomID int, isInteract int8, limit int) ([]*Message, error) {
	var (
		msgList []*Message
		err     error
	)
	if id <= 0 {
		// 最新消息直接查缓存
		msgList, err = s.getLastMessageFromCache(roomID, isInteract, limit)
	} else {
		// 老消息直接查数据库
		msgList, err = s.selectByCondition(id, roomID, isInteract, limit)
	}
	if err != nil {
		return nil, err
	}
	return msgList, nil
}

func (s *MsgService) SelectByClause(query interface{}, sort []string, limit int) ([]*Message, error) {
	sess := mongo.Factory.Acquire()
	defer mongo.Factory.Release(sess)
	c := sess.DB(_db).C(_chat_message_table)
	var res []*Message
	err := c.Find(query).Sort(sort...).Limit(limit).All(&res)
	if err != nil {
		return nil, err
	}
	return res, nil

}

func (s *MsgService) selectByCondition(id int, roomID int, IsInteract int8, limit int) ([]*Message, error) {
	sess := mongo.Factory.Acquire()
	defer mongo.Factory.Release(sess)
	c := sess.DB(_db).C(_chat_message_table)
	var res []*Message
	err := c.Find(bson.M{
		"is_interact": IsInteract,
		"room_id":     roomID,
		"id":          bson.M{"$lt": id},
	}).Sort("-id").Limit(limit).All(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *MsgService) saveMessageToCache(msg *Message) error {
	key := genChatHistoryKey(msg.RoomId, msg.IsInteract)
	buf, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	_, err = kit.RdsCacheHelper.ZAdd(key, redis.Z{Score: float64(msg.Id), Member: string(buf)})
	return err
}

func (s *MsgService) gcCacheMessage(roomID int, isInteract int8) error {
	key := genChatHistoryKey(roomID, isInteract)
	count, err := kit.RdsCacheHelper.ZCard(key)
	if err != nil {
		return err
	}
	if count > GC_MAX_MASSAGE_COUNT {
		err = kit.RdsCacheHelper.ZRemRangeByRank(key, 0, count-OVERFLOW_MASSAGE_COUNT-1)
		return err
	}
	return nil
}

func (s *MsgService) getLastMessageFromCache(roomID int, isInteract int8, count int) ([]*Message, error) {
	list, err := kit.RdsCacheHelper.ZRevRangeWithScores(genChatHistoryKey(roomID, isInteract), 0, int64(count-1))
	if err != nil {
		return nil, err
	}
	var msgs []*Message

	for _, v := range list {
		var msg Message
		str := v.Member.(string)
		err = json.Unmarshal([]byte(str), &msg)
		if err == nil {
			msgs = append(msgs, &msg)
		}
	}
	return msgs, nil
}

func (s *MsgService) GetWipeMsg(msg *Message) map[string]interface{} {
	msgSend := utils.Interface2Map(msg)
	if msg.ReplyId > 0 {
		rData, err := DefaultMsgService.FindByID(msg.ReplyId)
		if err != nil {
			zaplogger.Error("get chat message error", zap.Error(err), zap.Int("reply_id:", msg.ReplyId))
		} else {
			rMap := map[string]interface{}{
				"content":  rData.Content,
				"msg_type": rData.MsgType,
				"nickname": rData.Nickname,
			}
			msgSend["reply_data"] = rMap
		}
	}
	return msgSend
}

func genChatHistoryKey(roomID int, isInteract int8) string {
	return fmt.Sprintf(CHAT_HISTORY_KEY, roomID, isInteract)
}
