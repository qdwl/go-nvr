package model

import "time"

type User struct {
	Id       int64     `json:"-" gorm:"primarykey"`
	UserId   string    `json:"userId,omitempty" gorm:"unique"`
	UserName string    `json:"username,omitempty"`
	Password string    `json:"password,omitempty"`
	RoleId   int64     `json:"roleId"`
	Remark   string    `json:"remark,omitempty"`
	LoginAt  time.Time `json:"loginAt,omitempty"`
	CreateAt time.Time `json:"createAt,omitempty"`
	UpdateAt time.Time `json:"updateAt,omitempty"`
}

func (table User) TableName() string {
	return "t_user"
}
