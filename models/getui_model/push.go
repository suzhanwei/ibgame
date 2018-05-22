package getui_model

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"ibgame/logs"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/cast"
)

type Single struct {
	Message      MessageItem      `json:"message"`
	Notification NotificationItem `json:"notification"`
	Cid          string           `json:"cid"`
	Requestid    string           `json:"requestid"`
}

type MessageItem struct {
	Appkey            string `json:"appkey"`
	IsOffline         bool   `json:"is_offline"`
	OfflineExpireTime int    `json:"offline_expire_time"`
	Msgtype           string `json:"msgtype"`
}
type NotificationItem struct {
	Style               StyleItem `json:"style"`
	TransmissionType    bool      `json:"transmission_type"`
	TransmissionContent string    `json:"transmission_content"`
}

type StyleItem struct {
	Type  int    `json:"type"`
	Text  string `json:"text"`
	Title string `json:"title"`
}

const (
	APPKEY       = "al0zZ6nvSO9tvxPPrTVHD9"
	APPID        = "g5heXUmtiv9UcdMBvEGJI1"
	MASTERSECRET = "UpL4wlvdSh7xWMdRyyHP21"
)

func PushSingel(param Single) (ret string, err error) {

	url := "https://restapi.getui.com/v1/g5heXUmtiv9UcdMBvEGJI1/push_single"

	payload := strings.NewReader("{  \n   \"message\":{  \n      \"appkey\":\"GVvUv4M8FZAF7u5a9H79m6\",\n      \"is_offline\":true,\n      \"offline_expire_time\":10000000,\n      \"msgtype\":\"notification\"\n   },\n   \"notification\":{  \n      \"style\":{  \n         \"type\":0,\n         \"text\":\"请填写通知内容\",\n         \"title\":\"请填写通知标题\"\n      },\n      \"transmission_type\":true,\n      \"transmission_content\":\"透传内容\"\n   },\n   \"cid\":\"b1cb2c9b335f6af4ba885aaf3ca21929\",\n   \"requestid\":\"121111\"\n}")
	req, e2 := http.NewRequest("POST", url, payload)
	if e2 != nil {
		logs.Error.Println("http req err", e2)
		err = e2
		return
	}
	now := time.Now().UnixNano() / 1000000
	str := APPKEY + cast.ToString(now) + MASTERSECRET
	logs.Info.Println(str)

	s := getSha256(str)

	var a Auth

	a.Appkey = APPKEY
	a.Timestamp = cast.ToString(now)
	a.Sign = s

	r, _ := getAuth(a)

	req.Header.Add("authtoken", r)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "7e79383a-0751-c672-0a18-4f6d77b8c5b9")

	resp, e := http.DefaultClient.Do(req)
	if e != nil {
		logs.Error.Println("http resp", e)
		err = e
		return
	}
	b, e1 := ioutil.ReadAll(resp.Body)
	if e1 != nil {
		logs.Error.Println("http resp", e1)
		err = e
		return
	}
	defer resp.Body.Close()
	ret = string(b)
	return
}

// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

type Auth struct {
	Sign      string `json:"sign"`
	Timestamp string `json:"timestamp"`
	Appkey    string `json:"appkey"`
}

func getAuth(p Auth) (ret string, err error) {

	payloadBytes, e := json.Marshal(p)

	if e != nil {
		logs.Error.Println("http resp", e)
		err = e
		return
	}
	body := bytes.NewReader(payloadBytes)

	req, e1 := http.NewRequest("POST", "https://restapi.getui.com/v1/"+APPID+"/auth_sign", body)
	if e1 != nil {
		logs.Error.Println("http resp", e1)
		err = e1
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, e2 := http.DefaultClient.Do(req)
	if e2 != nil {
		logs.Error.Println("http resp", e1)
		err = e2
		return
	}

	b, e3 := ioutil.ReadAll(resp.Body)
	if e3 != nil {
		logs.Error.Println("http resp", e3)
		err = e3
		return
	}
	var re Result
	_ = json.Unmarshal(b, &re)
	fmt.Println(string(b))
	ret = re.AuthToken
	defer resp.Body.Close()
	return
}

type Result struct {
	Result    string `json:"result"`
	AuthToken string `json:"auth_token"`
}

func getSha256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}
