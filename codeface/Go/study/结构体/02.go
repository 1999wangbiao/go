package main

import "fmt"

type student1 struct {
	 name string
	 age int
}

func main()  {
	f:=student1{
		name: "wangbiao",
		age:  12,
	}
	fmt.Println(f)
}