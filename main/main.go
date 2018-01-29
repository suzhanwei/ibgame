package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

//用户信息
type userinfo struct {
	Uid      int    `json:"uid"`
	Username string `json:"username"`
	Created  int64  `json:"create"`
	Password int64  `json:"password"`
}

//sayhello
func Sayhello(w http.ResponseWriter, r *http.Request) {
	var engine *xorm.Engine
	users := make([]userinfo, 0)
	var e error
	engine, e = xorm.NewEngine("mysql", "szw:123456@tcp(10.231.31.240:3306)/test?charset=utf8")
	uname := r.FormValue("username")
	pword, err := strconv.ParseInt(r.FormValue("password"), 10, 64)
	if err != nil {
		panic(err)
	}
	if e == nil {
		if err := engine.Where("username=? and password=?", uname, pword).Find(&users); err != nil {
			panic(err)
		} else {
			ret := make(map[string]int64)
			if len(users) > 0 {
				ret[users[0].Username] = users[0].Password
				bytes, err := json.MarshalIndent(ret, " ", "    ")
				if err == nil {
					w.Write(bytes)
				}
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
