package main

import "fmt"

func main()  {
	//定义一个空接口
	var x interface{}
	s := "hello"
	x = s
	fmt.Printf("type:%T value:%v\n",x,x)
	i := 100
	x = i
	fmt.Printf("type:%T value:%v\n",x,x)
	b := true
	x = b
	fmt.Printf("type:%T value:%v\n",x,x)
}
