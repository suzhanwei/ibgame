package manage_model

import "github.com/astaxie/beego/orm"

func init() {
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//设置默认数据库
	orm.RegisterDataBase("default", "mysql", "work:123456@tcp(39.107.94.42:3306)/go?charset=utf8", 30)
	//注册定义的model
	orm.RegisterModel(new(PlayerInfo))
	// 创建table
	orm.RunSyncdb("default", false, true)
}
