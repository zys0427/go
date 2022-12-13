package main

import (
	"fmt"
	db2 "go.com/httpbase/base6/db"
	"go.com/httpbase/base6/service"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/addUrl", addUrl)
	http.HandleFunc("/jumpUrl", jumpUrl)
	err := http.ListenAndServe("localhost:9999", nil)
	if err != nil {
		log.Fatal(err)
	}

}

// 添加Url
func addUrl(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	defer r.Body.Close()

	// POST 获取参数
	url := r.PostFormValue("url")
	//log.Println(url)

	// URL 校验
	urlCheck := service.UrlCheck{}
	if ok := urlCheck.CheckUrl(url); !ok {
		fmt.Println("Ulr 不符合规则")
		return
	}
	
	// URL 转换
	encodeUrl := service.EncodeUrl{}
	urlMd5 := encodeUrl.GetCode(url, "md5")

	// URL 保存
	saveUrl(urlMd5, url)
	log.Println(urlMd5, url)
	return
}

func saveUrl(code string, url string) error{
	// redis 保存
	redisZys := service.RedisZys{}
	client := redisZys.RedisConnZys()
	client.Set(code, url, 0)

	// Mysql 保存
	db2 := db2.MysqlZys{}
	db2.Init()
	err := db2.ExecZys(code, url)
	return err
}

// 跳转Url
func jumpUrl(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data := r.URL.Query()
	urlKey := data.Get("url")
	url := getUrl(urlKey)
	w.Header().Set("Location", url)
	w.WriteHeader(302)
	return
}

func getUrl(code string) string {
	fmt.Println(code)
	redisZys := service.RedisZys{}
	client := redisZys.RedisConnZys()
	url, _ := client.Get(code).Result()
	if url == ""  {
		mysqlZys := db2.MysqlZys{}
		mysqlZys.Init()
		url = mysqlZys.GetOneUrlByCode(code)
		encodeUrl := service.EncodeUrl{}
		code := encodeUrl.GetCode(url, "md5")
		client.Set(code, url, 0)
	}
	return url
}



