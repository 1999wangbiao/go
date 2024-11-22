package common

import (
	"gvb_server/global"
	"gvb_server/models/system"
)

type Option struct {
	system.Page
	Debug bool
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {

	//查询的总共条数
	count = global.DB.Select("id").Find(&list).RowsAffected
	offset := (option.Page.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	if option.Sort == "" {
		option.Sort = "created_at asc" //设置默认排序
	}
	//是否调试
	if !option.Debug {
		err = global.DB.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error
	} else {
		err = global.DB.Debug().Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error
	}

	/* 配合前面的广告查询是否is_show:true

	//查询的总共条数
	query := global.DB.Where(model)
	count =query.Select("id").Find(&list).RowsAffected
	offset := (option.Page.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	if option.Sort == "" {
		option.Sort = "created_at asc" //设置默认排序
	}
	query := global.DB.Where(model)
	//是否调试
	if !option.Debug {
		err = global.DB.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error
	} else {
		err = global.DB.Debug().Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error
	}
	*/

	return list, count, err
}
