package user_model

type UserLogin struct {
	Id       int    `xorm:"not null pk autoincr INT(10)"`
	Name     string `xorm:"not null default '' index VARCHAR(100)"`
	Uid      string `xorm:"not null default '' VARCHAR(40)"`
	Password string `xorm:"not null default '' VARCHAR(40)"`
	Ctime    int64  `xorm:"not null default 0 BIGINT(20)"`
	Cid      string `xorm:"not null default '' VARCHAR(40)"`
}

type RegisterParam struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Cid      string `json:"cid"`
}
