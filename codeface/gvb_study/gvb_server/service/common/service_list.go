package common

import (
	"gvb_server/global"
	"gvb_server/models/system"
)

type Option struct {
	system.Page
	Debug bool
}

func ComList[T any](model T, option Option) (list T, count int64, err error) {

	var imageList []T
	//查询的总共条数
	count = global.DB.Select("id").Find(&imageList).RowsAffected
	offset := (option.Page.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	if !option.Debug {
		err = global.DB.Limit(option.Limit).Offset(offset).Find(&imageList).Error
	} else {
		err = global.DB.Debug().Limit(option.Limit).Offset(offset).Find(&imageList).Error
	}
	return list, count, err
}
