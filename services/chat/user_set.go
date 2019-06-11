package chat

import (
	"carp.cn/whale/utils"
)

var DefaultUserSet *UserSet

func NewUserSet() *UserSet {
	return &UserSet{
		online: utils.NewSafeMap(),
	}
}

type UserSet struct {
	online *utils.SafeMap
}

func (u *UserSet) AddClient(client *Client) bool {
	return u.online.Set(client.User.GetUid(), client)
}

func (u *UserSet) GetClient(uid int) *Client {
	return u.online.Get(uid).(*Client)
}

func (u *UserSet) GetClientCnt() int {
	return u.online.Count()
}

func (u *UserSet) GetClients() []*Client {
	var list []*Client
	for _, c := range u.online.Items() {
		list = append(list, c.(*Client))
	}
	return list
}

func (u *UserSet) KickOutClient(uid int) {
	client := u.GetClient(uid)
	u.RemoveClient(uid)
	if client != nil {
		client.Close()
	}
}

func (u *UserSet) RemoveClient(uid int) {
	u.online.Delete(uid)
}

func (u *UserSet) Echo (f func (client *Client) bool) {
	u.online.Echo(func(k interface{}, v interface{}) {
		f(v.(*Client))
	})
}

func InitUserSet() {
	DefaultUserSet = NewUserSet()
}
