package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	filePath :="src/文件/追加文件/01.txt"
	file ,err := os.OpenFile(filePath,os.O_WRONLY|os.O_APPEND|os.O_CREATE,777)
	if err !=nil {
		fmt.Println("open file err=",err)
		return
	}
	defer file.Close()
	str:="hello"
	writer := bufio.NewWriter(file)
	writer.WriteString(str)
	writer.Flush()
}
