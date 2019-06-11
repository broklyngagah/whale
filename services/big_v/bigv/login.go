package bigv

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"carp.cn/whale/zaplogger"
	"go.uber.org/zap"
	"time"
	"carp.cn/whale/utils"
	"strconv"
	"carp.cn/whale/model/jy_member"
	"carp.cn/whale/pkg/rpc"
	"carp.cn/whale/kit"
)

const (
	AUTH_SMS_KEY = "Auth:SMS:%s"
	SMS_LENGHT   = 6
	SMS_EXPIRE   = time.Minute * 10
)

func genRdsSmsKey(tel string) string {
	return fmt.Sprintf(AUTH_SMS_KEY, tel)
}

func setSmsKey(tel string, sms string) error {
	key := genRdsSmsKey(tel)
	return kit.RdsCacheHelper.Set(key, sms, SMS_EXPIRE)
}

func getSmsKey(tel string) (string, error) {
	return kit.RdsCacheHelper.Get(tel)
}

//----------------------------------------------------------------------------------
func SendSms(tel string, sms string) {

}

//----------------------------------------------------------------------------------
func BigVSms(c *gin.Context) {
	fmt.Println("BigSms")
	var req rpc.Request
	err := c.Bind(&req)
	if err != nil {
		zaplogger.Error("BigVSms c.Bind error", zap.Error(err))
		c.Status(http.StatusBadRequest)
		return
	}

	if req.Method != "BigSms" {
		zaplogger.Error("BigVSms request method error", zap.String(" method:", req.Method))
		c.Status(http.StatusBadRequest)
		return
	}

	if len(req.Params) < 1 {
		zaplogger.Error("BigVSms request params error", zap.Reflect("params:", req.Params))
		c.String(http.StatusOK, "请输入电话号码")
		return
	}
	tel, ok := req.Params[0].(string)
	if !ok {
		zaplogger.Error("BigVSms request params[0] type error", zap.Reflect("params[0]:", req.Params[0]))
		c.Status(http.StatusBadRequest)
		return
	}

	if !utils.IsMoblieNumber(tel) {
		zaplogger.Error("BigVSms request tel error", zap.String("tel:", tel))
		c.String(http.StatusOK, "电话号码错误")
		return
	}

	// 将短息验证码保存至redis
	sms := strconv.Itoa(utils.GenRandDigits(SMS_LENGHT))
	// 发送短信
	go SendSms(tel, sms)

	err = setSmsKey(tel, sms)
	if err != nil {
		zaplogger.Error("BigVSms save sms to redis error", zap.Error(err), zap.String("tel:", tel), zap.String("sms:", sms))
	}
	c.Status(http.StatusOK)
}

func BigVLogin(c *gin.Context) {

	var req rpc.Request
	err := c.Bind(&req)
	if err != nil {
		zaplogger.Error("BigVLogin c.Bind error", zap.Error(err))
		c.Status(http.StatusBadRequest)
		return
	}

	if req.Method != "BigVLogin" {
		zaplogger.Error("BigVLogin request method error", zap.String(" method:", req.Method))
		c.Status(http.StatusBadRequest)
		return
	}

	if len(req.Params) < 2 {
		zaplogger.Error("BigVLogin request params error", zap.Reflect("params:", req.Params))
		c.String(http.StatusOK, "请输入电话号码以及验证码")
		return
	}
	tel, ok := req.Params[0].(string)
	if !ok {
		zaplogger.Error("BigVLogin request params[0] type error", zap.Reflect("params[0]:", req.Params[0]))
		c.Status(http.StatusBadRequest)
		return
	}

	if !utils.IsMoblieNumber(tel) {
		zaplogger.Error("BigVLogin request tel error", zap.String("tel:", tel))
		c.String(http.StatusOK, "电话号码错误")
		return
	}

	sms, ok := req.Params[1].(string)
	if !ok {
		zaplogger.Error("BigVLogin request params[1] type error", zap.Reflect("params[1]:", req.Params[1]))
		c.Status(http.StatusBadRequest)
		return
	}

	smsAuth, err := getSmsKey(tel)
	if err != nil {
		zaplogger.Error("BigVLogin get sms from redis error", zap.String("tel:", tel))
		c.String(http.StatusOK, "请求失败")
		return
	}
	if sms != smsAuth {
		zaplogger.Error("BigVLogin sms is no match",
			zap.String("tel:", tel), zap.String("from client msg:%s", sms), zap.String("must be:%s", smsAuth))
		c.String(http.StatusOK, "验证码错误")
		return
	}

	user, err := JYMemberDB.UserOp.GetByMap(map[string]interface{}{
		"tel": tel,
	})
	if err != nil {
		zaplogger.Error("BigVLogin JYMemberDB.UserOp.GetByMap error", zap.Error(err), zap.String("tel", tel))
		c.String(http.StatusOK, "请求失败")
		return
	}

	if user == nil {
		zaplogger.Error("BigVLogin JYMemberDB.UserOp.GetByMap user is null", zap.String("tel", tel))
		c.String(http.StatusOK, " 请是使用移动端注册")
		return
	}
	c.JSON(http.StatusOK, req)

}
