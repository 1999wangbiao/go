package main

import "fmt"

func addr() (func(int) int)  {
	var  x int
	return func(y int) int {
		x+=y
		return x
	}
}
func main()  {
	var  f = addr()
	f1 := addr()
	fmt.Println(f(10))
	fmt.Println("f1:",f1(10))
	fmt.Println("f1:",f1(10))
	fmt.Println(f(10))
	fmt.Println(f(10))
	fmt.Println(f(10))
	fmt.Println(f(10))
fmt.Printf("type of f:%T\n",f)

	fmt.Println(f1(10))
	fmt.Println(f1(10))
	fmt.Println(f1(10))
	fmt.Printf("type of f1:%T",f1)
}
