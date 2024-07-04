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
//单行查询
func queryRowDemo()  {
	sqlStr := "select id,userName,departname,time from user where id=?"
	var u user
	//确保queryRow之后调用scan方法,否则持有数据库链接不会释放
	rowdb := db.QueryRow(sqlStr,1)//id==1
	err=rowdb.Scan(&u.id,&u.userName,&u.departname,&u.time)
	if err != nil {
		fmt.Println("a")
		fmt.Println(err)
		fmt.Println("sdf")
		return
	}
	fmt.Printf("id:%d, userName:%s departname:%s time:%s\n",u.id,u.userName,u.departname,u.time)

}
//多行查询
func queryMultiRowDemo()  {
	sqlStr := "select id,userName,departname,time from user where id>?"
	rows,err := db.Query(sqlStr,0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	//循环读取结果集中的数据
	for rows.Next()  {
		var u  user
		err:=rows.Scan(&u.id,&u.userName,&u.departname,&u.time)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(u.id,u.userName,u.departname,u.time)
	}

}
//插入数据
func insertRowDemo()  {
	sqlStr := "insert into user(userName,departname,time) values(?,?,?)"
	ret ,err:=db.Exec(sqlStr,"林国","人事","1290")
	if err != nil {
		fmt.Println(err)
		return
	}
	//在设置主键ID自增的情况下才能显示 id
	theID,err:=ret.LastInsertId()//新插入的数据ID\
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n",theID)
}
//更新数据
func updateRowDemo()  {
	sqlStr:="update user set time=? where userName=?"
	ret ,err:=db.Exec(sqlStr,123,"wangbaio")
	if err != nil {
		fmt.Println(err)
		return
	}
	//影响的行数
	n,err:=ret.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("update  success,affected rows:",n)
}
//删除数据
func deleteRowDemo()  {
	sqlStr := "delete from user where id=?"
	ret ,err:=db.Exec(sqlStr,13)
	if err != nil {
		fmt.Println(err)
		return
	}
	n, err:=ret.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("delete success,affected rows:",n)
}
func main()  {
     //单行查询
	fmt.Println("---单行查询----")
	queryRowDemo()
	//多行查询
	fmt.Println("---多行查询----")
	queryMultiRowDemo()
	//插入数据
	fmt.Println("---插入数据----")
	insertRowDemo()
	//更新函数
	fmt.Println("---更新数据----")
	updateRowDemo()
	//删除数据
	fmt.Println("---删除数据----")
	deleteRowDemo()
}