package main

import (
	"fmt"
	"sync"
)

var x1 int64
var wg1 sync.WaitGroup//实现goroutine同步
var lock sync.Mutex//锁

func add1()  {
	for i:=0;i<5000 ;i++  {
		lock.Lock()//加锁
		x1=x1+1
		lock.Unlock()//解锁

	}
	wg1.Done()//goroutine结束就登记-1
}

func main()  {
	wg1.Add(2)
	go add1()
	go add1()
	wg1.Wait()//等待所有登记的goroutine都结束
	fmt.Println(x1)
}
