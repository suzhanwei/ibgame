package main

import (
	
	"fmt"
	"ibgame/actions/user_action"
	"ibgame/actions/manage_actions"
	"ibgame/logs"
	"net/http"

)

func main() {
	http.HandleFunc("/add", manage_actions.AddPlayer)
	http.HandleFunc("/index",manage_actions.Index)
	http.HandleFunc("/login", user_action.Login)
	http.HandleFunc("/register", user_action.Register)
	err := http.ListenAndServe(":12356", nil) //设置监听的端口
	if err != nil {
		logs.Error.Println("ListenAndServe:error")
	}else{
		fmt.Println("监听12356")
	}
}
