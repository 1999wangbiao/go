package main

//第一种方式
/*
var arr [5]int = []int{1,2,3,4,5}
var slice =arr[1:3]
*/
//第二种发生
/*
	var a = make([]string, 5, 10)
*/
//第三种方式
/*
var slice []int = []int {1,2,3}
*/
import "fmt"

func main() {
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a)) //可以切片最大长度
}

//[     0 1 2 3 4 5 6 7 8 9]
//15
//20
