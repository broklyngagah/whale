package cerr

const (
	ERR_JSON_UNMARSHAL        = 10100 + iota // 读取json数据错误
	ERR_JSON_MARSHAL                         // json数据错误
	ERR_CALL_DENIED                          // 没有登录
	ERR_OTHER                                // 其他错误
	ERR_REQUEST                              // 请求失败
	ERR_PARAM_TYPE                           // 参数类型错误
	ERR_PARAM_VALUE                          // 参数值错误
	ERR_BINDING                              // 参数绑定错误
	ERR_BASE64_DECODE                        // Base64 解码错误
	ERR_READ_HTTP_BODY                       // 解析http body错误
	ERR_OBJECT_NULL                          // 空对象
	ERR_WRITE_TO_CLIENT                      // 写入错误
	ERR_DECRYPT                              // 解密错误
	ERR_ENCRYPT                              // 加密错误
	ERR_PHONE_NUMBER                         // 电话号码错误
	ERR_SMS_CODE                             // sms code 错误
)

//后台服务错误
const (
	ERR_DB_QUERY    = 20000 + iota //数据库查询错误
	ERR_DB_INSERT                  //数据库写入错误
	ERR_DB_DELETE                  //数据库删除错误
	ERR_DB_UPDATE                  //数据库写入错误
	ERR_TX_BEGIN                   //事务开始失败
	ERR_TX_COMMIT                  //事务提交失败
	ERR_TX_ROLLBACK                //事务回滚失败
)

//
const (
	RESPONSE_SUCCESSS  = 1000 // 请求成功
	RESPONSE_FAIL      = 1001 // 请求失败	toast提示
	RESPONSE_ARTICLE   = 1100 // 文章不存在或被删除	跳转文章已被删除页
	RESPONSE_NOT_LOGIN = 1099 // 未登录	跳转到登录页
	RESPONSE_RElOGIN   = 1098 // 单点登录	弹出另一台设备登录的提示，并退出登录
)
