package v1

import (
	"carp.cn/whale/utils"
	"fmt"
	"carp.cn/whale/model/jy_member"
	"carp.cn/whale/pkg/cerr"
)

var DefaultReg = NewReg()

type Reg struct {
}

func NewReg() *Reg {
	return &Reg{}
}

func (r *Reg) Sendsms(req map[string]interface{}) map[string]interface{} {

	//TODO:完善
	return map[string]interface{}{}
}

func (r *Reg) Login(req map[string]interface{}) map[string]interface{} {
	tel, ok := req["tel"].(string)
	if !ok {
		cerr.RaiseErrorCode(cerr.ERR_PHONE_NUMBER, fmt.Sprintf("Error phone number type:%v", req["tel"]))
	}
	if ok := utils.IsMoblieNumber(tel); !ok {
		cerr.RaiseErrorCode(cerr.ERR_PHONE_NUMBER, fmt.Sprintf("Error phone number:%s", tel))
	}

	smsCode, ok := req["sms_code"].(string)
	if !ok {
		cerr.RaiseErrorCode(cerr.ERR_SMS_CODE, fmt.Sprintf("Error SMS code type:%v", req["sms_code"]))
	}

	fmt.Println(tel, smsCode)

	// TODO: user register and login
	user := defaultUserServer.FindUserByTel(tel)
	if user != nil {
		defaultUserServer.Login()
	} else {
		user = defaultUserServer.Register()
	}

	// 每日奖励
	defaultDailyServer.DailyLoginReward(user)

	res := utils.Interface2Map(user)

	userInfo, err := JYMemberDB.UserInfoOp.GetByMap(map[string]interface{}{"uid": user.Id})
	cerr.CheckError(err, cerr.ERR_DB_QUERY)

	for k, v := range utils.Interface2Map(userInfo) {
		if k == "id" {
			continue
		}
		res[k] = v
	}
	return map[string]interface{}{
		"data": res,
	}
}
