package main

import (
	"fmt"
	"log"
	"net/http"
)

/**
net/http 包下只需两个方法
- HandleFunc：注册路由以及对应的处理函数
- ListenAndServe：监听并启动服务
  - 第一个参数：IP+port，IP可省略默认为本地地址
  - 第二个参数：Handler接口，是处理所有的HTTP请求的实例，nil 代表使用标准库中的实例处理
    实现Handler接口，就可以基于net/http标准库实现Web框架的入口。
*/

/**
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

Handler 接口只有一个方法：ServeHTTP
*/

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

// handler echoes r.URL.Path
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

// handler echoes r.URL.Header
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
