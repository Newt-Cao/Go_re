/* 切片的声明与空切片判断 */

package main

import "fmt"

// 声明切片 slice 的四个方法
func main() {
	// 只声明，不开辟空间元素。
	var Slice1 []int = []int{} // var Slice1 []int

	// 开辟三个空间地址。
	Slice1 = []int{1, 2, 3}     //有初始化值，且 len = cap
	var Slice2 = make([]int, 3) //默认初始化值0，且 len 可能等于 cap ，也可能不等于，也可以自己定义 cap 。

	// 推导写法，可简写。
	Slice3 := []int{1, 2, 3}
	Slice4 := make([]int, 3)

	//  %v 打印出详细信息
	fmt.Printf("len = %d,slice = %v,cap = %v\n", len(Slice1), Slice1, cap(Slice1)) //len = 0,slice = []
	fmt.Printf("len = %d,slice = %v,cap = %v\n", len(Slice2), Slice2, cap(Slice2)) //len = 3,slice = [1 2 3]
	fmt.Printf("len = %d,slice = %v,cap = %v\n", len(Slice3), Slice3, cap(Slice3)) //len = 3,slice = [0 0 0]
	fmt.Printf("len = %d,slice = %v,cap = %v\n", len(Slice4), Slice4, cap(Slice4)) //len = 3,slice = [0 0 0]

	// 如何判断一个切片是否为空切片
	// nil 是一个预先声明的标识符，表示指针、通道、函数、接口、映射或切片类型的零值。
	if Slice1 == nil {
		fmt.Println("Slice1是空切片")
	} else {
		fmt.Println("Slice1不是空切片")
	}
}
