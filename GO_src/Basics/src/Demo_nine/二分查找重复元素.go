package main

import "fmt"

func search2(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			high = mid - 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	//因为循环跳出条件为low = high + 1，当target比所有元素都大的话,此时high就是最后一个元素，low会越界，所以要判断一下
	//当target比所有元素都小的话，此时low不会越界，但要判断一下nums[low] 和 target 相不相等,如nums数组为[1],不判断的话会返回0
	if low >= len(nums) || nums[low] != target {
		return -1
	}
	//因为不管怎么样high一定是 mid-1
	//而循环跳出条件为low = high + 1，所以返回low

	return low
}

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	//判断是否越界
	if low >= len(nums) || nums[high] != target {
		return -1
	}
	//因为不管怎么样low一定是 mid+1
	//而循环跳出条件为low = high + 1，也就是high = low - 1，所以返回high

	return high
}

func main() {
	r := search2([]int{2, 4, 8, 10, 66, 66, 66, 68}, 66)
	fmt.Println(r)
}
