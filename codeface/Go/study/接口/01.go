package main

import "fmt"

/*
type 接口类型名 interface{
	方法名 1 (参数列表 1 ) 返回值列表 1
	方法名 2 (参数列表 2 ) 返回值列表 2
	......
}
 */

type Cat struct {
}

func (c Cat)Say() string {
	return "喵喵喵"
}
type Dog struct {
}
func (d Dog)Say() string {
	return "汪汪汪"
}
func main()  {
	c:= Cat{}
	fmt.Printf("猫:%s\n",c.Say())
	d:=Dog{}
	fmt.Printf("狗:%s\n",d.Say())
}
