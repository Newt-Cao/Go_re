// 匿名函数的几种使用方式

package main

import (
	"fmt"
)

// 全局匿名函数
var res2 = func(n5, n6 int) int {
	return n5 * n6
}

func main() {
	// 第一种，直接调用
	func(n1, n2 int) int {
		return n1 + n2
	}(1, 2)

	// 第二种，赋值给一个变量，变量调用
	res := func(n3, n4 int) int {
		return n3 - n4
	}

	fmt.Printf("res: %v\n", res(2, 1))

	// 第三种，定义全局匿名函数
	fmt.Printf("res2: %v\n", res2(1, 2))
}
