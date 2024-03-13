package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// 小写
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}

// 大写
func MD5Encode(data string) string {
	return strings.ToUpper(MD5Encode(data))
}

func MakePassword(plainPwd string, randomNum string) string {
	return Md5Encode(plainPwd + randomNum)
}

func ValidPassword(plainPwd string, randomNum string, password string) bool {
	return Md5Encode(plainPwd+randomNum) == password
}
