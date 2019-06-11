package bigv

import (
	"carp.cn/whale/model/jy_member"
	"carp.cn/whale/utils"
	"fmt"
)

type VUser struct {
	Base *JYMemberDB.User
	Info *JYMemberDB.UserInfo
}

func (u *VUser) UpdateVUserBase() error {
	return JYMemberDB.UserOp.Update(u.Base)
}

func (u *VUser) UpdateVUserInfo() error {
	return JYMemberDB.UserInfoOp.Update(u.Info)
}

func (u *VUser) LoadFromDB(tel string) error {
	if !utils.IsMoblieNumber(tel) {
		return fmt.Errorf("tel error:%s", tel)
	}
	user, err := JYMemberDB.UserOp.GetByMap(map[string]interface{}{
		"tel": tel,
	})
	if err != nil {
		return err
	}

	if user == nil {
		return fmt.Errorf("telephone number %s is not registered", tel)
	}
	info, err := JYMemberDB.UserInfoOp.GetByMap(map[string]interface{}{
		"uid": user.Id,
	})
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("user %d Info is no find", user.Id)
	}
	u.Base = user
	u.Info = info
	return nil
}
