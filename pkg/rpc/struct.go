package rpc

import (
	"reflect"
	"carp.cn/whale/pkg/cerr"
)

type Request struct {
	Method string        `json:"func_name"`
	Params []interface{} `json:"params"`
}

type Response struct {
	Method string      `json:"func_name"`
	Result interface{} `json:"data"`
	Error  cerr.IError `json:"error"`
}

type Method struct {
	Method reflect.Method
	host   reflect.Value
	idx    int
}
