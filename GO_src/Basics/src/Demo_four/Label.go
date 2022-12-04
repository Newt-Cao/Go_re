// 流程控制中的标签使用

package main

import (
	"fmt"
)

// 结束对应标签下的流程控制
func Mylaber() {
Label:
	for i := 0; i < 10; i++ {
		fmt.Printf("i:%d\n", i)
		if i >= 5 {
			break Label
		}
	}
}

// 打印1--100内的奇数
func Odd() {
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			continue
		} else {
			fmt.Printf("Odd_number:%d\n", i)
		}
	}
}

// 判断随机输入的数是负数还是正数，输入0退出
func Judge() {
	var positive int
	var negative int
	var num []int
	num = make([]int, 5)
	// var num1 int
	var judge string

Judge1:
	fmt.Println("请输入五个数：")
	fmt.Scanln(&num)
	for {
		if len(num) == 0 {
			fmt.Println("是否要退出程序？y/n")
			fmt.Scanln(&judge)
			if judge == "y" {
				fmt.Println("程序已退出！")
				break
			} else {
				goto Judge1
			}
		}
		if len(num) > 0 {
			for j := 0; j == len(num); j++ {
				if j > 0 {
					positive++
					continue
				} else if j < 0 {
					negative++
				}
			}
		}
	}

	fmt.Printf("正数：%d，负数：%d", positive, negative)
}

func main() {
	//Mylaber()
	//Odd()
	Judge()
}
