package main

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() (err error){
	dsn:="root:root@tcp(127.0.0.1:3306)/go?charset=utf8"
	//不会效验密码是否正确
	db,err = sql.Open("mysql",dsn)
	if err != nil {
		return err
	}
	//尝试与数据库建立连接(效验dsn是否正确)
	err = db.Ping()
	if err != nil{
		return err
	}
	return err
}

func main()  {
	err := initDB()//调用初始数据库的函数
	if err != nil {
		fmt.Println("init db failed,err=",err)
		return
	}
	fmt.Println("连接成功")
}
