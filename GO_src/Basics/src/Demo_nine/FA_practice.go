package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		练习：随机生成五个数，并反转打印。

		分析：
			1.声明一个[5]int数组。
			2.生成种子和随机数。
			3.遍历，把随机数赋值给数组的每一个元素。
			4.交换的次数为：len/2，第一个和最后一个交换，第二个和倒数第二个交换。
	*/

	var array [5]int

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(100)
	}

	fmt.Println("交换前：", array)

	/*
		方法一：
			for i := 0; i < len(array)/2; i++ {

			temp := array[len(array)-1-i]

			array[len(array)-1-i] = array[i]

			array[i] = temp

			}

		方法二：
			temp0, temp1 := array[0], array[1]
			temp3, temp4 := array[3], array[4]

			array[0], array[1] = temp4, temp3
			array[3], array[4] = temp1, temp0
	*/

	fmt.Println("交换后：", array)

}
