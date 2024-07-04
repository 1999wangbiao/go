package main

import "fmt"
//结构体的继承
type Animal struct {
	name string
}

func (a *Animal)move()  {
	fmt.Printf("%s会动\n",a.name)
}

type Dog struct {
	feet int
	*Animal//匿名嵌套
}

func (d *Dog)wang()  {
	fmt.Printf("%s会汪汪汪叫\n",d.name)
}
func main()  {
	dog := &Dog{
		4,
		&Animal{
			"乐乐"},
	}
	dog.move()
	dog.wang()
}