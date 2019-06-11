package cerr

import (
	"fmt"
	"runtime"

	"go.uber.org/zap"

	"carp.cn/whale/zaplogger"
	"strings"
	"go.uber.org/zap/zapcore"
	"runtime/debug"
)

type IError interface {
	GetCode() int
	GetMsg() string
	error
}


type CodeError struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("%s [code=%d]", e.Msg, e.Code)
}

func (e *CodeError) GetCode() int {
	return e.Code
}

func (e *CodeError) GetMsg() string {
	return e.Msg
}

func NewError(err error) CodeError {
	return CodeError{
		Code: -1,
		Msg:  err.Error(),
	}
}

func NewErrorWithCode(code int, msg ...string) *CodeError {
	return &CodeError{
		Code: code,
		Msg:  strings.Join(msg, ","),
	}
}

func RaiseError(err error) {
	if err == nil {
		return
	}
	// runtime.Caller速度慢
	pc, _, lineno, ok := runtime.Caller(1)
	src := ""
	if ok {
		src = fmt.Sprintf("%s:%d", runtime.FuncForPC(pc).Name(), lineno)
	}
	zaplogger.Error("RaiseError from", zap.String("src:", src),
		zap.String(" Error:", err.Error()))
	panic(err)
}

func RaiseErrorCode(code int, msg ...string) {
	err := NewErrorWithCode(code, msg...)
	pc, _, lineno, ok := runtime.Caller(1)
	src := ""
	if ok {
		src = fmt.Sprintf("%s:%d", runtime.FuncForPC(pc).Name(), lineno)
	}
	zaplogger.Error("RaiseError from", zap.String("src:", src),
		zap.String(" Error:", err.Error()))
	panic(err)
}

func CheckError(err error, code int, fields ...zapcore.Field) {
	if err != nil {
		zaplogger.Error(fmt.Sprintf("error code:%d", code), append(fields, zap.String(" Error:", err.Error()))...)
		panic(NewErrorWithCode(code, err.Error()))
	}
}

func CheckIError(c IError, fields ...zapcore.Field) {
	if c != nil {
		zaplogger.Error(fmt.Sprintf("error code:%d", c.GetCode()), append(fields, zap.Reflect("Interface:", c))...)
		panic(c)
	}
}

func CheckErrDoNothing(err error, msg string, fields ...zapcore.Field) {
	if err != nil {
		zaplogger.Error(msg, append(fields, zap.String(" Error:", err.Error()))...)
	}
}


func ErrorWipe(e interface{}) *CodeError {
	res := &CodeError{}
	switch v := e.(type) {
	case string:
		res.Msg = v
	case CodeError:
		res.Code = v.Code
		res.Msg = v.Msg
	case *CodeError:
		res.Code = v.Code
		res.Msg = v.Msg
	case IError:
		res.Code = v.GetCode()
		res.Msg = v.GetMsg()
	case runtime.Error:
		zaplogger.Error(v.(error).Error(), zap.String("\nstack : ", string(debug.Stack())))
		res.Msg = "runtime error"
	case error:
		zaplogger.Error(v.(error).Error(), zap.String("\nstack : ", string(debug.Stack())))
		res.Msg = v.Error()
	default:
		res.Msg =fmt.Sprintf("%v", v)
	}


	zaplogger.Error("ApiResponse", zap.Reflect("error:", res))

	if res.Code > 10000 && res.Code < 1000 {
		res.Code = RESPONSE_FAIL
	}

	return res
}