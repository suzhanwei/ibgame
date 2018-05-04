package manage_actions

import (
	"encoding/json"
	"fmt"
	"html/template"
	"ibgame/logs"
	"ibgame/models/manage_model"
	"log"
	"net/http"

	"github.com/spf13/cast"
)

//添加
func AddPlayer(w http.ResponseWriter, r *http.Request) {
	var pi manage_model.AddParam
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("views/new.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		pi.Name = r.FormValue("name")
		pi.PlayerID = cast.ToInt64(r.FormValue("id"))
		pi.NickName = r.FormValue("nickname")
		pi.Position = cast.ToInt(r.FormValue("position"))
		pi.SecondPosition = cast.ToInt(r.FormValue("second"))
		pi.Type = cast.ToInt(r.FormValue("type"))
		if re, e := manage_model.Add(pi); e != nil {
			logs.Error.Println("models.Get:err", e)
		} else {
			logs.Info.Println(re)
			ret := map[string]interface{}{"code": 0, "msg": "ok", "data": re}
			bytes, err := json.MarshalIndent(ret, " ", "    ")
			if err != nil {
				logs.Error.Println("json:err", err)
			}
			w.Write(bytes)
		}
	}
}
