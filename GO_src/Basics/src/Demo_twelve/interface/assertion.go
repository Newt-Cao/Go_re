package main

import "fmt"

func main() {
	// 断言机制
	var assertion interface{}
	var num int = 10
	assertion = num
	fmt.Printf("%T,%v", assertion, assertion) // int,10
	// 强转后赋值
	var num2 int
	num2 = assertion.(int)
	fmt.Printf("%T,%v", num2, num2) // int,10
	// 判断
	if num2, ok := assertion.(int); ok {
		fmt.Println("转换成功！", num2)
	} else {
		fmt.Println("转换失败！")
	}
}
