package main

import "fmt"

type Person struct {
	name string
	sex  string
}

// 方法
func (p Person) eating(food string) string {
	fmt.Println("方法内：", p.name+" is eating "+food)
	return "Beautiful " + food
}

// 方法继承，使用定义了方法的匿名字段即可实现继承
type Teacher struct {
	// 继承Person，Person没有属性名称是匿名字段
	Person
	// 老师教的科目
	subject string
}

// 方法继承
type Student struct {
	// 继承Person
	Person
	// 学生作业
	work string
}

func main() {

	// 实现类
	teacher := Teacher{Person{"Teacher", "male"}, "Chinese"}
	student := Student{Person{"Student", "male"}, "Learn Chinese"}

	// 调用方法
	teacher.eating("Apple")
	student.eating("Banana")
	//方法内： Teacher is eating Apple
	//方法内： Student is eating Banana
}
