package main

import (
	"fmt"
	"reflect"
)

type myInt int64

func reflectType2(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
}

func main() {
	var a *float32 //指针
	var b myInt    //自定义类型
	var c rune     //类型别名

	reflectType2(a) //type: kind:ptr
	reflectType2(b) //type:myInt kind:int64
	reflectType2(c) //type:int32 kind:int32

	type person struct {
		name string
		age  int
	}
	type book struct {
		title string
	}
	var d = person{
		name: "沙河小王子",
		age:  18,
	}

	var e = book{title: "跟着小王子"}
	reflectType2(d) //type:person kind:struct
	reflectType2(e) //type:book kind:struct
}
