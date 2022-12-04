package main

import "fmt"

func star(i int) {

	// 循环层数用i决定
	for a := 1; a <= i; a++ {

		// 空格打印数用i-a决定
		for c := 1; c <= i-a; c++ {
			fmt.Print(" ")
		}

		// 打印内容用2*a-1决定
		for b := 1; b <= 2*a-1; b++ {

			// 镂空打印 b==1、b==2*a-1、a==i 分别对应的是 打印第一个、打印最后一个、打印最后一层
			if b == 1 || b == 2*a-1 || a == i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}

		}
		// 换行
		fmt.Println()
	}
}

func main() {
	star(3)
}
