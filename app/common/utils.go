package common

import (
	"crypto/md5"
	"encoding/base64"
)

func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}
func Md5(encode string) (decode string) {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(encode))
	cipherStr := md5Ctx.Sum(nil)
	return string(base64Encode(cipherStr))
}
