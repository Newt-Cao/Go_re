package main

import "fmt"

func main() {

	// 一、new(type)*type：由系统自动分配初始化的内存地址和值，默认初始化值为 0 。

	num := new(int)
	fmt.Printf("%T,%v,%v,%v\n", num, num, &num, *num) //*int,0xc000018088,0xc000006028,0

	// *type：不是自动分配，需要手动分配指定地址和对应的值，初始化地址为 nil 。

	var num2 *int
	fmt.Printf("%T,%v\n", num2, num2) //*int,<nil>
	// 手动分配地址
	var num3 int
	num2 = &num3
	fmt.Printf("%T,%v,%v,%v\n", num2, num2, &num2, *num2) //*int,0xc0000180a8,0xc000006038,0
}
