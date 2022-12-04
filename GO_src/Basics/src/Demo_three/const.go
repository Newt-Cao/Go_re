/* 常量的声明和关键字 iota 使用 */

package main

import "fmt"

// const 可定义枚举类型。
// 例一（普通）：
const (
	// 可以添加关键字 iota ，第一行iota的默认值是 0 ，往下每行的iota会累加 1 。（只可以配合const使用）
	BEIJING   = iota //iota = 0  BEIJING = 10*iota
	SHANGHAI         //iota = 1
	GUANGZHOU        //iota = 2
)

//例二（公式）：
const (
	a, b = iota + 1, iota + 2 //iota = 0, a = iota + 1,b = iota + 2,a = 1,b = 2
	c, d                      //iota = 1, c = iota + 1,b = iota + 2,c = 2,d = 3

	e, f = iota * 2, iota * 3 //iota = 2, e = iota * 2,f = iota * 3,e = 4,f = 6
	g, h                      //iota = 3, g = iota * 2,h = iota * 3,g = 6,h = 9
)

const (
	i, j = iota, iota // 0 0
	k, l = iota, iota // 1 1
)

func main() {
	// 常量（只读属性）不允许修改！
	const length int = 10
	fmt.Println("length = ", length)

	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("GUANGZHOU = ", GUANGZHOU)

	fmt.Println(a, b, c, d, e, f, g, h)
}
