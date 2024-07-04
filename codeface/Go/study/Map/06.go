package main

import "fmt"

func main()  {
	mapSlice := make([]map[string]string,3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d  value:%s\n",index,value)
	}
	fmt.Println("after init\n")
	//对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string,10)
	mapSlice[0] ["name"]="小王子"
	mapSlice[0] ["password"] = "123456"
	mapSlice[0] ["address"] = "沙河"
	for index,value:=range mapSlice {
		fmt.Printf("index:%d value:%s\n",index,value)
		for key, v:=range value {
			fmt.Printf("key:%s v:%s\n",key,v)
		}

	}
}
