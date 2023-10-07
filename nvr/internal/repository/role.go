package repository

import (
	"github.com/qdwl/go-nvr/nvr/internal/model"
	"github.com/zeromicro/go-zero/core/logx"
)

func AddRole(role *model.Role) error {
	err := model.DB.Create(role).Error
	if err != nil {
		logx.Error("[DB ERROR] : ", err)
		return err
	}

	return err
}

func GetRoleList(page int, size int) (count int64, list []*model.Role, err error) {
	list = make([]*model.Role, 0)
	err = model.DB.Model(new(model.Role)).Count(&count).Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		logx.Error("[DB ERROR] : ", err)
		return
	}

	return
}

func DeleteRole(id int64) error {
	err := model.DB.Delete(&model.Role{}, id).Error
	if err != nil {
		logx.Error("[DB ERROR] : ", err)
		return err
	}

	return err
}

func UpdateRole(role *model.Role) error {
	err := model.DB.Save(role).Error
	if err != nil {
		logx.Error("[DB ERROR] : ", err)
		return err
	}

	return err
}
