package main

import (
	"fmt"
	"reflect"
)

type Order struct {
	Id         int
	customerId int
}
	
func reflectQuery(x interface{}) {
	oid := reflect.TypeOf(x)
	cid := reflect.ValueOf(x)
	fmt.Printf("Type = %v\nValue = %v\nRealType = %T\nRealType = %T", oid, cid, oid, cid)
}

func main() {
	//实例化
	order := Order{
		Id:         10,
		customerId: 20,
	}
	// 调用
	reflectQuery(order)
}
