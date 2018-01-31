package main

import (
	"controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", controllers.Login)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
