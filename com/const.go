package com

var APP_VERSIONS = map[string]string{

	"1_2_*": "v1_2",
	"1_*_*": "v1",
}

var API_SIGN_KEY = map[string]string{
	"v1": "sjf&ds#d$FfdWr+rw",
}

var Encrypt_KEY = map[string]map[string]string{
	"v1": map[string]string{
		"ij#n+yus": "iojyxgas+x*$a$*s",
		"anj#*yud": "a*jyxga#+wdfa%nd",
	},
}

var PLATFORM_MAP = map[string]int{
	"jy*&#ios*&":  1,
	"jyand$ro*id": 2,
}

var MAPI_SIGN_ = map[string]string{
	"v1": "sjf&ds#d$FfdWr+rw",
}

const (
	IOS_SUBMIT = 1
)

const (
	TIME_LAYOUT = "20060102150405"
	RESPONSE_TIME_LAYOUT = "2016-01-02 15:04:05"
)

const (
	CHAT_TYPE_PONG       = "pong"
	CHAT_TYPE_LOGIN      = "login"
	CHAT_TYPE_SEND       = "say"
	CHAT_TYPE_CLOSE_ROOM = "closeRoom"
	CHAT_TYPE_SUBSCRIBE  = "subscribe"
	CHAT_TYPE_FOLLOW     = "follow"
	CHAT_TYPE_HISTORY    = "history"
	CHAT_TYPE_LOGIN_CHECK = "loginCheck"
)

const (
	VISITOR = "游客"
	SYSTEM = "系统"
)
