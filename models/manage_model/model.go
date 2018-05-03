package manage_model

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type PlayerInfo struct {
	ID             int64 `pk`
	PlayerID       int64
	Name           string
	NickName       string
	Position       int
	SecondPosition int
	Type           int
}

func (pi PlayerInfo) Add() (ret string) {

	o := orm.NewOrm()

	pi.Name = "Donovan Mitchell"
	pi.NickName = "米切尔"
	pi.PlayerID = 45
	pi.Position = 2
	pi.SecondPosition = 1
	pi.Type = 1

	_, err := o.Insert(&pi)
	if err == nil {
		ret = "success"
	}
	return
}
