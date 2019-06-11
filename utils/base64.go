package utils

import "encoding/base64"


func Base64Decode(d string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(d)
}

func Base64Encode(d []byte) string {
	return base64.StdEncoding.EncodeToString(d)
}
