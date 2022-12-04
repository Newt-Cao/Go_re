package main

import "fmt"

// 插入排序，讲[40,23,19,-4,100]升序

func Insert_Sort(arr *[5]int) {
	// 外循环循环次数是n-1，即有5个元素就要比较至少4次才可以
	for i := 1; i < len(arr); i++ {
		// 声明一个变量接收无序表的第一个值
		disorderValue := arr[i]

		// 插入的下标是在无序表第一个元素的前一个下标的元素左右找合适位置（有序表的最后一个元素下标）
		insertIndex := i - 1

		// 内层循环，数组内之间的数比较，如果是升序就找大于有序表最后一个元素下标，降序找小于有序表最后一个元素下标，数组下标必须insertIndex >= 0
		for insertIndex >= 0 && arr[insertIndex] > disorderValue {
			// 当找到合适位置后插入
			arr[insertIndex+1] = arr[insertIndex]
			// 遍历完一个元素，继续往前面遍历，直到-1就是数组最前端无数可比较退出，所以这里最后insertIndex每次退出for循环时都是-1
			insertIndex--
		}

		// 还要判断insertIndex本身是否等于disorderVal，如果相等没必要交换，已经是排好了，又因为insertIndex==-1，所以要加1
		if insertIndex+1 != i {
			arr[insertIndex+1] = disorderValue
		}
	}
}

func main() {
	arr := [5]int{40, 23, 19, -4, 100}
	Insert_Sort(&arr)
	fmt.Println(arr) // [-4 19 23 40 100]
}
