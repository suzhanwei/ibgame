package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "https://restapi.getui.com/v1/g5heXUmtiv9UcdMBvEGJI1/push_single"
	payload := strings.NewReader("{  \n   \"message\":{  \n      \"appkey\":\"al0zZ6nvSO9tvxPPrTVHD9\",\n      \"is_offline\":true,\n      \"offline_expire_time\":1000000000,\n      \"msgtype\":\"notification\"\n   },\n   \"notification\":{  \n      \"style\":{  \n         \"type\":0,\n         \"text\":\"通知内容\",\n         \"title\":\"通知标题\"\n      },\n      \"transmission_type\":true,\n      \"transmission_content\":\"透传内容\"\n   },\n   \"cid\":\"b1cb2c9b335f6af4ba885aaf3ca21929\",\n   \"requestid\":\"1441111119900\"\n}")
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("authtoken", "d0fcec0e0943c89f2048cc97b61ba26336647262a0636098e271f32f2f3405d3")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "139f244b-cfce-f21c-3ac6-b510a192ceb4")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
