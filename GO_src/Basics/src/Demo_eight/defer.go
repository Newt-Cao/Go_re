/* 关键字 defer */

package main

import "fmt"

// 如果一个函数体含有 defer关键字，则当该函数所有的方法执行完，最后再执行defer，相当于在右括号之后。(一个函数体内可有多个defer，defer后可连接函数，也可嵌套。)

// defer在执行中会同时把值拷贝到defer栈中，后续修改值时不改变defer中的值
func sum(n1, n2 int) int {
	n1++
	n2++
	res := n1 + n2
	fmt.Printf("res: %v\n", res)
	defer fmt.Println(n1)
	defer fmt.Println(n2)
	return res
}

func main() {
	defer fmt.Println("Game Over")

	fmt.Println("player1")
	fmt.Println("player2")

	res2 := sum(10, 20)
	fmt.Printf("res2: %v\n", res2)

}
