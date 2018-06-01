package manage_model

//player base info
type PlayerInfo struct {
	Pid            int64  `xorm:"BIGINT(20)"`
	Name           string `xorm:"not null default '' unique(uk_name) VARCHAR(128)"`
	NickName       string `xorm:"not null default '' unique(uk_name) VARCHAR(128)"`
	Position       int    `xorm:"not null unique(uk_type_position) INT(11)"`
	SecondPosition int    `xorm:"not null default 0 INT(11)"`
	Type           int    `xorm:"not null default 0 unique(uk_type_position) INT(11)"`
}

type AddParam struct {
	PlayerID       int64  `json:"player_id"`
	Name           string `json:"name"`
	NickName       string `json:"nick_name"`
	Position       int    `json:"position"`
	SecondPosition int    `json:"second_position"`
	Type           int    `json:"type"`
}
