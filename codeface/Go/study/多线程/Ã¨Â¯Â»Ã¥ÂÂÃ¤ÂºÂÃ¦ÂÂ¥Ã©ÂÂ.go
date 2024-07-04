package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x2 int64
	wg2 sync.WaitGroup
	//lock2 sync.Mutex
	rwlock2 sync.RWMutex
)

func write()  {
	//lock2.Lock()  //加互斥锁
	rwlock2.Lock()//加写锁
	x2 = x2 + 1
	time.Sleep(10*time.Microsecond)//假设读操作耗时10毫秒
	rwlock2.Unlock()//解写锁
	//lock2.Unlock()//解互斥锁
	wg2.Done()
}

func read()  {
	//lock2.Lock()   //加互斥锁
	rwlock2.RLock()//加读锁
	x2 = x2 + 1
	time.Sleep(10*time.Microsecond)//假设读操作耗时1毫秒
	rwlock2.RUnlock()//解读锁
	//lock2.Unlock()//解互斥锁
	wg2.Done()
}
func main()  {
	start := time.Now()//现在时间
	for i := 0;i<1000 ;i++  {
		wg2.Add(1)
		go write()
	}

	for i:=0;i<1000 ;i++  {
		wg2.Add(1)
		go read()
	}

	wg2.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}