package services

import (
	"time"
	"carp.cn/whale/model/jy_member"
	"carp.cn/whale/zaplogger"
	"go.uber.org/zap"
	"fmt"
	"carp.cn/whale/kit"
	"encoding/json"
	"carp.cn/whale/pkg/cerr"
)

const (
	RdsUserKey       = "Whale:User:uid-%d"
	RdsUserKeyExpire = time.Hour * 24 * 2 // 用户信息保存两天
)

var DefaultUserService = NewUserService()

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func genRdsUserKey(id int) string {
	return fmt.Sprintf(RdsUserKey, id)
}

func (s *UserService) GetUserByID(id int) *JYMemberDB.User {
	u, err := s.getUserByIDFromCache(id)
	cerr.CheckErrDoNothing(err, "UserService GetUserByIDFromCache error", zap.Int(" id", id))
	if u != nil {
		return u
	}

	u, err = JYMemberDB.UserOp.Get(id)
	if err != nil {
		zaplogger.Error("UserService GetUserByIDFromDB error", zap.Error(err), zap.Int("id", id))
		return nil
	}
	err = s.saveUserToRds(u)
	if err != nil {
		zaplogger.Error("UserService SaveUserToRds error", zap.Error(err), zap.Reflect("user:", u))
	}
	return u
}

func (s *UserService) UpdateUser(user *JYMemberDB.User) error {
	err := JYMemberDB.UserOp.Update(user)
	if err != nil {
		return err
	}
	uErr := s.saveUserToRds(user)
	if uErr != nil {
		zaplogger.Error("UpdateUser to cache error", zap.Error(uErr))
		key := genRdsUserKey(user.Id)
		rErr := kit.RdsCacheHelper.Del(key)
		if rErr != nil {
			zaplogger.Error("UpdateUser del key error", zap.Error(rErr), zap.String("key", key))
		}
	}
	return nil
}

func (s *UserService) getUserByIDFromCache(id int) (*JYMemberDB.User, error) {
	key := genRdsUserKey(id)
	str, err := kit.RdsCacheHelper.Get(key)
	if err != nil {
		return nil, err
	}
	var user JYMemberDB.User
	err = json.Unmarshal([]byte(str), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil

}

func (s *UserService) saveUserToRds(user *JYMemberDB.User) error {
	key := genRdsUserKey(user.Id)
	buf, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return kit.RdsCacheHelper.Set(key, string(buf), RdsUserKeyExpire)
}
