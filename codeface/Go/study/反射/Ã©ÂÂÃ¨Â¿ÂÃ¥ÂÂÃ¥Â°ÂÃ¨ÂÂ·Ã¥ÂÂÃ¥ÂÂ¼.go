package main

import (
	"fmt"
	"reflect"
)

func reflectVaule3(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		//v.Int()从反射中获取整型的原始值,然后通过int64()强制转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		//v.Float()从反射中获取浮点型的原始值,然后通过float32()强制转换
		fmt.Printf("type is float32,value is %f\n", float32(v.Float()))
	case reflect.Float64:
		//v.Float从反射中获取浮点型的原始值,然后通过float64()强制转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

func main() {
	var a float32 = 3.14
	var b int64 = 100

	reflectVaule3(a)
	reflectVaule3(b)
	c := reflect.ValueOf(10)
	fmt.Println(c.Kind())
	fmt.Printf("type c :%T\n", c)
	fmt.Println(c)
	var d int64 = 100
	e := reflect.ValueOf(&d).Elem()
	fmt.Println("type of p:", e.Type())
	fmt.Println("setTability of p:", e.CanSet())

	e.SetInt(13)
	fmt.Println(e.Interface())
	fmt.Println(d)
}
