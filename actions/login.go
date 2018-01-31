package actions

import (
	"encoding/json"
	"logs"
	"models"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func Login(w http.ResponseWriter, r *http.Request) {
	uname := r.FormValue("username")
	pword, err := strconv.ParseInt(r.FormValue("password"), 10, 64)
	if err != nil {
		logs.Error.Println("strconv.ParseInt:err", err)
	}
	if r, e := models.Get(uname, pword); e != nil {
		logs.Error.Println("models.Get:err", err)
	} else {
		logs.Info.Println(r)
		ret := map[string]interface{}{"code": 0, "msg": "ok", "data": r}
		bytes, err := json.MarshalIndent(ret, " ", "    ")
		if err != nil {
			logs.Error.Println("json:err", err)
		}
		w.Write(bytes)
	}
}
