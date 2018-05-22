package manage_actions

import (
	"encoding/json"
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

	if r.Method == "GET" {
		t, _ := template.ParseFiles("views/new.tpl")
		log.Println(t.Execute(w, nil))
	} else {
		pi.Name = r.FormValue("name")
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
