package main

import (
	"actions"
	"logs"
	"net/http"
)

func main() {
	http.HandleFunc("/login", actions.Login)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		logs.Error.Println("ListenAndServe:error")
	}
}
