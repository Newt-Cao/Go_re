package main

import "fmt"

func main() {
	type Point struct {
		x, y int
	}
	type Rectangle struct {
		left, right Point
	}

	R1 := Rectangle{Point{1, 2}, Point{3, 4}}
	fmt.Printf("R1.left.x的地址：%p\n", &R1.left.x)
	fmt.Printf("R1.left.x的地址：%p\n", &R1.left.y)
	fmt.Printf("R1.right.x的地址：%p\n", &R1.right.x)
	fmt.Printf("R1.right.x的地址：%p\n", &R1.right.y)
	/*
		int64 依次 +8
		R1.left.x的地址：0xc00000e200
		R1.left.x的地址：0xc00000e208
		R1.right.x的地址：0xc00000e210
		R1.right.x的地址：0xc00000e218
	*/

	type square struct {
		left, right *Point
	}
	R2 := square{&Point{1, 2}, &Point{3, 4}}
	fmt.Printf("R2.left的地址：%p\n", &R2.left)
	fmt.Printf("R2.right的地址：%p\n", &R2.right)
	fmt.Printf("R2.left储存的地址：%p\n", R2.left)
	fmt.Printf("R2.right储存的地址：%p\n", R2.right)
	/*
		int64 同样是依次 +8
		R2.left的地址：0xc000068230
		R2.right的地址：0xc000068238
		R2.left储存的地址：0xc0000180b0
		R2.right储存的地址：0xc0000180c0
	*/
}
