package models

type userLogin struct {
	Id       int    `json:"id"`
	Uid      int    `json:"uid"`
	Name     string `json:"name"`
	Ctime    int64  `json:"ctime"`
	Password int64  `json:"password"`
}
