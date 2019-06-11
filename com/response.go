package com

import (
	"github.com/gin-gonic/gin/json"
	"carp.cn/whale/utils"
)

type ApiResponse struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
	//URLCode string      `json:"url_code"`
	Time    string      `json:"time"`
}

func (r *ApiResponse) GetEncryptData(EnKey string) (string, error) {
	buf, err := json.Marshal(r)
	if err != nil {
		return "" , err
	}

	enData, err := utils.AesEncrypt(buf, utils.ToBytes(EnKey))
	if err != nil {
		return "", err
	}
	return utils.Base64Encode(enData), nil
}
