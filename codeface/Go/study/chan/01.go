package main

import "fmt"

func main(){
	ch1:=make(chan int)
	ch1<-10
	fmt.Println("发送成功")
}


