package utils

import "regexp"

const (
	MobileNumber = "^((13[0-9])|(14[5|7])|(15([0-3]|[5-9]))|(17[0-9])|(18[0,5-9]))\\d{8}$"
)


func IsMoblieNumber(phone string) bool {
	reg := regexp.MustCompile(MobileNumber)
	return reg.MatchString(phone)
}
