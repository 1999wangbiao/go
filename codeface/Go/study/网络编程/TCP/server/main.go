package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn)  {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	var buf [1024]byte
	for  {
		n,err := reader.Read(buf[:])
		recvStr := string(buf[:n])
		fmt.Println("接收到的数据:",recvStr)
		if err== io.EOF {
			break
		}
		if err != nil{
			fmt.Println("read from client failed,err:",err)
			break
		}
	}
}

func main() {
	lister, err := net.Listen("tcp", "127.0.0.1:3000")
	if err != nil {
		fmt.Println("Listen failed err:", err)
		return
	}
	defer lister.Close()
	for {
		conn, err := lister.Accept()
		if err != nil {
			fmt.Println("accept failed ,err:", err)
			continue
		}
		go process(conn)
	}

}
