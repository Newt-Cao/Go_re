package main

import "fmt"

// 快速排序
func Quick_sort(slice []int) []int {
	//判断是否为不需要排序的数组
	if len(slice) < 2 {
		return slice
	}
	// 选择一个基准数
	pivot := slice[0]
	// 声明两个子数组
	var (
		left  []int
		right []int
	)
	// 遍历该数组并把小于和大于基准数的元素分类添加
	for _, value := range slice[1:] {
		if value <= pivot {
			left = append(left, value)
		} else {
			right = append(right, value)
		}
	}
	// 递归调用
	return append(Quick_sort(left), append([]int{pivot}, Quick_sort(right)...)...)
}

func Quick_sort2(head, tail int, array *[6]int) {
	// 基准数
	pivot := array[(head+tail)/2]
	l := head
	r := tail

	for l < r {
		// 找到大于pivot值，然后继续向右移动
		for array[l] < pivot {
			l++
		}
		// 找到小于pivot值，然后继续向左移动
		for array[r] > pivot {
			r--
		}
		if l >= r {
			break
		}
		// 把找到的下标值进行交换，大于的放到pivot右边，小于的放到pivot左边
		array[l], array[r] = array[r], array[l]
		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}
	}
	if l == r {
		l++
		r--
	}
	if head < r {
		Quick_sort2(head, r, array)
	}
	if tail > l {
		Quick_sort2(l, tail, array)
	}
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	i, j := left, right
	// i从左边开始遍历，j从右边开始遍历。
	p := nums[i]
	// 设置基准数
	for i < j {
		for i < j && nums[i] < p {
			// 寻找一个比基准数大的数
			i++
		}
		for i < j && nums[j] > p {
			// 寻找一个比基准数小的数
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
		// 替换这两个数
	}
	nums[j] = p
	// 如果i和j相遇，就将基准数和j下标的值进行替换

	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)
}

func main() {
	slice := []int{52, 6, 9, 1, 56, 6218, 48, 99}
	slice = Quick_sort(slice)
	fmt.Println(slice) //[1 6 9 48 52 56 99 6218]

	quickSort(slice, 0, len(slice)-1)
	fmt.Println(slice) //[1 6 9 48 52 56 99 6218]

	arr := [6]int{52, 6, 1, 1, 56, 6218}
	Quick_sort2(0, len(arr)-1, &arr)
	fmt.Println(arr) //[1 6 9 48 52 56]
}
