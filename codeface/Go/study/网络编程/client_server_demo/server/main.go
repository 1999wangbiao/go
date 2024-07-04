package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"proto/proto"
)

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		fmt.Println("收到clien发来的数据:", msg)
		if err != nil {
			fmt.Println("decode mag failed err:", err)
			return
		}
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed ,err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen failed, err:", err)
			continue
		}
		go process(conn)

	}

}
