package util

import (
	"crypto/md5"
	"encoding/hex"
)

// EncodeMD5 计算文件名的MD5
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	// hex进行编码
	return hex.EncodeToString(m.Sum(nil))
}
