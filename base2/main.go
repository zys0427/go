/**
启动 http 监控服务  端口 9999
 */
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	engine := new(Engine)

	log.Fatal(http.ListenAndServe(":9999", engine))
}

type Engine struct {
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path=%q", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[q%]= q%\n", k, v)
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 NOT FOUND: %s\n",req.URL)
	}

}
