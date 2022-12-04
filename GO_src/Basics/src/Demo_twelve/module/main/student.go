package main

import (
	"factory"
	"fmt"
)

func main() {
	// 当factory里结构体Students为大写时，直接调用
	stu := factory.Students{"Vian", 59, 99.9}
	fmt.Println(stu.Name) // Vian

	// 当factory里结构体students为小写时，需要调用该包的大写函数间接调用小写结构体
	stu2 := factory.StudentsInterface("Vian", 59, 99.9)
	fmt.Println((*stu2).Name) // Vian

	// 当factory里结构体students为小写，结构体里字段也为小写时，需要绑定一个结构体大写函数间接调用小写结构体和结构体里小写字段
	fmt.Println(stu2.AgeField()) // 59

}
