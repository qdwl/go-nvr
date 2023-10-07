package model

import "time"

type Role struct {
	Id       int64     `json:"-" gorm:"primarykey"`
	Name     string    `json:"roleName,omitempty"`
	Type     string    `json:"roleType,omitempty"`
	Remark   string    `json:"remark,omitempty"`
	CreateAt time.Time `json:"createAt,omitempty"`
	UpdateAt time.Time `json:"updateAt,omitempty"`
}

func (table Role) TableName() string {
	return "t_role"
}
