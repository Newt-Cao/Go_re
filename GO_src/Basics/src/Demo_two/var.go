/*******************************************************
Author       : Mr.Chan
Date         : 2022-03-27 18:08:50
 LastEditors  : Mr.Chan
 LastEditTime : 2022-09-08 09:30:29
 FilePath     : /vivianchan.cn/Basics/src/Demo_two/var.go
*******************************************************/

/* 四种变量的声明 */

package main

import "fmt"

// 局部变量和全局变量用方法一、二、三，方法四只能声明函数体内局部变量。
var gA int = 100

func main() {

	// 方法一：声明一个变量，默认值是：0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a) // int

	// 方法二：声明一个变量，赋予一个初始值。
	var bb string = "abcd"
	fmt.Printf("bb = %s\ntype of bb = %T\n", bb, bb) // abcd string

	// 方法三：在初始化阶段可以省略数据类型，让其自动匹配数据类型。
	var c = 100
	fmt.Println("c = ", c)            // 100
	fmt.Printf("type of c = %T\n", c) // int

	// 方法四（最常用）：省去var关键字，用":"直接匹配，只能声明局部变量。
	d := 100
	fmt.Println("d = ", d)
	fmt.Printf("type of d = %T\n", d)

	// 方法五：使用两个小括号
	var e = (int)(10)
	fmt.Printf("%v %T\n", e, e) // 10 int
	var f = (*int)(&e)
	fmt.Printf("%v,%v,%T,%p\n", f, *f, f, &f) // 0xc00001a0d0,10,*int,0xc000006030

	// 全局变量
	fmt.Println("gA = ", gA)

	// 声明多个变量
	// 方法一：
	var dd, ee int = 100, 200
	fmt.Println("dd = ", dd, "ee = ", ee)
	var ff, gg = 100, "abcd"
	fmt.Println("ff = ", ff, "gg = ", gg)

	// 方法二：
	var (
		h int     = 100
		j float64 = 3.14
		k bool    = true
	)
	fmt.Println(h, "\n", j, "\n", k)
}
