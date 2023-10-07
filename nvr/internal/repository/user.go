package repository

import (
	"github.com/qdwl/go-nvr/nvr/internal/model"
	"github.com/zeromicro/go-zero/core/logx"
)

func AddUser(user *model.User) error {
	err := model.DB.Create(user).Error
	if err != nil {
		logx.Error("[DB ERROR] : ", err)
		return err
	}

	return err
}

func GetUserByName(name string) (user *model.User, err error) {
	user = new(model.User)
	err = model.DB.Where("user_name = ?", name).Take(user).Error
	if err != nil {
		logx.Error("[DB ERROR] : ", err)
		return
	}

	return
}

func GetUserList(page int, size int) (count int64, list []*model.User, err error) {
	list = make([]*model.User, 0)
	err = model.DB.Model(new(model.User)).Count(&count).Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		logx.Error("[DB ERROR] : ", err)
		return
	}

	return
}

func DeleteUser(id int64) error {
	err := model.DB.Delete(&model.User{}, id).Error
	if err != nil {
		logx.Error("[DB ERROR] : ", err)
		return err
	}

	return err
}

func UpdateUser(user *model.User) error {
	err := model.DB.Save(user).Error
	if err != nil {
		logx.Error("[DB ERROR] : ", err)
		return err
	}

	return err
}
