package main

import (
	"encoding/json"
	"fmt"
)

//学生
type Student struct {
	ID int//id
	Gender string//性别
	Name string//姓名
}
//班级
type Class struct {
	Title string //班号
	Students []Student//学生信息数组
}

func main()  {
	//创建一个班级变量c
	c := &Class{
		Title:   "10班", //10班
		Students: make([]Student,0,200),// 创建200个学生容量
	}
	//在班级中添加学生
	for i:=0;i<10 ;i++  {
		//编写学生
		stu := &Student{
			ID:     i,
			Gender: "男",
			Name:   fmt.Sprintf("stu%02d",i),
		}
		//添加学生
		c.Students = append(c.Students,*stu)//将新信息追加到集合
	}
	fmt.Println(*c)
	//JSON序列化:结构体-->JSON格式的字符串
	data,err := json.Marshal(c)
	if err != nil{
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n",data)
	//JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"10班","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	var c1 Class
	err = json.Unmarshal([]byte(str),&c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n",c1)
}
