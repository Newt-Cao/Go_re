// 闭包（知识盲区）

package main

import "fmt"

// 实现一个简单的闭包，累加器
func AddUpper() func(int) int {
	var num int = 10
	return func(i int) int {
		num = num + i
		return num
	}
}

func main() {

	f := AddUpper()

	fmt.Printf("f: %v\n", f(1))
	fmt.Printf("f: %v\n", f(2))

}
