package user_model

import (
	"ibgame/logs"

	"ibgame/models"

	"github.com/go-xorm/xorm"
)

//用户信息

func Get(uname string, pword int64) (ret map[string]string, err error) {

	var engine *xorm.Engine
	users := make([]UserLogin, 0)
	var e error
	if engine, e = models.GetEngine(); e == nil {
		if err = engine.Where("name=? and password=?", uname, pword).Find(&users); err != nil {
			logs.Error.Println("db:err", err)
			return
		} else {
			ret = make(map[string]string)
			if len(users) > 0 {
				ret[users[0].Name] = "success"
			}
		}
	} else {
		logs.Error.Println("db :err ", e)
	}
	return
}
