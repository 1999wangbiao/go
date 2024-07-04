package main

import (
	"fmt"
	"reflect"
)

func reflectSetValue1(x interface{})  {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64{
		v.SetInt(200)//修改的是副本, reflect 包会引发 panic
	}
}

func reflectSetValue2(x interface{})  {
	v := reflect.ValueOf(x)
	// 反射中使用Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64{
		v.Elem().SetInt(200)
	}
}

func main(){
	var a int64 = 100
	//reflectSetValue1(&a)
	reflectSetValue2(&a)
	fmt.Println(a)
}