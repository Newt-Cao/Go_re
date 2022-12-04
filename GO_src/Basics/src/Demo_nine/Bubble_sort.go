package main

import (
	"fmt"
)

/* 对数组 arr = [24,69,80,57,13] 进行冒泡排序。 */

func Bubble_sort(arr *[5]int) {
	// 第一版：基础代码。

	for j := 0; j < len(arr)-1; j++ {

		for i := 0; i < len(arr)-1-j; i++ {

			if arr[i] > arr[i+1] {
				temp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp
			} else {
				continue
			}
		}
	}

	// 打印排序后的数组
	fmt.Println(*arr)
}

func Bubble_sort2(arr *[10]int) {
	// 第二版：优化了第一版的无差别遍历，假设未到最后一轮时，数组已经有序，则不需要再进行后续轮次的排序（使用flag标识）。

	for j := 0; j < len(arr)-1; j++ {
		flag := true //当数组有序时，flag一直未true，说明不需要继续排序

		for i := 0; i < len(arr)-1-j; i++ {

			if arr[i] > arr[i+1] {
				temp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp
				flag = false //需要交换时说明数组还是无序，标识为flag，不退出
			}
		}

		if flag {
			break //不需要继续排序时退出
		}
	}
	fmt.Println(*arr)
}

func Bubble_sort3(arr *[10]int) {
	// 第三版：鸡尾酒排序，即先进行从左到右的排序，再进行从右到左的排序，这样能减少轮次数,缩小范围。使用鸡尾酒排序可解决一些特殊情况，排序的轮次数也能减少一半。

	// 只需n/2轮的比较，因为每轮里都会吧最大值移到队尾，最小值移到队首
	for loop := 1; loop < len(arr)/2; loop++ {
		sorted := false

		var j int
		// 先正向冒泡，把最大值移动到队尾
		for j = loop - 1; j < len(arr)-loop; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				sorted = true
			}
		}

		// 如果跑了一轮没有交换元素，说明已经排好序了
		if !sorted {
			break
		}

		// 再反向冒泡，把最小值移动到队首
		for ; j >= loop; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				sorted = true
			}
		}

		if !sorted {
			break
		}
	}
	fmt.Println(*arr)
}

func main() {

	// 定义数组
	arr := [5]int{24, 69, 80, 57, 13}
	arr2 := [10]int{4, 20, 1, 3, 6, 8, 5, 4, 650, 30}
	arr3 := [10]int{60, 5, 8, 20, 17, 6, 30, 5, 42, 10}
	//调用函数
	Bubble_sort(&arr)
	Bubble_sort2(&arr2)
	Bubble_sort3(&arr3)
	/*
		[13 24 57 69 80]
		[1 3 4 4 5 6 8 20 30 650]
		[5 5 6 8 10 17 20 30 42 60]
	*/

}
