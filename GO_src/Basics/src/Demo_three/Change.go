// 交换值

package main

import (
	"fmt"
)

func change(a, b int) {
	var temp int
	temp = a
	a = b
	b = temp
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
}

func change1(a, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
}

func main() {

	// 值交换 swap ，内存地址不变。
	n1 := 1
	n2 := 2
	fmt.Println(&n1, &n2)
	fmt.Printf("n1: %v\nn2: %v\n", n1, n2)

	n1, n2 = n2, n1
	fmt.Printf("n1: %v\nn2: %v\n", n1, n2)
	fmt.Println(&n1, &n2)

	change(n1, n2)

	// 指针传递

	var n3 *int
	n3 = &n1
	var n4 *int
	n4 = &n2
	fmt.Printf("n3: %v n4: %v\n", n3, n4)
	fmt.Printf("n3: %v n4: %v\n", *n3, *n4)

	n3, n4 = n4, n3
	fmt.Printf("n3: %v n4: %v\n", n3, n4)
	fmt.Printf("n3: %v n4: %v\n", *n3, *n4)

}
