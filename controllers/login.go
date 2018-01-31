package controllers

import (
	"encoding/json"
	"fmt"
	"models"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func Login(w http.ResponseWriter, r *http.Request) {
	uname := r.FormValue("username")
	pword, err := strconv.ParseInt(r.FormValue("password"), 10, 64)
	if err != nil {
		panic(err)
	}
	if r, e := models.Get(uname, pword); e != nil {
		fmt.Println(e)
	} else {
		ret := map[string]interface{}{"code": 0, "msg": "ok", "data": r}
		bytes, err := json.MarshalIndent(ret, " ", "    ")
		if err != nil {
			fmt.Println("json", err)
		}
		w.Write(bytes)
	}
}
