package user_model

import (
	"ibgame/logs"
	"ibgame/models/mysql"
	"ibgame/utils"
	"math/rand"
	"time"

	"github.com/go-xorm/xorm"
)

const (
	ReKey = "jiamisuzhanweiqw"
)

//用户信息
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Register 注册
func Register(param RegisterParam) (ret map[string]string, err error) {
	var engine *xorm.Engine
	ret = make(map[string]string)
	var user UserLogin
	if engine, err = mysql.GetEngine(); err != nil {
		logs.Error.Println("engine error", err)
	}

	var str string
	if str, err = utils.Aes128CBCDecrypt(param.Token, ReKey, true); err != nil {
		logs.Error.Println(" error", err)
	}

	now := time.Now().Unix()
	user.Name = param.Name
	user.Password = str
	uid := randSeq(10)
	user.Uid = uid
	user.Cid = param.Cid
	user.Ctime = now
	user.Mail = param.Mail
	user.MailVerify = 0
	user.Utime = now
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
