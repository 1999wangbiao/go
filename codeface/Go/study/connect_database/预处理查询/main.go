package main

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)
type user struct {
	id int
	userName string
	departname string
	time string
}
var db *sql.DB
func init()  {//初始化链接数据库
	var err error
	dsn:="root:root@tcp(127.0.0.1:3306)/go?charset=utf8"
	//不会效验密码是否正确
	db,err = sql.Open("mysql",dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
}
//预处理查询
func prepareQueryDemo()  {
	//1. 输入SQL语句
	sqlStr :="select id,userName,departname,time from user where id>?"
	//2. 进行预处理
	stmt ,err:= db.Prepare(sqlStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()
	//3. 进行处理
	rows,err:=stmt.Query(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next()   {
		var u user
		err := rows.Scan(&u.id,&u.userName,&u.departname,&u.time)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("id:%d, userName:%s departname:%s time:%s\n",u.id,u.userName,u.departname,u.time)
	}
}
func main()  {
	prepareQueryDemo()
}
