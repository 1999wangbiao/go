package main

import (
	"bufio"
	"fmt"
	"net"
)
//服务端(接收客户端的信息)
func process(conn net.Conn){
	defer conn.Close()  //关闭连接
	for {
		reader := bufio.NewReader(conn)//获得数据 reader
		var buf [128]byte
		n, err := reader.Read(buf[:])//读取reader的数据
		if err !=nil {
			fmt.Println("read from failed ,err:",err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据:",recvStr)
		conn.Write([]byte(recvStr))  // 发送 数据  (往链接里写数据 )
	}
}

func main()  {
	listen, err := net.Listen("tcp","127.0.0.1:20000")//建立监听器
	if err!=nil {
		fmt.Println("listen failed,err:",err)
		return
	}
	for {
		conn,err := listen.Accept()
		if err!=nil {
			fmt.Println("accept failed ,err",err)
		    continue
		}
		go process(conn)
	}
}