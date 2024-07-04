package main

import "fmt"
//for range 从通道循环取值
func main()  {
	ch1 := make(chan int)
	ch2 := make (chan int)
	//开启 goroutine将0-100 的数发送到ch1
	go func() {
		for i:=0;i<100 ;i++  {
			ch1<-i
		}
		close(ch1)
	}()
	go func() {
		for{
			i,ok:=<-ch1
			if !ok {
				break
			}
			ch2<- i*i
		}
		close(ch2)
	}()
	for i:= range ch2 {
		fmt.Println(i)
	}
}


