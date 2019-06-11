package v1

import (
	"carp.cn/whale/model/jy_member"
	"fmt"
	"carp.cn/whale/pkg/cerr"
)

var defaultUserServer = NewUserServer()

type UserServer struct{}

func NewUserServer() *UserServer {
	return &UserServer{}
}

func (u *UserServer) Register() *JYMemberDB.User {

	return &JYMemberDB.User{}
}

func (u *UserServer) Login() {}

func (u *UserServer) FindUserByTel(phone string) *JYMemberDB.User {
	user, err := JYMemberDB.UserOp.GetByMap(map[string]interface{}{"tel": phone})
	if err != nil {
		cerr.CheckErrDoNothing(err, fmt.Sprintf("find user by tel error, tel:%s", phone))
		return nil
	}
	return user
}

func (u *UserServer) FindUserById(id int) *JYMemberDB.User {
	user, err := JYMemberDB.UserOp.Get(id)
	if err != nil {
		cerr.CheckErrDoNothing(fmt.Errorf("find user by id error, id:%d", id), "")
		return nil
	}
	return user
}
