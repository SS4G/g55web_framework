package main

import (
	"flag"
	"fmt"
	"framework"
	"net/http"
)

func main() {
	var port = flag.Int("port", 12334, "server port")
	server := &http.Server{
		// 自定义的请求核心处理函数
		Handler: framework.NewCore(),
		// 请求监听地址
		Addr: fmt.Sprintf(":%d", *port),
	}
	server.ListenAndServe()
}
