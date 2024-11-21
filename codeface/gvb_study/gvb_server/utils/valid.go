package utils

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

// GetValidMsg 返回结构体中的msg参数
func GetValidMsg(err error, obj any) string {
	//使用时,需要传obj的指针
	typeOf := reflect.TypeOf(obj)
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			//循环每一段错误信息
			//根据报错字段名.获取结构体的具体字段
			if fieldByName, exits := typeOf.Elem().FieldByName(e.Field()); exits {
				msg := fieldByName.Tag.Get("msg")
				return msg
			}
		}
	}
	return err.Error()
}
