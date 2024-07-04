package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//客户端(给服务端发送信息)
func main()  {
	//创建链接
	conn,err:=net.Dial("tcp","127.0.0.1:20000")

	if err!=nil {
		fmt.Println("err :",err)
		return
	}
	defer conn.Close()//关闭连接
	// NewReader returns a new Reader whose buffer has the default size.
	//NewReader返回其缓冲区具有默认大小的新读取器
	//从控制台输入数据
	inputReader := bufio.NewReader(os.Stdin)
	for{
		input,_ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input,"\r\n")
		if strings.ToUpper(inputInfo)=="Q" {
			return
		}
		_,err = conn.Write([]byte(inputInfo))
		if err != nil{
			return
		}
		buf := [512]byte{}
		n,err := conn.Read(buf[:])
		if err != nil{
			fmt.Println("recv failed, err:",err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
