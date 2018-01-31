package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

//用户信息

func Get(uname string, pword int64) (ret map[string]string, err error) {
	var engine *xorm.Engine
	users := make([]userinfo, 0)
	var e error
	engine, e = xorm.NewEngine("mysql", "szw:123456@tcp(10.231.31.240:3306)/test?charset=utf8")
	if e == nil {
		if err = engine.Where("username=? and password=?", uname, pword).Find(&users); err != nil {
			panic(err)
		} else {
			ret = make(map[string]string)
			if len(users) > 0 {
				ret[users[0].Username] = "success"
			}
		}
	}
	return
}
