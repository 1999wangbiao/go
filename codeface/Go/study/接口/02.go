package main

//一个类型实现多接口

import "fmt"
//Mover 接口
type Mover interface {
	move()
}
//Sayer 接口
type Sayer interface {
	say()
}
//定义dog结构体
type dog struct {
	name string
}

//实现Mover接口
func (d dog) move()  {
        fmt.Printf("%s会动\n",d.name)
}
//实现Sayer接口
func (d dog) say()  {
	fmt.Printf("%s会汪汪汪叫\n",d.name)
}
func main()  {
	var x  Mover
	var y Sayer

	var a = dog{"旺财"}
	x = a
	y = a

	x.move()
	y.say()
}
