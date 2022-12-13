package service

import "fmt"

type UrlCheck struct {
}

func (uk UrlCheck) CheckUrl(url string) bool {
	return uk.isEmpty(url) && uk.isHttp(url)
}

// 是否为空
func (uk UrlCheck) isEmpty(url string) bool {
	fmt.Println(url)
	if url == "" {
		return false
	}
	return true
}

// 校验是否符合Http
func (uk UrlCheck) isHttp(url string) bool {

	return true
}