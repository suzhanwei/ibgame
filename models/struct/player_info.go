package strucct

type PlayerInfo struct {
	Pid            int64  `xorm:"not null pk autoincr BIGINT(20)"`
	Name           string `xorm:"not null default '' unique(uk_name) VARCHAR(128)"`
	NickName       string `xorm:"not null default '' unique(uk_name) VARCHAR(128)"`
	Position       int    `xorm:"not null unique(uk_type_position) INT(11)"`
	SecondPosition int    `xorm:"not null default 0 INT(11)"`
	Type           int    `xorm:"not null default 0 unique(uk_type_position) INT(11)"`
}
