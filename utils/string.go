package utils

import (
	"unsafe"
	"fmt"
	"strconv"
)

func ToString(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

func ToBytes(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(&str))
}

func ConvertString(i interface{}) string {
	switch t := i.(type) {
	case string:
		return t
	case int:
		return strconv.Itoa(int(t))
	case int8:
		return strconv.Itoa(int(t))
	case int16:
		return strconv.Itoa(int(t))
	case int32:
		return strconv.Itoa(int(t))
	case int64:
		return strconv.Itoa(int(t))
	case uint:
		return strconv.Itoa(int(t))
	case uint8:
		return strconv.Itoa(int(t))
	case uint16:
		return strconv.Itoa(int(t))
	case uint32:
		return strconv.Itoa(int(t))
	case uint64:
		return strconv.Itoa(int(t))
	case float32:
		return fmt.Sprintf("%.3f", t)
	case float64:
		return fmt.Sprintf("%.3f", t)
	default:
		return fmt.Sprintf("%v", t)
	}
}
