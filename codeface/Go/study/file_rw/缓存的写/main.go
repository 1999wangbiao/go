package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	//创建一个新文件,写入内容 5句"hello"
	// 1, 打开文件 ./02.txtsrc/文件/demo01/01.txt
	file,err :=os.OpenFile("src/文件/创建文件/02.txt",os.O_WRONLY|os.O_CREATE,777)
	if err!=nil {
		fmt.Printf("open file err=%v\n",err)
		return
	}
	//写入
	str := "hello\n"
	//写入时使用带缓存的 *writer
	write:=bufio.NewWriter(file)
	//写入
	for i:=0;i<5 ;i++  {
		write.WriteString(str)
	}
	//因为是缓存,所以要flush
	write.Flush()
	//关闭文件
	defer file.Close()
}
