package main

import (
	"actions/user_action"
	"logs"
	"net/http"
)

func main() {
	http.HandleFunc("/login", actions.Login)
	http.HandleFunc("/register", actions.Register)
	err := http.ListenAndServe(":12356", nil) //设置监听的端口
	if err != nil {
		logs.Error.Println("ListenAndServe:error")
	}
}
