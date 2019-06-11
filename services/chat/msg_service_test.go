package chat

import (
	"carp.cn/whale/config"
	"carp.cn/whale/zaplogger"
	"carp.cn/whale/db"
	"carp.cn/whale/kit"
	"testing"
	"time"
	"carp.cn/whale/utils"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	config.LoadFromFile("../../config.json")

	zaplogger.SetLogger("../../logs", "test.go", "debug", true)

	db.InitDB()
	kit.Init()
}

var msg = &Message{
	Uid:        12,
	IsClient:   1,
	Secret:     2,
	RoomId:     13,
	Content:    "内容 发图片就存图片地址",
	Stocks:     "$10034$",
	IsInteract: 2,
	IsRed:      2,
	MsgType:    1,
	ReplyId:    0,
	CTime:      time.Now().Unix(),
	Nickname:   "snlan",
	Level:      45,
	HeadImg:    "/base/var/1.jpg",
	IsSub:      1,
	Date:       utils.FormatDateForChat(time.Now().Unix()),
	Time:       utils.FormatTimeForChat(time.Now().Unix()),
}

func TestMsgService_SaveAndGC(t *testing.T) {
	err := DefaultMsgService.SaveAndGC(msg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", msg)
}

func TestMsgService_saveToDB(t *testing.T) {
	err := DefaultMsgService.insertToDB(msg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", msg)
}

func TestMsgService_FindByID(t *testing.T) {
	msg, err := DefaultMsgService.FindByID(1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", msg)
}

func TestMsgService_GetMsgByCondition(t *testing.T) {
	msgs, err := DefaultMsgService.GetMsgByCondition(0, 13, 2, 10)
	if err != nil {
		panic(err)
	}
	for _, msg := range msgs {
		fmt.Printf("%+v\n", msg)
	}
}

func TestMsgService_SelectByClause(t *testing.T) {
	key := fmt.Sprintf("/%s/", "ty")
	fmt.Println(key)
	msgs, err := DefaultMsgService.SelectByClause(bson.M{
		"is_interact": 2,
		"room_id":     13,
		"c_time":      bson.M{"$gt": 0, "$lte": 1524567586},
		"content":     bson.M{"$regex": fmt.Sprintf("/%s/", key), },
	}, []string{"-id"}, 10)
	if err != nil {
		panic(err)
	}
	for _, msg := range msgs {
		fmt.Printf("%+v\n", msg)
	}
	fmt.Println(len(msgs))
}
