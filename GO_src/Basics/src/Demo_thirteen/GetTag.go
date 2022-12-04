package main

import (
	"fmt"
	"reflect"
)

type Human struct {
	Name string `json:"姓名"`
	Age  int    `json:"年龄"`
}

func main() {
	// 实例
	Human := Human{}
	// 判断该实例是否为结构体
	if reflect.TypeOf(Human).Kind() == reflect.Struct {
		// 如果是则转换成reflect.Type类型
		HumanReflect := reflect.TypeOf(Human)
		// 统计该结构体字段数量
		HumanFieldNum := HumanReflect.NumField()
		// 遍历输出字段对应的标签
		for i := 0; i < HumanFieldNum; i++ {
			fmt.Printf("%s\n", HumanReflect.Field(i).Tag.Get("json"))
			/*
				姓名
				年龄
			*/
		}
	}
}
