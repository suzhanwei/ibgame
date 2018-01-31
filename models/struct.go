package models

type userinfo struct {
	Uid      int    `json:"uid"`
	Username string `json:"username"`
	Created  int64  `json:"create"`
	Password int64  `json:"password"`
}
