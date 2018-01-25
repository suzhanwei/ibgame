package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Userinfo struct {
	Uid        int    `json:"uid"`
	Username   string `json:"username"`
	Departname string `json:"departname"`
	Created    int64  `json:"create"`
	Password   int64  `json:"password"`
}

func Sayhello(w http.ResponseWriter, r *http.Request) {
	var engine *xorm.Engine
	users := make([]Userinfo, 0)
	var e error
	engine, e = xorm.NewEngine("mysql", "root:123456@/test?charset=utf8")
	uname := r.FormValue("username")
	pword, err := strconv.ParseInt(r.FormValue("password"), 10, 64)
	if err != nil {
		panic(err)
	}
	if e == nil {
		if err := engine.Where("username=? and password=?", uname, pword).Find(&users); err != nil {
			panic(err)
		} else {
			if len(users) > 0 {
				fmt.Fprintln(w, "success")
			}
		}
	}
}
func main() {
	http.HandleFunc("/", Sayhello)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
