package bigv

import (
	"carp.cn/whale/utils"
	"fmt"
)

type BigV struct {
}

func (v *BigV) Login(tel string, sms string) (string, error) {
	if utils.IsMoblieNumber(tel) {
		return "", fmt.Errorf("tel code error: %s", tel)
	}

	if !authAccessSms(sms) {
		return "", fmt.Errorf("sms error: %s", sms)
	}
	return "", nil

}

func authAccessSms(sms string) bool {
	return true
}
