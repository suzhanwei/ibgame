package manage_actions

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("views/index.tpl")
		log.Println(t.Execute(w, nil))
	}
}
