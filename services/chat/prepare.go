package chat

import "fmt"

func renderErrorMessage(code int) []byte {
	return []byte(fmt.Sprintf(`{"method":"", "data":null, "error":{"code":%d}}`, code))
}

func CloseMessage(code int) []byte {
	return []byte(fmt.Sprintf(`{"func_name":"Kickout", "data":{"kick_out": %d}, "error":{}}`, code))
}
