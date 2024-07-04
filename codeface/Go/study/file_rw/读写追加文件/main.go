package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main()  {
	filePath := "src/文件/读写追加文件/01.txt"
	file ,err := os.OpenFile(filePath,os.O_CREATE|os.O_APPEND|os.O_RDWR,777)
	if err!=nil {
		fmt.Println("open file err=%v\n",err)
		return
	}
	defer file.Close()
	//读文件
	reader := bufio.NewReader(file)
	for  {
		str,err:=reader.ReadString('\n')
		fmt.Print(str)
		if err== io.EOF {
			fmt.Print("文件末尾")
			break
		}
	}
	//追加写文件
	writer := bufio.NewWriter(file)
	str:="wangbiao"
	writer.WriteString(str)
	writer.Flush()

	//reader1 := bufio.NewReader(file)
	//for  {
	//	str,err:=reader1.ReadString('\n')
	//	fmt.Print(str)
	//	if err== io.EOF {
	//		fmt.Print("文件末尾")
	//		break
	//	}
	//}

}
