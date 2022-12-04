package main

import "fmt"

// 空接口
type One interface {
}

type Two struct {
	Name string
}

func main() {
	// 直接调用空接口变量赋值
	var one One = 10
	fmt.Printf("%T,%v\n", one, one) // int,10
	// 使用结构体对象赋值
	var two = Two{"Vian"}
	one = two
	fmt.Printf("%T,%v\n", one, one) // main.Two,Vian
	// 同第一种
	var one2 interface{} = 8.8
	fmt.Printf("%T,%v\n", one2, one2) // float64,8.8
}
