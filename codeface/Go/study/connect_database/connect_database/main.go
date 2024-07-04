package main
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)


func main()  {
	//连接数据库
	dsn:="root:root@tcp(127.0.0.1:3306)/go?charset=utf8"
	db,err:=sql.Open("mysql",dsn)
	if err != nil {
		fmt.Println("连接数据库失败,err",err)
		//panic(err)
	}
	fmt.Println("连接成功")
	res , err := db.Exec("insert into user(id, userName, departname,time)values (1,'wangbaio','技术部门','1234')")
	fmt.Println("aa")
	if err != nil {
		fmt.Println("sdf")
	}
	id,_ := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
}
