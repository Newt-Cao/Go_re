package main

/* 切片的截取与复制 */

import "fmt"

func main() {
	// 截取切片，[0,2) , 截取索引从 0 ~ 2 的元素，不包含 2 。
	s := []int{1, 2, 3}
	fmt.Println(len(s), cap(s), s) // 3 3 [1 2 3]
	s1 := s[0:2]
	fmt.Println(len(s1), cap(s1), s1) // 2 3 [1 2]

	// 截取数组
	a := [3]int{1, 2, 3}
	fmt.Println(len(a), cap(a), a) // 3 3 [1 2 3]
	a1 := a[0:2]
	fmt.Println(len(a1), cap(a1), a1) // 2 3 [1 2]

	fmt.Println(s1) //[1,2]

	fmt.Println("------------------")
	// 指向同一地址,同时修改
	s1[0] = 100

	fmt.Printf("%p\n", &s[0])
	fmt.Printf("%p\n", &s1[0])

	// 拷贝(拷贝底层数组的值)，不同地址。
	s2 := make([]int, 3) //[0,0,0]
	copy(s2, s)          //拷贝s的值到s2中

	fmt.Println(s2)
}
