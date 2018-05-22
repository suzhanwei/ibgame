package manage_model

import (
	"ibgame/logs"
	"ibgame/models"
)

func Add(p AddParam) (ret string, err error) {

	var pi PlayerInfo

	pi.Name = p.Name
	pi.NickName = p.NickName
	pi.Position = p.Position
	pi.SecondPosition = p.SecondPosition
	pi.Type = p.Type
	if engine, e := models.GetEngine(); e == nil {
		if _, e1 := engine.InsertOne(&pi); e != nil {
			logs.Error.Println("db :err ", e1)
			return "false", e1
		}
	} else {
		logs.Error.Println("db :err ", e)
		return "false", e
	}
	ret = "success"
	return
}
