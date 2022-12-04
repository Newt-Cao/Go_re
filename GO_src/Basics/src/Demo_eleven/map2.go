/* map 的一些操作 */

package main

import "fmt"

// 传参,是一个引用传递 ，指向同一个地址
func printMap(cityMap map[string]string) {
	for key, value := range cityMap {
		fmt.Println("key = ", key, "value = ", value)
	}
}

// 传参，修改
func ChangeValue(cityMap map[string]string) {
	cityMap["England"] = "London"

	for key, value := range cityMap {
		fmt.Println("key = ", key, "value = ", value)
	}
}

func main() {
	// 定义
	cityMap := make(map[string]string)

	// 添加
	cityMap["China"] = "Beijing"
	cityMap["USA"] = "Newyork"
	cityMap["Japan"] = "Tokyo"

	// 遍历（针对数组、切片、字符串返回索引和值，map返回键和值，channel返回通道内的值）
	for key, value := range cityMap {
		fmt.Println("key = ", key, "value = ", value)
	}

	// 删除
	delete(cityMap, "Japan")

	// 修改
	cityMap["USA"] = "DC"

	fmt.Println("----------------------------")

	for key, value := range cityMap {
		fmt.Println("key = ", key, "value = ", value)
	}

	// 函数调用
	printMap(cityMap)    //遍历
	ChangeValue(cityMap) //修改
}
