/***
hello word
 */
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatalln(http.ListenAndServe(":9999", nil))
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.PATH=%q \n", req.URL.Path)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
