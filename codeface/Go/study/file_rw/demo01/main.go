package main

import (
	"fmt"
	"os"
)

func main()  {
	file,err :=os.Open("D:/experiment/goWorkFace/src/文件/demo01/01.txt")
	if err!=nil {
		fmt.Printf("open file err=",err)
	}
	fmt.Printf("file is %v",file)
    file.WriteString("万是一个")
	file.Close()
}
