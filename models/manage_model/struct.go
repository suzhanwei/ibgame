package manage_model

//player base info
type PlayerInfo struct {
	PlayerId       int64  `xorm:"not null pk autoincr BIGINT(20)"`
	Name           string `xorm:"not null default '' unique(uk_name) VARCHAR(128)"`
	NickName       string `xorm:"not null default '' unique(uk_name) VARCHAR(128)"`
	Position       int    `xorm:"not null unique(uk_type_position) INT(11)"`
	SecondPosition int    `xorm:"not null default 0 INT(11)"`
	Type           int    `xorm:"not null default 0 unique(uk_type_position) INT(11)"`
	Score          int    `xorm:"not null default 0 INT(11)"`
	Rebound        int    `xorm:"not null default 0 INT(11)"`
	Assist         int    `xorm:"not null default 0 INT(11)"`
	Steal          int    `xorm:"not null default 0 INT(11)"`
	Cap            int    `xorm:"not null default 0 INT(11)"`
	AppearNum      int    `xorm:"not null default 0 INT(11)"`
}

type AddParam struct {
	Name           string `json:"name"`
	NickName       string `json:"nick_name"`
	Position       int    `json:"position"`
	SecondPosition int    `json:"second_position"`
	Type           int    `json:"type"`
	Score          int    `json:"score"`
	Rebound        int    `json:"rebound"`
	Assist         int    `json:"assist"`
	Steal          int    `json:"steal"`
	Cap            int    `json:"cap"`
	AppearNum      int    `json:"appear_num"`
}

// const(
// 	1="当家球星"
// 	2="得分手"
// 	3="防守者"
// 	4="三分手"
// 	5="组织者"
// 	6="篮板手"
// 	7="第六人"
// 	8="3d"
// 	9="替补"
// 	10="板凳"
// )
