package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)
func main()  {
	rand.Seed(time.Now().UnixNano())
	scoreMap := make(map[string] int,200)
	for i:=0; i<100; i++  {//给map符100个值
		key :=fmt.Sprint("stu%02d",i)//生成stu开头的字符串
		value := rand.Intn(100)//生成0-99的随机数
		scoreMap[key]=value
	}
	//取出map中的所有key存入切片keys
	keys := make([]string,0,200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	for _,key :=range keys {
		fmt.Println(key,scoreMap[key])

	}
}
