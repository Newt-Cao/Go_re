package main

import "fmt"

func binarySearch(array *[5]int, userFind int) (int, bool) {
	// 声明中间值、首下标、尾下标
	var (
		middleIndex int
		firstIndex  int = 0
		trailIndex  int = len(array) - 1
	)
	// 左边的元素不能大于右边的元素
	for firstIndex <= trailIndex {
		// 获取中间元素下标，firstIndex和trailIndex的值不断改变，middleIndex的值也不断改变，缩小范围。
		middleIndex = (firstIndex + trailIndex) / 2
		// 当查询值大于中间值时，查询范围在中间值右边。firstIndex = middleIndex + 1 ，右边的第一个元素。
		// 当查询值大于中间值时，查询范围在中间值右边。trailIndex = middleIndex - 1 ，左边的最后一个元素。
		if userFind > array[middleIndex] {
			firstIndex = middleIndex + 1
		} else if userFind < array[middleIndex] {
			trailIndex = middleIndex - 1
		} else {
			return middleIndex, false
			// 查询到时，返回一个bool。
		}
	}
	return -1, true
}

// 递归
func BinarySearch(array *[5]int, userFind, leftIndex, rightIndex int) {
	// 结束条件
	if leftIndex > rightIndex {
		fmt.Println("找不到指定元素！")
		return
	}
	// 中间元素下标
	MiddleIndex := (leftIndex + rightIndex) / 2
	// 中间元素和查找值的判断
	if array[MiddleIndex] > userFind {
		BinarySearch(array, userFind, leftIndex, MiddleIndex-1)
	} else if array[MiddleIndex] < userFind {
		BinarySearch(array, userFind, MiddleIndex+1, rightIndex)
	} else {
		fmt.Println("找到了，该元素下标为:", MiddleIndex)
	}
}

func main() {
	// 声明一个数组
	array := [5]int{2, 42, 50, 51, 122}
	if index, ok := binarySearch(&array, 50); ok {
		fmt.Println("找不到指定元素！")
	} else {
		fmt.Printf("元素 %v 的下标是：%v\n", array[index], index)
	}
	// 递归
	array2 := [5]int{1, 2, 3, 5, 6}
	BinarySearch(&array2, 5, 0, len(array2)-1)
}
