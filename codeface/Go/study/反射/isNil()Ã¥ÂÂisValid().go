package main

import (
	"fmt"
	"reflect"
)

// IsNil()常被用于判断指针是否为空；IsValid()常被用于判定返回值是否有效。
func main() {
	var a *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil()) // var a *int IsNil: true
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())  //nil IsValid: false
	b := struct{}{}
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid())  //不存在的结构体成员: false
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).MethodByName("abc").IsValid()) //不存在的结构体成员: false
	c := map[string]int{}
	fmt.Println("map中不存在的键:", reflect.ValueOf(c).MapIndex(reflect.ValueOf("那咋")).IsValid()) //map中不存在的键: false
}

//var a *int IsNil: true
//nil IsValid: false
//不存在的结构体成员: false
//不存在的结构体成员: false
//map中不存在的键: false
