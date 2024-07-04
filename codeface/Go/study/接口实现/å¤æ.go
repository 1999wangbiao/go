package main

import "fmt"

// 接口是一组方法的集合
/*
// 接口定义，如果返回值参数列表没有返回值，或者只有一个返回值，可以不要括号“()”
type 接口名称  interface {
	方法1(输入参数列表1) (返回值参数列表1)
    方法2(输入参数列表2) (返回值参数列表2)
	……
	方法N(输入参数列表N) (返回值参数列表N)
}

// 接口实现，如果返回值参数列表没有返回值，或者只有一个返回值，可以不要括号“()”
func (接收者变量, 接收者类型) 方法1(输入参数列表1) (返回值参数列表1){

}
func (接收者变量, 接收者类型) 方法2(输入参数列表2) (返回值参数列表2){

}
……
func (接收者变量, 接收者类型) 方法N(输入参数列表N) (返回值参数列表N){

}

*/

// 定义接口
type Person1 interface {
	eating(food string) string
}

// 定义结构体
type Teacher1 struct {
	name string
}

type Student1 struct {
	name string
}

// 使用Teacher结构体实现方法
func (t Teacher1) eating(food string) string {
	fmt.Println("The " + t.name + " eating " + food)
	return "Good teacher"
}

// 使用Student结构体实现方法
func (s Student1) eating(food string) string {
	fmt.Println("The " + s.name + " eating " + food)
	return "Good student"
}

func main() {

	var person Person1

	teacher := Teacher1{"Teacher"}
	student := Student1{"Student"}

	// 实现多态
	person = teacher
	person.eating("Apple")

	person = student
	person.eating("Banana")

}
