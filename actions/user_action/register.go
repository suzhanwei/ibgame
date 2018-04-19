package actions

import (
	"encoding/json"
	"io/ioutil"
	"logs"
	"models/user_model"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logs.Error.Println("ioutil.ReadAll err", err)
		return
	}
	logs.Info.Println("ServeHTTP", "reqeust = ", string(data))
	var param models.RegisterParam
	err = json.Unmarshal(data, &param)

	if param.Name == "" {
		re := "注册名字为空"
		ret := map[string]interface{}{"code": 401, "msg": "error", "data": re}
		bytes, err := json.MarshalIndent(ret, " ", "    ")
		if err != nil {
			logs.Error.Println("json:err", err)
		}
		w.Write(bytes)
		return
	}
	if re, e := user_model.Register(param); e != nil {
		logs.Error.Println("models.Register:err", err)
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
