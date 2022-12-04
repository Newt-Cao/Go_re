/* 多个 defer 时的执行顺序 */

package main

import "fmt"

// 先入栈，后出栈，顺序刚好相反。

func func1() {
	fmt.Println("A")
}

func func2() {
	fmt.Println("B")
}

func func3() {
	fmt.Println("C")
}

func main() {
	defer func1()
	defer func2()
	defer func3()
}

// 运行顺序 3 - 2 - 1
