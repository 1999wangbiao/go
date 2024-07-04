package main

import (
	"fmt"
	"net"
	"网络/解决粘包/proto"
)

func main()  {
	conn,err := net.Dial("tcp","127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err:",err)
		return
	}
	for i:=0;i<10 ;i++  {
		msg := `hello,Hello . Hoe are you?`
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed,err:",err)
			return
		}
		conn.Write(data)
	}
	defer conn.Close()
}
