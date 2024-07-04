package main
/*
var 变量 chan 元素类型

var ch1 chan int //声明一个传递整形的通道
var ch1 chan bool //声明一个传递布尔值的通道
var ch1 chan []int  //声明一个传递int切片的通道

通道是引用类型,通道类型的空值是 nil
var ch chan int
fmt.Println(ch) //<nil>


声明的 通道后需要使用make 函数初始化之后才能使用
创建channel的格式如下:
   make (chan 元素类型, [缓冲大小])

channel操作
定义
ch := make(chan int)
发送(将一个值发送到通道中)
ch <- 10  //把10发送到 ch 中
接收(从一个通道中接受值)
x:=<-ch
<-ch
关闭
close(ch)

 */
