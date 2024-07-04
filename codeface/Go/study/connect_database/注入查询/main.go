package main

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)
var db *sql.DB
var err  error
type user struct {
	id int
	userName string
	departname string
	time string
}
//初始化链接数据库
func init()  {
	dsn:="root:root@tcp(127.0.0.1:3306)/go?charset=utf8"
	//不会效验密码是否正确
	db,err = sql.Open("mysql",dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func sqlInjectDemo(name string)  {
	sqlStr := fmt.Sprintf("select id,userName ,departname,time from user where userName='%s'",name)
	fmt.Printf("SQL:%s\n",sqlStr)
	rows,err := db.Query(sqlStr)
	for  rows.Next(){
		var u user
		err=rows.Scan(&u.id,&u.userName,&u.departname,&u.time)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(u.id,u.userName,u.departname,u.time)
	}
	//rows,err = db.QueryRow(sqlStr).Scan(&u.id,&u.userName,&u.departname,&u.time)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(u)
}
func main()  {
	fmt.Println("正常查询")
	sqlInjectDemo("林国")
	fmt.Println("注入查询")
	sqlInjectDemo("xxx' or 1=1 #")
}
