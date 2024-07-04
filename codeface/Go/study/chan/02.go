package main

import (
	"fmt"
)

func recv(c chan int)  {
	ret := <-c
	fmt.Println("接收成功",ret)
}
func main()  {
	ch := make(chan int, 1)
	ch <- 10
	go recv(ch)
	//time.Sleep(time.Second)
	fmt.Println("发送成功")
	//time.Sleep(time.Second)
}
