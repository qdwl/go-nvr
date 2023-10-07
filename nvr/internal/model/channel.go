package model

type Channel struct {
	Id        int64   `json:"-" gorm:"primarykey"`
	Name      string  `json:"name,omitempty"`
	Enable    int64   `json:"enable,omitempty"`
	Ip        string  `json:"ip,omitempty"`
	Port      uint32  `json:"port,omitempty"`
	UserName  string  `json:"userName,omitempty"`
	Password  string  `json:"password,omitempty"`
	Protocol  string  `json:"protocol,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	Remark    string  `json:"remark,omitempty"`
}

func (table Channel) TableName() string {
	return "t_channel"
}
