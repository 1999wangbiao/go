package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{})  {
	//我不知道别人调用这个函数是会传进来什么类型的变量
	 obj := reflect.TypeOf(x)
	fmt.Printf("type:%v\n",obj)
}


func main()  {
	var a float32 = 1.234
	reflectType(a)
	var b int = 1234
	reflectType(b)
}
