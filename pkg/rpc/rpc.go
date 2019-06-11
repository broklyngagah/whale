package rpc

import (
	"fmt"
	"reflect"
	"runtime"
	"runtime/debug"

	"unicode"
	"unicode/utf8"

	"go.uber.org/zap"
	"carp.cn/whale/zaplogger"
	"carp.cn/whale/pkg/cerr"
)

// 方法名注册规则
type Rule func(typeName, funcName string) string

//----------------------------------------------------------------------------------------

type RpcHelper struct {
	Methods map[string]*Method
}

func NewRpcHelper() *RpcHelper {
	return &RpcHelper{make(map[string]*Method)}
}

//处理客户端消息
// baseParam 一般为基础参数如client 等
func (h *RpcHelper) Handler(req *Request, baseParam ...interface{}) *Response {
	resp := &Response{}
	method, params, err := h.Parse(req, baseParam...)
	if err != nil {
		resp.Error = err
		return resp
	}

	return h.Call(method, params)
}

// 解析客户端请求，
// default_params : 服务器端调用时自动带入的参数, 和客户端请求的参数共同组成method的参数。
func (h *RpcHelper) Parse(req *Request, defaultParams ...interface{}) (*Method, []reflect.Value, *cerr.CodeError) {
	method, ok := h.Methods[req.Method]
	if !ok {
		return nil, nil, &cerr.CodeError{Code: ERR_METHOD_NOT_FOUND, Msg: fmt.Sprintf("[RPC]: method not found.[method=%s]", req.Method)}
	}

	defaultParamsLen := len(defaultParams)
	//长度应减去method的receiver
	var lens int
	if req.Params == nil {
		lens = 0
	} else {
		lens = len(req.Params)
	}

	if lens != (method.Method.Type.NumIn() - defaultParamsLen - 1) {
		return nil, nil, &cerr.CodeError{Code: ERR_PARAM_COUNT_NOT_MATCH,
			Msg: fmt.Sprintf("[RPC]: params not matched. got %d, need %d.", lens, method.Method.Type.NumIn()-defaultParamsLen-1)}
	}

	params := make([]reflect.Value, lens+defaultParamsLen)
	//第一个参数是*chat.Client
	for idx, hdnParam := range defaultParams {
		params[idx] = reflect.ValueOf(hdnParam)
	}

	for i := 0; i < lens; i++ {
		targetType := method.Method.Type.In(i + 1 + defaultParamsLen) //跳过receiver和default_params
		newParam, ok := convertParam(req.Params[i], targetType)
		if !ok {
			errMsg := fmt.Sprintf("[RPC]: convert param faild. expect %s, found=%v value=%v.",
				targetType, reflect.TypeOf(req.Params[i]), req.Params[i])
			return nil, nil, &cerr.CodeError{Code: ERR_PARAM_INVALID, Msg: errMsg}
		}
		params[i+defaultParamsLen] = newParam
	}

	return method, params, nil
}

func (h *RpcHelper) Call(method *Method, params []reflect.Value) (resp *Response) {

	resp = &Response{}

	resp.Method = method.Method.Name

	defer func() {
		if re := recover(); re != nil {
			switch v := re.(type) {
			case cerr.CodeError:
				resp.Error = &v
			case *cerr.CodeError:
				resp.Error = v
			case cerr.IError:
				resp.Error = re.(*cerr.CodeError)
			case string:
				zaplogger.Error(v, zap.String("\nstack : ", string(debug.Stack())))
				resp.Error = &cerr.CodeError{Code: ERR_RUNTIME}
			case runtime.Error:
				zaplogger.Error(v.Error(), zap.String("\nstack : ", string(debug.Stack())))
				resp.Error = &cerr.CodeError{Code: ERR_RUNTIME}
			case error:
				zaplogger.Error(v.Error(), zap.String("\nstack : ", string(debug.Stack())))
				resp.Error = &cerr.CodeError{Code: ERR_RUNTIME}
			default:
				debug.PrintStack()
				resp.Error = &cerr.CodeError{Code: ERR_RUNTIME}
			}
		}
	}()

	result := method.host.Method(method.idx).Call(params)
	if len(result) > 0 {
		resp.Result = result[0].Interface()
	}
	return
}

func (h *RpcHelper) RegisterMethod(v interface{}) {
	reflectType := reflect.TypeOf(v)
	host := reflect.ValueOf(v)
	for i := 0; i < reflectType.NumMethod(); i++ {
		m := reflectType.Method(i)
		char, _ := utf8.DecodeRuneInString(m.Name)
		if !unicode.IsUpper(char) {
			continue
		}
		h.Methods[m.Name] = &Method{m, host, m.Index}
	}
}

// 将对象的方法按照一定的变换规则进行映射
func (h *RpcHelper) RegisterMethodByRule(v interface{}, rule Rule) {
	reflectType := reflect.TypeOf(v)
	host := reflect.ValueOf(v)
	for i := 0; i < reflectType.NumMethod(); i++ {
		m := reflectType.Method(i)
		char, _ := utf8.DecodeRuneInString(m.Name)
		//非导出函数不注册
		if !unicode.IsUpper(char) {
			continue
		}
		h.Methods[rule(reflectType.String(), m.Name)] = &Method{m, host, m.Index}
	}
}
//----------------------------------------------------------------------------------------
// JSON standard : all number are Number type, that is float64 in golang.
func convertParam(v interface{}, targetType reflect.Type) (newV reflect.Value, ok bool) {
	defer func() {
		if re := recover(); re != nil {
			ok = false
			zaplogger.Error("[RPC]: convertParam", zap.Reflect("recover:", re))
		}
	}()

	ok = true

	if targetType.Kind() == reflect.Interface {
		newV = reflect.ValueOf(v)
	} else if reflect.TypeOf(v).Kind() == reflect.Float64 {
		f := v.(float64)
		switch targetType.Kind() {
		case reflect.Int:
			newV = reflect.ValueOf(int(f))
		case reflect.Uint8:
			newV = reflect.ValueOf(uint8(f))
		case reflect.Uint16:
			newV = reflect.ValueOf(uint16(f))
		case reflect.Uint32:
			newV = reflect.ValueOf(uint32(f))
		case reflect.Uint64:
			newV = reflect.ValueOf(uint64(f))
		case reflect.Int8:
			newV = reflect.ValueOf(int8(f))
		case reflect.Int16:
			newV = reflect.ValueOf(int16(f))
		case reflect.Int32:
			newV = reflect.ValueOf(int32(f))
		case reflect.Int64:
			newV = reflect.ValueOf(int64(f))
		case reflect.Float32:
			newV = reflect.ValueOf(float32(f))
		default:
			ok = false
		}
	} else if reflect.TypeOf(v).Kind() == targetType.Kind() {
		newV = reflect.ValueOf(v)
	} else if targetType.Kind() == reflect.Ptr { //if it is pointer, get it element type
		newV = reflect.ValueOf(&v) //targetType.Elem()
	} else {
		ok = false
	}

	return
}