package main

import "fmt"

func main() {
	var array [8]int = [8]int{10,66,8,7,66,2,2,66}
	var array2 = array
	var minIndex, maxIndex []int

	for i := 0; i < len(array2)-1; i++ {
		for j := 0; j < len(array2)-1-i; j++ {
			if array2[j] > array2[j+1] {
				temp := array2[j]
				array2[j] = array2[j+1]
				array2[j+1] = temp
			}
		}
	}
	for i, _ := range array {
		if array[i] == array2[0] {
			minIndex = append(minIndex, i+1)
		} else if array[i] == array2[7] {
			maxIndex = append(maxIndex, i+1)
		}
	}
	fmt.Print("评分最低的评委是：")
	for _, v := range minIndex {
		fmt.Print(v," ")
	}
	fmt.Println()
	fmt.Print("评分最高的评委是：")
	for _, v := range maxIndex {
		fmt.Print(v," ")
	}
}
