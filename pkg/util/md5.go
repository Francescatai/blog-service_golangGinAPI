package util

import (
	"crypto/md5"
	"encoding/hex"
)

// 將文件名MD5後再進行寫入
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}