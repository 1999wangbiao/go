package main

import (
	"database/sql"
	"fmt"
    _"github.com/go-sql-driver/mysql"
)
var db *sql.DB
var err  error
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
func transactionDemo()  {
	tx, err := db.Begin()//开启事务
	if err != nil {
		if tx!=nil {
			tx.Rollback()//回滚
		}
		fmt.Println(err)
		return
	}
	sqlStr1 := "update user set userName='王彪' where id=?"
	ret1,err:=tx.Exec(sqlStr1,2)
	if err != nil {
		tx.Rollback()//回滚
		fmt.Println(err)
		return
	}
	affRow1, err:=ret1.RowsAffected()
	if err != nil {
		tx.Rollback()//回滚
		fmt.Println(err)
		return
	}

	sqlStr2 := "update user set userName='熊澜' where id=?"
	ret2,err:=tx.Exec(sqlStr2,0)
	if err != nil {
		tx.Rollback()//回滚
		fmt.Println(err)
		return
	}
	affRow2, err:=ret2.RowsAffected()
	if err != nil {
		tx.Rollback()//回滚
		fmt.Println(err)
		return
	}
	fmt.Println(affRow1,affRow2)
	if affRow1==1&&affRow2==1{
		fmt.Println("事务提交啦....")
		tx.Commit()
	}else {
		tx.Rollback()
		fmt.Println("事务回滚啦....")
	}
	fmt.Println("exec trans success!")
}
func main()  {
	transactionDemo()
}