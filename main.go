package main

import (
	"learn/actions/user_action"
	"learn/logs"
	"net/http"
)

func main() {
	http.HandleFunc("/login", user_action.Login)
	http.HandleFunc("/register", user_action.Register)
	err := http.ListenAndServe(":12356", nil) //设置监听的端口
	if err != nil {
		logs.Error.Println("ListenAndServe:error")
	}
}
