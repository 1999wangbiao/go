package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)
//定义结构体
type charCount struct {
	ChCount int //英文个数
	NumCount int //数字个数
	SpaceCount int //空格个数
	OtherCount int //其他个数
}
//统计文件中数字,英文,空格和其他字符数量
func main()  {
	filePath:="src/实用/文件字符数量/01.txt"
	file,err:=os.Open(filePath)
	if err!=nil {
		fmt.Println("open file err=",err)
		return
	}
	defer file.Close()
	reader:=bufio.NewReader(file)
	var count charCount
	for{
		//一行一行获得文件
		str,err:=reader.ReadString('\n')
		if err==io.EOF {
			break
		}
		//遍历该行
		for _,v := range str{
			switch  {
			case v>='a'&&v<='z':
				fallthrough//穿透处理
			case v>='A'&&v<='Z':
				count.ChCount++
			case v==' '||v=='\t':
				count.SpaceCount++
			case v>='0'&&v<='9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}
	fmt.Println(count.ChCount,count.NumCount,count.SpaceCount,count.OtherCount)

}
