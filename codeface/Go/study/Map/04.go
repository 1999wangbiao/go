package main

import "fmt"

func main()  {
	scoreMap := make(map[string] int)
	scoreMap["张三"] = 90
	scoreMap["李四"] = 100
	scoreMap["小明"] = 60
	for k,v := range scoreMap {
		fmt.Printf("key=%s  vaule=%d\n" ,k,v)
	}
	delete(scoreMap,"小明")
	for k,v := range scoreMap {
		fmt.Printf("key=%s  vaule=%d\n" ,k,v)
	}
}
