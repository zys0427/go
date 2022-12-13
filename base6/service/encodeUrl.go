package service

import (
	"crypto/md5"
	"fmt"
)

type EncodeUrl struct {
}

func (e EncodeUrl) GetCode(code string, category string) string {
	switch category {
	case "has256":
		return e.has256(code)
	case "md5":
		return e.md5(code)
	}
	return ""
}

func (e EncodeUrl) has256(strKey string) string {
	return "has256_" + strKey
}

func (e EncodeUrl) md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5Str := fmt.Sprintf("%x", has)
	return md5Str
}
