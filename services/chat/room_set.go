package chat

import (
	"carp.cn/whale/model/jy_member"
	"carp.cn/whale/utils"
	"fmt"
	"errors"
	"carp.cn/whale/com"
	"go.uber.org/zap"
	"encoding/json"
	"carp.cn/whale/pkg/rpc"
	"carp.cn/whale/pkg/cerr"
)

type Room struct {
	Info     *JYMemberDB.UserRoom // 房间信息
	UserList *utils.SafeMap       // key:client
}

func NewRoom(info *JYMemberDB.UserRoom) *Room {
	return &Room{
		Info:     info,
		UserList: utils.NewSafeMap(),
	}
}

func (r *Room) AddUser(client *Client) bool {
	return r.UserList.Set(client, struct{}{})
}

func (r *Room) DeleteUser(client *Client) {
	r.UserList.Delete(client)
}

func (r *Room) Echo(f func(client *Client)) {
	r.UserList.Echo(func(k interface{}, _ interface{}) {
		go func() {
			cli := k.(*Client)
			f(cli)
		}()
	})
}

func (r *Room) Count() int {
	return r.UserList.Count()
}

//-----------------------------------------------------------------------------------------
var DefaultRoomSet *RoomSet

type RoomSet struct {
	pair *utils.SafeMap // key: room_id(int), value: *Room
}

//-----------------------------------------------------------------------------------------
func NewRoomSet() *RoomSet {
	return &RoomSet{
		pair: utils.NewSafeMap(),
	}
}

func (rs *RoomSet) AddRoom(room *Room) {
	rs.pair.Set(room.Info.Id, room)
}

func (rs *RoomSet) GetRoomById(roomId int) *Room {
	iRoom := rs.pair.Get(roomId)
	if iRoom == nil {
		return nil
	}
	return iRoom.(*Room)
}

func (rs *RoomSet) InitRoomSet() error {
	all, err := JYMemberDB.UserRoomOp.SelectAll()
	if err != nil {
		return err
	}
	for _, r := range all {
		rs.AddRoom(NewRoom(r))
	}
	return nil
}

func (rs *RoomSet) CloseRoom(roomId int) error {
	room := rs.GetRoomById(roomId)
	if room == nil {
		return fmt.Errorf("room does not exist id:%d", roomId)
	}
	if room.Info.Status != com.RoomStatusClosed {
		room.Info.Status = com.RoomStatusClosed
		err := JYMemberDB.UserRoomOp.Update(room.Info)
		if err != nil {
			return err
		}
	}

	room.Echo(func(client *Client) {
		client.CloseFlag()
	})
	return nil
}

func (rs *RoomSet) RemoveUser(client *Client) error {
	if client.User == nil {
		return errors.New("client user is null")
	}
	if client.User.room != nil {
		client.User.room.DeleteUser(client)
	}
	return nil
}

func (rs *RoomSet) SendRoomMsg(roomId int, funcName string, data map[string]interface{}) error {

	room := rs.GetRoomById(roomId)
	if room == nil {
		return fmt.Errorf("not find room id :%d", roomId)
	}

	resp := &rpc.Response{
		Method: funcName,
		Result: data,
	}

	bytes, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	room.Echo(func(client *Client) {
		err := client.WriteMsg(bytes)
		cerr.CheckErrDoNothing(err, "client write msg error",
			zap.Int("room_id:", client.User.room.Info.Id), zap.Int("uid:", client.User.GetUid()))
	})
	return nil
}

func InitRoomSet() {
	DefaultRoomSet = NewRoomSet()
	DefaultRoomSet.InitRoomSet()
}
