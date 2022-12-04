/* 函数的定义 */

package main

import "fmt"

// 如何定义一个函数： func 函数名（形参 数据类型，形参 数据类型） 返回值数据类型/(多个返回值类型 X,X) {方法}

// 单个返回值：
func foo1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100

	return c
}

// 多个返回值，匿名：
func foo2(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	return 200, 300
}

// 多个返回值，有形参名称：
//       (a , b int) 形参相同时缩写（r1 , r2  int ） 返回值的数据类型相同时可缩写
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	// r1,r2 属于foo3的形参 ，相当于局部变量 ，golang中定义一个变量的初始值默认为 0 。
	// r1,r2 作用域是整个 foo3 ， 函数体{}的空间。
	fmt.Println("r1 = ", r1) // r1 = 0

	// 给对应的名称返回值变量赋值
	r1 = 1000
	r2 = 1000

	return
}

func main() {
	// 调用函数foo1
	c := foo1("abc", 555)
	fmt.Println("c = ", c)

	// 调用函数foo2
	result1, result2 := foo2("abcd", 400)
	fmt.Println("res1 = ", result1, "res2 = ", result2)

	// 调用函数foo3
	result3, result4 := foo3("abcde", 500)
	fmt.Println("res1 = ", result3, "res2 = ", result4)
}
