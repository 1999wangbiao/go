package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
//定义学生基本信息
type Student struct {
	id int
	name string
	age int
}
type Class struct {
	 Map map[int] *Student

}
func getInput() string {
	//使用os.Stdin开启输入流
	//函数原型 func NewReader(rd io.Reader) *Reader
	//NewReader创建一个具有默认大小缓冲、从r读取的*Reader 结构见官方文档
	in := bufio.NewReader(os.Stdin)
	//in.ReadLine函数具有三个返回值 []byte bool error
	//分别为读取到的信息 是否数据太长导致缓冲区溢出 是否读取失败
	str, _, err := in.ReadLine()
	if err != nil {
		return err.Error()
	}
	return string(str)
}
//添加学生信息
func (c *Class)AddStudent(){
	var (
		name  string
	)
	fmt.Println("输入学生姓名")
	name = getInput()
	fmt.Println("输入学生年龄")
	age, _ := strconv.Atoi(getInput())
chongxin:
      fmt.Println("输入学生ID")
      id, _ := strconv.Atoi(getInput())
	stu := &Student{
		id:   id,
		name: name,
		age:age,
	}
	_,err:=c.Map[id]
	if err {
		fmt.Println("已存在该学生id")
		goto chongxin
	}
	c.Map[id]=stu
	fmt.Println("添加成功")
}


//查看 学生信息
func (c *Class)ShowStudent(){
	for _,stu:=range c.Map {
		fmt.Printf("学号:%d 姓名:%s 年龄:%d\n",stu.id,stu.name,stu.age)
	}
}
//修改学生信息
func (c *Class)UpdateStudent(){
	idChong:
		fmt.Println("输入要修改的学生ID")
		id, _ := strconv.Atoi(getInput())
		_,err:=c.Map[id]
	if !err {
		fmt.Println("不存在该ID请重新输入")
		goto idChong
	}
		fmt.Println("输入修改后的姓名")
		name := getInput()
		fmt.Println("输入修改后的年龄")
	    age, _ := strconv.Atoi(getInput())
		stu:=&Student{
			id:  id,
			name: name,
			age:  age,
		}
		c.Map[id]=stu
		fmt.Println("修改成功")
}
//删除学生信息
func (c *Class)DeleteStudent(){
	Deleteid:
	fmt.Println("输入要删除的学生ID")
	id, _ := strconv.Atoi(getInput())
	_,err:=c.Map[id]
	if !err {
		fmt.Println("不存在该ID请重新输入")
		goto Deleteid
	}
	delete(c.Map,id)
	fmt.Println("删除成功")
}
//菜单函数
func show()  {
	fmt.Println(`
	欢迎来到学生管理系统!
	1. 查看所有学生
	2. 新增学生信息
	3. 修改学生信息
	4. 删除学生信息
	5. 退出
	请选择您需要的功能:
	`)
}
func main()  {
	c:=Class{}
	c.Map=make(map[int]*Student,50)
	show:
	show()
	var tmp  int
	fmt.Scanln(&tmp)
	switch tmp {
	case 1:
			c.ShowStudent()
			goto show
	case 2:
			c.AddStudent()
			goto show
	case 3:
			c.UpdateStudent()
			goto show
	case 4:
			c.DeleteStudent()
			goto show
	case 5:

	default:
		fmt.Println("输入错误,重新输入")
		goto show
	}

}
