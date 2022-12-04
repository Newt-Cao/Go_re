package main

import (
	"fmt"
)

// 选择排序，将数组[30,27,100,12,8]升序
func Select_Sort(arr *[5]int) {

	// 外层循环是循环排序次数n-1次
	for i := 0; i < len(arr); i++ {
		// 声明假设最小值下标为数组的第一个值下标
		minIndex := i

		// 内层循环是循环数组内的数然后比较，从i+1开始是因为已经假定minIndex := i为最小值，和自己比较没意义
		for j := i + 1; j < len(arr); j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}

		// 找到最小值后交换，minIndex == i即最小值等于自己本身，不需要交换，交换无意义
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
}

func main() {
	arr := [5]int{30, 27, 100, 12, 8}
	Select_Sort(&arr)
	fmt.Println(arr) // [8 12 27 30 100]
}
