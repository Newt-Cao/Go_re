package main

import (
	"fmt"
	"reflect"
)

type user struct {
	Id   int
	Name string
}

func main() {
	// 只是开了一个指针内存，里面并没有地址，无法给字段赋值
	var customer *user
	fmt.Printf("customer: %v\n", customer) // nil
	// 通过反射可以放入地址并赋值，这些都是底层反射实现的
	// 现在底层创建指向该结构体的reflect.Type类型
	customerReflect := reflect.TypeOf(customer) // 获取reflect.Type
	fmt.Println(customerReflect.Kind())         // ptr
	customerReflect = customerReflect.Elem()    // 获取reflect.Type类型指向的类型
	fmt.Println(customerReflect.Kind())         // struct

	// 为reflect.Type对应的类型New一个值
	customerValue := reflect.New(customerReflect) // New返回一个Value类型值，为指向类型reflect.Type申请一个指针，相当于给customerReflect里存了个地址，但是没赋值
	fmt.Println(customerValue.Kind())             // ptr
	fmt.Println(customerValue.Elem().Kind())      // struct
	// 利用断言把创建的实例指向customerValue
	customer = customerValue.Interface().(*user)
	// 通过给底层的指向该实例的类型赋值，达到创建结构体实例的效果
	customerValue.Elem().FieldByName("Id").SetInt(001)
	customerValue.Elem().FieldByName("Name").SetString("Vian")
	fmt.Println(*customer) // {1 Vian}
}
