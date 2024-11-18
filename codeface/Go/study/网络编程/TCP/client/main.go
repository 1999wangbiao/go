package main

import (
	"fmt"
	"net"
)

// 客户端
func main() {
	conn, err := net.Dial("tcp", "192.168.1.253:1030")
	if err != nil {
		fmt.Println("dial failed ,err", err)
		return
	}
	defer conn.Close()

	for i := 0; i < 10; i++ {
		msg := `Hello, Hello. How are you?\n`
		conn.Write([]byte(msg))
	}

}
