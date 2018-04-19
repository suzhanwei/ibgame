package models

import (
	"logs"
	"math/rand"
	"models"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

//用户信息

func Register(param models.RegisterParam) (ret map[string]string, err error) {
	var engine *xorm.Engine
	ret = make(map[string]string)
	var user models.UserLogin
	if engine, err = xorm.NewEngine("mysql", "work:123456@tcp(39.107.94.42:3306)/go?charset=utf8"); err != nil {
		logs.Error.Println("engine error", err)
	}
	user.Name = param.Name
	user.Password = param.Password
	uid := randSeq(10)
	user.Uid = uid
	user.Ctime = time.Now().Unix()

	if _, err := engine.Insert(&user); err != nil {
		logs.Error.Println("insert error", err)
	}
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
