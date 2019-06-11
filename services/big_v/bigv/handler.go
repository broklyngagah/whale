package bigv

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"carp.cn/whale/zaplogger"
	"go.uber.org/zap"
	"encoding/json"
	"carp.cn/whale/pkg/rpc"
	"carp.cn/whale/kit"
)

const (
	CookieName   = "SessionID"
	VUserContext = "VUser"
)

func BigVHandler(rpc *rpc.RpcHelper, c *gin.Context) {

	defer func() {
		if err := recover(); err != nil {

		}
	}()

	fmt.Println("zhong")
	c.String(http.StatusOK, "----------")
}

func CookieMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID, err := c.Cookie(CookieName)
		if err != nil || sessionID == "" {
			zaplogger.Error("check cookie error ", zap.String("cookie_name", CookieName))
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}
		buf, err := kit.MemCacheHelper.Get(sessionID)
		if err != nil {
			zaplogger.Error("get session Info from memcache error", zap.Error(err), zap.String("session_id", sessionID))
			c.Abort()
			return
		}
		var vuser VUser
		err = json.Unmarshal(buf, &vuser)
		if err != nil {
			zaplogger.Error("CookieMiddleWare json.Unmarshal error", zap.Error(err), zap.String("data:", string(buf)))
			c.Abort()
			return
		}

		// 将VUser的信息保存至调用的上下文中
		c.Set(VUserContext, &vuser)
		c.Next()
	}
}

func CORSMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		zaplogger.Info("set cors...")
		context.Next()
	}
}
