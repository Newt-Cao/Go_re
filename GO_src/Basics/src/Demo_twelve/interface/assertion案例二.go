package main

import "fmt"

type Student struct {
}

func Assertion(anyType ...interface{}) {
	for i, v := range anyType {
		switch v.(type) {
		case bool:
			fmt.Printf("第 %d 个是布尔类型\n", i+1)
		case string:
			fmt.Printf("第 %d 个是字符串类型\n", i+1)
		case int, int32, int64:
			fmt.Printf("第 %d 个是整型\n", i+1)
		case float32, float64:
			fmt.Printf("第 %d 个是浮点型\n", i+1)
		default:
			fmt.Printf("第 %d 个不确定类型\n", i+1)
		}
	}
}

func main() {
	var student Student
	Assertion(10, true, "Vian", student, 66.66)
	/*
		第 1 个是整型
		第 2 个是布尔类型
		第 3 个是字符串类型
		第 4 个不确定类型
		第 5 个是浮点型
	*/
}
