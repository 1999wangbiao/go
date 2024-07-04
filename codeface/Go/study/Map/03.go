package main

import "fmt"
//  value, ok := map[key]
func main()  {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["李四"] = 100
	//如果key存在ok为true,v为对应的值; 不存在ok为false, v为值类型的零值
	v,ok :=scoreMap["王彪"]
	if ok {
		fmt.Println(v)
	}else {
		fmt.Println("查无此人")
	}

	for k, vaule := range scoreMap  {
		fmt.Printf("k=%s   vaule=%d\n",k,vaule)
	}
}
