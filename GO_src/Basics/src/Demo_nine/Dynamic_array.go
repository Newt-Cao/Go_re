/* 动态数组 */

package main

import "fmt"

// 传递动态数组,引用传递，指针传递，所有形参都是 []int
func printArray(myArray1 []int) {
	// 遍历动态数组，下划线 _ 表示匿名的变量，接受返回的 index 但是不返回
	for _, value := range myArray1 {
		fmt.Println("value = ", value)
	}
	// 可修改对应地址的值
	myArray1[0] = 100
}

func main() {
	// 定义一个动态数组,切片 slice
	myArray := []int{1, 2, 3, 4}

	// 打印动态数据类型
	fmt.Printf("type of maArray = %T\n", myArray) //[]int

	// 调用
	printArray(myArray)
	// 修改
	fmt.Println("------------------------------")
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
}
