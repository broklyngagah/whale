package game_error

import (
	"errors"
	"fmt"
	"runtime"

	"go.uber.org/zap"

	"changit.cn/contra/server/zaplogger"
)

type GameError struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

func (ge *GameError) Error() string {
	return fmt.Sprintf("%s [code=%d]", ge.Msg, ge.Code)
}

func NewError(err error) GameError {
	return GameError{
		Code: -1,
		Msg:  err.Error(),
	}
}

func NewGameError(code int, msg string) *GameError {
	return &GameError{code, msg}
}

func NewErrorWithCode(code int, msg ...string) GameError {
	var m string
	if len(msg) >= 1 {
		m = msg[0]
	} else {
		m = ""
	}
	return GameError{
		Code: code,
		Msg:  m,
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

func WrapError(err interface{}) error {
	switch err.(type) {
	case error:
		return err.(error)
	case string:
		return errors.New(err.(string))
	default:
		return errors.New(fmt.Sprintf("unknown :%v", err))
	}
}
