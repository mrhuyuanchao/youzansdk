package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// Md5Encrypt MD5加密
func Md5Encrypt(content string, isUpper bool) string {
	h := md5.New()
	h.Write([]byte(content))
	sumStr := h.Sum(nil)
	result := hex.EncodeToString(sumStr)
	if isUpper {
		result = strings.ToUpper(result)
	}
	return result
}
