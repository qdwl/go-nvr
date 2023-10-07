package repository

import (
	"github.com/qdwl/go-nvr/nvr/internal/model"
	"github.com/zeromicro/go-zero/core/logx"
)

func AddChannel(channel *model.Channel) error {
	err := model.DB.Create(channel).Error
	if err != nil {
		logx.Error("[DB ERROR] : ", err)
		return err
	}

	return err
}

func DeleteChannel(id int64) error {
	err := model.DB.Delete(&model.Channel{}, id).Error
	if err != nil {
		logx.Error("[DB ERROR] : ", err)
		return err
	}

	return err
}

func UpdateChannel(channel *model.Channel) error {
	err := model.DB.Save(channel).Error
	if err != nil {
		logx.Error("[DB ERROR] : ", err)
		return err
	}

	return err
}

func GetChannelList(page int, size int, channel *model.Channel) (count int64, list []*model.Channel, err error) {
	list = make([]*model.Channel, 0)
	err = model.DB.Model(&model.Channel{}).Where(channel).Count(&count).Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		logx.Error("[DB ERROR] : ", err)
		return
	}

	return
}
