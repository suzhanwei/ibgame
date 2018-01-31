package models

import (
	"logs"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

//用户信息

func Register(param RegisterParam) (ret map[string]string, err error) {
	var engine *xorm.Engine
	users := make([]UserLogin, 1)
	var e error
	users[0].Name = param.Name
	users[0].Password = param.Password
	uid := randSeq(10)
	users[0].Uid = uid
	users[0].Ctime = time.Now().Unix()
	if engine, e = xorm.NewEngine("mysql", "szw:123456@tcp(10.231.31.240:3306)/go?charset=utf8"); e == nil {
		if _, err := engine.Insert(&users); err != nil {
			logs.Error.Println("insert error", err)
		}
	}
	ret = make(map[string]string)
	ret[param.Name] = uid
	return
}

var letters = []rune("1234567890")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
