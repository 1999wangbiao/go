package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
)
//将消息编码
func Encode(message string)([]byte,error)  {
	var length  =  int32(len(message))
	var pkg = new(bytes.Buffer)
	//写入消息头
	err := binary.Write(pkg,binary.LittleEndian,length)
	if err!= nil {
		return nil,err
	}
	//写入消息实体
	err = binary.Write(pkg,binary.LittleEndian,[]byte(message))
	if err != nil {
		return nil,err
	}
	return pkg.Bytes(),nil
}
//解码消息
func Decode(reader *bufio.Reader) (string,error) {
	//读取消息的长度
	lengthByte,_ := reader.Peek(4)//读取前四个字节数据
	lenghtBuff := bytes.NewReader(lengthByte)
	var length  int32
	err := binary.Read(lenghtBuff,binary.LittleEndian,&length)
	if err != nil {
		return "",err
	}
	//buffered返回缓冲中现有的可读取的字节数
	if int32(reader.Buffered()) < length+4{
		return "", err
	}
	//读取真正的消息数据
	pack := make([]byte,int(4+length))
	_,err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}