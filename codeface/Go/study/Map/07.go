package main

import "fmt"

func main()  {
	sliceMap := make(map[string][]string,3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key:="中国"
	value,ok :=sliceMap[key]
	if !ok {
		value = make([]string,0,2)
	}
	value = append(value,"北京","上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}
