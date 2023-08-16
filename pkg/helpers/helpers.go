package helpers

import (
	"crypto/md5"
	"encoding/hex"
)

// Str2Md5 将 string 做 md5 处理
func Str2Md5(str string) string {
	data := md5.Sum([]byte(str))
	hashString := hex.EncodeToString(data[:])
	return hashString
}
