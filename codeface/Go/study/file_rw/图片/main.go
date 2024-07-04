package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//编写一个函数实现拷贝
func copyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile,err:=os.Open(srcFileName)
	if err!=nil {
		fmt.Println("open file err=",err)
	}
	//读出
	str:=bufio.NewReader(srcFile)
	defer  srcFile.Close()
	dstFile,err:=os.OpenFile(dstFileName,os.O_RDWR|os.O_CREATE,777)
	if err!=nil {
		fmt.Println("open file err=",err)
	}
	//写入
	reader:=bufio.NewWriter(dstFile)
	defer dstFile.Close()
	//调用copy函数
	return io.Copy(reader,str)

}
func main()  {
	filePath01 := "src/文件/拷贝图片/01.PNG"
	filePath02 := "src/文件/拷贝图片/02.PNG"
	_,err:=copyFile(filePath02,filePath01)
	if err !=nil {
		fmt.Println("copy file err=",err)
	}
}
