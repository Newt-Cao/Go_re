package main

import "fmt"

// 形参的传递可以多个,args...但必须放在独立形参的后面，除非单独使用
func A(n1, n2 float32, args ...int) float32 {
	fmt.Printf("n1: %T\n", n1)
	sum := n1 + n2
	for i := 0; i < len(args); i++ {
		sum += float32(args[i])
	}
	return sum
}

func B(function func() int) int {
	return 0
}

func main() {
	a := A
	res := a(1, 2, 4, 5, 6, 8)
	fmt.Printf("%T,%T,%v\n", a, A, res) // func() int,func() int,0
}
