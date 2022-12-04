package main

import "fmt"

func main() {
	/*
		练习：
			要求：
				1.传入一个 n 整型。
				2.输出一个斐波那契数列的切片。
				3.0 --> 1 , 1 --> 1
	*/
	func(n int) {
		// 声明一个 slice
		slice := make([]uint64, n)
		// 初始化第一个和第二个的元素值
		slice[0] = 1
		slice[1] = 1
		// 计算后面的元素
		for i := 2; i < n; i++ {
			slice[i] = slice[i-1] + slice[i-2]
		}
		// 打印
		fmt.Printf("1 - %v 的斐波那契数列是：%v,type = %T\n", n, slice, slice)
	}(20)

	func(n int) {
		slice := []uint64{}
		slice = append(slice, 1)
		slice = append(slice, 1)
		for i := 2; i < n; i++ {
			slice = append(slice, slice[i-1]+slice[i-2])
		}
		fmt.Printf("1 - %v 的斐波那契数列是：%v,type = %T\n", n, slice, slice)
	}(20)

}
