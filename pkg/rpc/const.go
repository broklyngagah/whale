package rpc

const (
	ERR_JSON_UNMARSHAL        = 2000 + iota //读取json数据错误
	ERR_PARSE_RPC_REQ                       //解析RPC请求错误
	ERR_METHOD_NOT_FOUND                    //没有找到方法
	ERR_PARAM_COUNT_NOT_MATCH               //调用参数数量不匹配
	ERR_PARAM_INVALID                       // 参数类型错误
	ERR_JSON_MARSHAL                        //json数据错误
	ERR_RUNTIME                             //运行时错误

)
