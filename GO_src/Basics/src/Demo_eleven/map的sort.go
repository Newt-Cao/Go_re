package main

import (
	"fmt"
	// sort包可以针对切片进行排序
	"sort"
)

func main() {
	/*
		map的排序：
			1.现将map的key放入到切片当中使用sort包排序；
			2.把排序好的key以key-value输出。
			3.sort包可以针对切片进行排序。
	*/

	// 声明
	mapTest := make(map[int]int)
	// 赋值
	mapTest[20] = 1
	mapTest[3] = 100
	mapTest[69] = 8
	mapTest[5] = 30
	mapTest[4] = 56
	// 遍历把key放到切片中
	mapSlice := []int{}
	for key, _ := range mapTest {
		mapSlice = append(mapSlice, key)
	}
	// 调用sort包
	sort.Ints(mapSlice)
	fmt.Println(mapSlice)
	// 排序完后使用key-value输出
	for _, value := range mapSlice {
		fmt.Printf("map[%v]=[%v]", value, mapTest[value])
		fmt.Println()
	}
	/*
		[3 4 5 20 69]
		map[3]=[100]
		map[4]=[56]
		map[5]=[30]
		map[20]=[1]
		map[69]=[8]
	*/
}
