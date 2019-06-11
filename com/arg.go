package com

import (
	"carp.cn/whale/utils"
	"strings"
	"math"
	"strconv"
	"fmt"
	"time"
)

type BaseArg struct {
	OptAct     string `json:"opact" form:"opact" binding:"required"`             // 接口名称
	EnData     string `json:"en_data" form:"en_data" binding:"required"`         // 加密后的数据
	EnKey      string `json:"en_key" form:"en_key" binding:"required"`           // 前端秘钥  => 将会被转成完整密钥
	Version    string `json:"version" form:"version" binding:"required"`         // 版本
	DeviceType string `json:"device_type" form:"device_type" binding:"required"` // 驱动类型
	Debug      string `json:"debug" form:"debug"`                                // 模式
}

func (b *BaseArg) GetData() ([]byte, error) {
	data, err := utils.Base64Decode(b.EnData)
	if err != nil {
		return nil, err
	}
	version := getVersion(b.Version)
	if version == "" {
		return nil, fmt.Errorf("version no find:%s", b.Version)
	}
	vEncrypt, find := Encrypt_KEY[version]
	if !find {
		return nil, fmt.Errorf("version ecrypt no find:%s", b.Version)
	}
	deKey, find := vEncrypt[b.EnKey]
	if !find {
		return nil, fmt.Errorf("ecrypt no find:%s", b.EnKey)
	}

	// 从新保存密钥
	b.EnKey = deKey

	return utils.AesDecrypt(data, []byte(deKey))
}

//-------------------------------------------------------------------------------------------------------
type SysArg struct {
	AppModel     string `json:"app_model"`     // 手机型号
	Platform     string `json:"platform"`      // 平台编号
	XGPushDevice string `json:"xgpush_device"` // 信鸽的设备标识
	JPushDevice  string `json:"jpush_device"`  // 极光的设备标识
	Tms          string `json:"tms"`           // 时间字符串 例：20141212101010
	Source       string `json:"source"`        // 渠道来源
	Imei         string `json:"imei"`          // 手机设备唯一标识
}

func (s *SysArg) GetTime() (*time.Time, error) {
	t, err := time.Parse(s.Tms, TIME_LAYOUT)
	return &t, err
}

//-------------------------------------------------------------------------------------------------------

type TOClient struct {
	EnData string `json:"en_data"`
}

//-------------------------------------------------------------------------------------------------------

func getVersion(ve string) string {

	cVersion := strings.Split(ve, "_")
	if len(cVersion) != 3 {
		return ""
	}

	appVersion := map[string][]string{}
	for key, val := range APP_VERSIONS {
		sV := strings.Split(key, "_")
		appVersion[val] = sV
	}

	vs := matchVersion(0, cVersion, appVersion)

	if len(vs) > 1 {
		vInt := map[float64]string{}
		flag := float64(0)
		for k, v := range vs {
			codeSum := float64(0)
			for i, code := range v {
				if code != "*" {

					if num, err := strconv.Atoi(code); err == nil {
						codeSum += math.Pow10(10-i) * float64(num)
					}
				}
			}
			vInt[codeSum] = k
			if flag < codeSum {
				flag = codeSum
			}
		}
		return vInt[flag]
	}

	for k := range vs {
		return k
	}
	return ""
}

func matchVersion(index int, cVersion []string, version map[string][]string) map[string][]string {
	if index > len(version) || index > len(cVersion) || index < 0 {
		return version
	}
	for k, v := range version {
		if !(v[index] == "*" || v[index] == cVersion[index]) {
			delete(version, k)
		}
	}
	return matchVersion(index+1, cVersion, version)
}
