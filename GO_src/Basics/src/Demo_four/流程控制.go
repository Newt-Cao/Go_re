package main

import (
	"fmt"
)

func Age() {
	var age int

	for i := 0; i <= 3; i++ {

		fmt.Println("输入你的年龄：")

		// 获取用户输入信息
		if fmt.Scanln(&age); age >= 22 {
			fmt.Printf("%d岁你已经可以结婚\n", age)

		} else {
			fmt.Printf("%d岁还没到法定结婚年龄\n", age)
		}
	}
}

func Score() {
	// 多重分支,嵌套
	var score int

	for {
		fmt.Println("请输入分数：")

		if fmt.Scan(&score); score >= 90 {
			fmt.Println("评分：A")

		} else if 90 > score && score >= 80 {
			fmt.Println("评分：B")

		} else {
			fmt.Println("评分：C")
		}

	}
}

func main() {
	Age()
	Score()
}
