package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main()  {
	filePath01 :="src/实用/复制文件/01.txt"
	filePath02 :="src/实用/复制文件/02.txt"
	file01,err01:=os.OpenFile(filePath01,os.O_RDWR,777)
	file02,err02:=os.OpenFile(filePath02,os.O_RDWR,777)
	if err01!= nil {
		fmt.Println("open o1.txt err=",err01)
		return
	}
	if err02!= nil {
		fmt.Println("open o1.txt err=",err02)
		return
	}
	defer file01.Close()
	defer file02.Close()
	//读 01.txt
	reader01:=bufio.NewReader(file01)
	for  {
		str01,err:=reader01.ReadString('\n')
		//写入02.txt
		writer01:=bufio.NewWriter(file02)
		writer01.WriteString(str01)
		writer01.Flush()
		if err==io.EOF {
			break
		}
	}

}
