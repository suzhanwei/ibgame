package user_model

import (
	"ibgame/logs"

	"ibgame/models"
)

//用户信息

func Get(uname string, pword int64) (ret map[string]string, err error) {

	users := make([]UserLogin, 0)
	if engine, e := models.GetEngine(); e == nil {
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
