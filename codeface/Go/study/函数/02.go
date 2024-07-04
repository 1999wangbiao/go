package main

import (
	"fmt"
)

func main()  {
	//定义匿名函数
	add := func(x,y int)int {
		  fmt.Println(x+y)
		return x+y
	}
	//使用匿名函数
	a:=add(10,20)
	fmt.Println(a)

	//直接使用匿名函数
	func(x,y int){
		fmt.Println(x+y)
	}(10,20)
}
