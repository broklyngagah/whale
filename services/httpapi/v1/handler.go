package v1

import (
	"github.com/gin-gonic/gin/binding"
	"carp.cn/whale/com"
	"go.uber.org/zap"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"time"
	"net/http"
	"carp.cn/whale/zaplogger"
	"carp.cn/whale/pkg/rpc"
	"carp.cn/whale/pkg/cerr"
	"carp.cn/whale/utils"
)

func ApiHandler(rpcHelper *rpc.RpcHelper, c *gin.Context) {
	var base com.BaseArg
	resp := &com.ApiResponse{
		Data:[]interface{}{},
		Time:time.Now().Format(com.RESPONSE_TIME_LAYOUT),
	}
	defer func() {
		if err := recover(); err != nil {
			comErr := cerr.ErrorWipe(err)
			resp.Code = comErr.Code
			resp.Msg = comErr.Msg
		} else {
			resp.Code = cerr.RESPONSE_SUCCESSS
			resp.Msg = "成功"
		}

		enData, err := resp.GetEncryptData(base.EnKey)
		if err != nil {
			zaplogger.Info("HTTP Response", zap.Reflect("[OUT]:", "null string"))
			c.String(http.StatusOK, "")
		} else {
			zaplogger.Info("HTTP Response", zap.Reflect("[OUT]:", resp))
			c.JSON(http.StatusOK, map[string]interface{}{"en_data": enData})
		}
		return
	}()


	err := c.ShouldBindWith(&base, binding.Form)
	cerr.CheckError(err, cerr.ERR_BINDING)
	zaplogger.Info("HTTP Base data", zap.Reflect("[IN]:", &base))

	// Decrypt
	req, err := (&base).GetData()
	cerr.CheckError(err, cerr.ERR_DECRYPT)
	zaplogger.Info("HTTP Request [IN]:" + string(req))

	param := map[string]interface{}{}
	cerr.CheckError(json.Unmarshal(req, &param), cerr.ERR_JSON_UNMARSHAL)


	result := rpcHelper.Handler(&rpc.Request{
		Method: utils.UrlFromate(base.OptAct),
		Params: []interface{}{param},
	})

	cerr.CheckIError(result.Error)

	ResponseHandler(result, resp)
}

func ResponseHandler(resp *rpc.Response, apiResp *com.ApiResponse) {
	if resp != nil && resp.Result != nil {
		// TODO: 后期统一用map 传值
		switch v := resp.Result.(type) {
		case []interface{}:
			apiResp.Data = v
		case map[string]interface{}:
			apiResp.Data = v["data"]
		default:
			apiResp.Data = v
		}
	}
}
