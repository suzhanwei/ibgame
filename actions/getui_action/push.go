package getui_action

import (
	"encoding/json"
	"net/http"
	//"html/template"
	"ibgame/logs"
	//"log"
	"ibgame/models/getui_model"
	//"github.com/spf13/cast"
)
const(
	APPKEY="al0zZ6nvSO9tvxPPrTVHD9"
	MASTERSECRET="UpL4wlvdSh7xWMdRyyHP21"
	NOMSGTYPE="notification"
)
//添加
func PushSingle(w http.ResponseWriter, r *http.Request) {

	var ps getui_model.Single
	//if r.Method == "GET" {
	//	t, _ := template.ParseFiles("views/no.tpl")
	//	log.Println(t.Execute(w, nil))
	//} else {
		//text:=cast.ToString(r.FormValue("text"))
		//title:=cast.ToString(r.FormValue("title"))
		//cid:=cast.ToString(r.FormValue("cid"))
		text:="12345"
		title:="1234567"
		cid:="123"
		ps.Message = getui_model.MessageItem{
			Appkey:APPKEY,
			IsOffline:true,
			Msgtype:NOMSGTYPE,
		}
		style:=getui_model.StyleItem{
			Type:0,
			Text:text,
			Title:title,
		}
		ps.Notification=getui_model.NotificationItem{
			Style:style,
			TransmissionContent:"",
			TransmissionType:false,
		}
		ps.Cid=cid
		ps.Requestid="123456"
		logs.Info.Println(ps)
		if re, e := getui_model.PushSingel(ps); e != nil {
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
//	}
}
