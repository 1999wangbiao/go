package main

import "fmt"

type fun func(int,int) int

func add(x,y int) int {
	return x+y
}
func sub(x,y int) int  {
	return x-y
}
func main()  {
	var c fun//声明一个fun型的变量
	c = add  //给 c赋add函数
	fmt.Printf("type of c:%T\n",c)//type of c:main.fun
	fmt.Println(c(1,2))

	f := sub  //直接给 f赋sub函数
	fmt.Printf("type of f:%T\n",f)//type of f:func(int, int) int
	fmt.Println(f(1,2))

}
