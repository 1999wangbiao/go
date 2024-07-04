package main

import "fmt"

//方法的定义实例
//person 是一个结构体
type person struct {
	name string
	age int8
}

//newperson 是一个person类型的构造函数
func newperson(name string,age  int8) *person  {
	return &person{
		name: name,
		age:  age,
	}
}
//dream 是为person类型定义方法
func (p person)dream()  {
	fmt.Printf("%s的梦想是学好Go语言!\n",p.name)
}
//setAge 是为person修改年龄的方法
func (p *person)setAge(newAge int8)  {
	p.age=newAge
}

func main()  {
	p1 :=newperson("王彪",18)
	p1.dream()
	fmt.Println(*p1)
	p1.setAge(20)
	fmt.Println(p1.age)
	fmt.Println(*p1)
}
