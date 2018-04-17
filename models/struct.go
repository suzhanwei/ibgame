package models

type UserLogin struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Uid      string `json:"uid"`
	Ctime    int64  `json:"ctime"`
	Password string `json:"password"`
}

type RegisterParam struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
