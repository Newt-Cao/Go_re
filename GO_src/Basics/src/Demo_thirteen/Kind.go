package main

import (
	"fmt"
	"reflect"
)

type Order struct {
	Id         int
	customerId int
}

func main() {
	order := Order{}
	// 类型
	oType := reflect.TypeOf(order)
	// 种类
	oKind := oType.Kind()
	fmt.Printf("order Type = %v\norder Kind = %v", oType, oKind)
}
/*
order Type = main.Order
order Kind = struct
*/
