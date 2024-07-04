package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func main() {
	stu1 := student{
		Name:  "小王子",
		Score: 90,
	}

	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind()) // student struct
	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		//type StructField struct {
		//    // Name是字段的名字。PkgPath是非导出字段的包路径，对导出字段该字段为""。
		//    // 参见http://golang.org/ref/spec#Uniqueness_of_identifiers
		//    Name    string
		//    PkgPath string
		//    Tag       StructTag // 字段的标签
		//    Offset    uintptr   // 字段在结构体中的字节偏移量
		//    Index     []int     // 用于Type.FieldByIndex时的索引切片
		//    Anonymous bool      // 是否匿名字段
		//}
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
		//name:Name index:[0] type:string json tag:name
		//name:Score index:[1] type:int json tag:score

	}

	// 通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	} //name:Score index:[1] type:int json tag:score
}
