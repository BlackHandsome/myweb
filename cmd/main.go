package main

import (
	"log"
	"net/http"

	"awesome.project/webtest/service"
)

func main() {
	s := service.New()
	http.HandleFunc("/", s.SayhelloName)       // 设置访问的路由
	http.HandleFunc("/login", s.Login)         // 设置访问的路由
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
