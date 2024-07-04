package main

import (
	"fmt"
	"strconv"
	"sync"
)

//存在问题不安全  fatal error: concurrent map writes
//var m = make(map[string]int)
//func get(key string) int {
//	return m[key]
//}
//
//func set(key string,value int)  {
//	m[key] = value
//}
//func main()  {  // 存在问题不安全  fatal error: concurrent map writes
//	wg3 := sync.WaitGroup{}
//	for i:=0;i<20;i++ {
//		wg3.Add(1)
//		go func(n int) {
//			key := strconv.Itoa(n)
//			set(key,n)
//			fmt.Printf("k=:%v,v:=%v\n",key,get(key))
//			wg3.Done()
//		}(i)
//	}
//	wg3.Wait()
//}
var m = sync.Map{}

func main() {
	wg3 := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg3.Add(1)
		go func(n int) {
			key := strconv.Itoa(n) //strconv.Itoa()函数的参数是一个整型数字,它可以将数字转换成对应的字符串类型的数字
			m.Store(key, n)
			//sync.Map 不能使用 map 的方式进行取值和设置等操作，而是使用 sync.Map 的方法进行调用，
			// Store 表示存储，Load 表示获取，Delete 表示删除。
			//使用 Range 配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值，
			// Range 参数中回调函数的返回值在需要继续迭代遍历时，返回 true，终止迭代遍历时，返回 false。
			value, _ := m.Load(key)
			fmt.Printf("k:=%v,v:=%v\n", key, value)
			wg3.Done()
		}(i)
	}
	wg3.Wait()
}
