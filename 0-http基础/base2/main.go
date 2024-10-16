package main

import (
	"fmt"
	"log"
	"net/http"
)

// Engine 定义一个空结构体
type Engine struct{}

/**
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

Handler 接口只有一个方法：ServeHTTP
*/

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	}
}

func main() {
	log.Fatal(http.ListenAndServe(":9999", new(Engine)))
}
