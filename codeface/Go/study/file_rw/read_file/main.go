package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)
/*
如果文件不大
ioutil.ReadFile("D:/experiment/goWorkFace/src/文件/读文件/01.txt")
 */
func main()  {
	file,err :=os.Open("D:/experiment/goWorkFace/src/文件/读文件/01.txt")
	if err!=nil {
		fmt.Printf("open file err=",err)
	}
	//关闭
	 defer  file.Close()
	//创建一个*Reader,带缓冲  4096
	reader := bufio.NewReader(file)
	//循环读取文件内容
	for{
		str, err := reader.ReadString('\n')//读到一个换行就结束
		fmt.Print(str)//打印文件内容
		if err == io.EOF {//表示文件末尾
			break
		}
	}
	fmt.Println("---------------------")
	/*
	func ReadFile(filename string) ([]byte, error) {
		return os.ReadFile(filename)
	}
	 */
	str ,err :=ioutil.ReadFile("D:/experiment/goWorkFace/src/文件/读文件/01.txt")
	if err!=nil {
		fmt.Printf("read file err=%v",err)
	}
	//fmt.Printf("%v",str)//[]byte
	fmt.Printf("%s",str)
	fmt.Printf("%v",string(str))
}
