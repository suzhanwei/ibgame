package user_model

import (
	"learn/logs"
	"learn/models"
	"math/rand"
	"time"

	"github.com/go-xorm/xorm"
)

//用户信息
func init() {
	rand.Seed(time.Now().UnixNano())
}

func Register(param RegisterParam) (ret map[string]string, err error) {
	var engine *xorm.Engine
	ret = make(map[string]string)
	var user UserLogin
	if engine, err = models.GetEngine(); err != nil {
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
