package utils

import "strings"

// TypeName = *JYTradeDB.jyUserSubscribeOp | FuncName = SelectAll => JyUserSubscribeOp.SelectAll
func UrlRule(typeName, funcName string) string {
	tn := typeName
	index := strings.LastIndexByte(tn, '.')
	t := tn[index+1: len(tn)]

	return strings.Title(t) + "." + strings.Title(funcName)
}

// url = JyUserSubscribeOp/SelectAll => jyUserSubscribeOp.SelectAll
func UrlFromate(uri string) string {
	u := strings.Split(uri, "/")
	if len(u) < 2 {
		return ""
	}
	for k ,v := range u {
		u[k] = strings.Title(v)
	}

	t := u[len(u)-2:len(u)]
	return strings.Join(t, ".")
}