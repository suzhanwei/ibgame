package models

type UserLogin struct {
	Id       int    `json:"id"`
	Uid      string `json:"uid"`
	Name     string `json:"name"`
	Ctime    int64  `json:"ctime"`
	Password string `json:"password"`
}

type RegisterParam struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
