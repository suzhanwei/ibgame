package models

import (
	"logs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

//用户信息

func Get(uname string, pword int64) (ret map[string]string, err error) {
	var engine *xorm.Engine
	users := make([]UserLogin, 0)
	var e error
	if engine, e = xorm.NewEngine("mysql", "szw:123456@tcp(10.231.31.240:3306)/go?charset=utf8"); e == nil {

		if err = engine.Where("name=? and password=?", uname, pword).Find(&users); err != nil {
			logs.Error.Println("db:err", err)
			return
		} else {
			ret = make(map[string]string)
			if len(users) > 0 {
				ret[users[0].Name] = "success"
			}
		}
	}
	return
}
