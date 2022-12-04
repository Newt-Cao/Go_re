package main

import (
	"fmt"
	"reflect"
)

type Stu struct {
	Name  string
	Skill string
}

func main() {
	stu := Stu{"Niko", "Fly"}
	// 判断种类是否为struct
	if reflect.TypeOf(stu).Kind() == reflect.Struct {
		stuReflectType := reflect.TypeOf(stu)
		stuReflectValue := reflect.ValueOf(stu)
		// 结构体字段数
		numField := stuReflectValue.NumField()
		fmt.Println("Number of fields = ", numField, "\n") // Number of fields =  2
		// 遍历获取字段对应名字、值、类型
		for i := 0; i < numField; i++ {
			fmt.Printf("Field Name = %v\nField Value = %v\nField Type = %T\n\n", stuReflectType.Field(i).Name, stuReflectValue.Field(i), stuReflectValue.Field(i))
		}
	}
}
