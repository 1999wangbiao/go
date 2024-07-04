package main

import (
	"fmt"
	"github.com/astaxie/beego/adapter/orm"
	"log"
	"net/http"
	"strconv"
)

func main()  {
	i, _ := strconv.ParseInt("1234",10,64)
	fmt.Println(i)
	orm.Debug = true
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:3351", nil))
	}()

}
