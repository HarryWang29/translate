package util

import (
	"encoding/base64"
	"log"
	"strings"
)

func Base64Decode(s string) string {
	//将base64的等号去掉，解码使用不自动补齐方式
	s = strings.ReplaceAll(s, "=", "")
	encrypted, err := base64.RawStdEncoding.DecodeString(s)
	if err != nil {
		log.Printf("%s", err)
		return ""
	}
	return string(encrypted)
}
