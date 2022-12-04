package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Rand() {
	// 设计一个 1-100 的随机数，并统计输出99时所用的次数

	/*
		1.使用死循环
		2.当循环至99时退出
		3.生成一个时间戳
		4.生成随机数
		5.声明一个变量计算生成99时所用的次数
		6.打印
	*/

	var numCount int
	for {
		// 时间戳，生成一个种子，Seed（x）x的值不同，对应的Intn（y）y值也不同，同理，如果x相同对应的y值也相同。
		rand.Seed(time.Now().UnixNano())

		// 随机数,Intn([0,n))
		r := rand.Intn(100) + 1
		fmt.Printf("r=%v\n", r)

		// 计算次数
		numCount++

		// 判断为99时结束
		if r == 99 {
			break
		}
	}
	// 打印
	fmt.Printf("生成99时所用的次数: %v次", numCount)
}

func PassCount(money float64) {
	/*
		某人有 100,000 元，每经过一次路口需要收费，规则如下：
			1.当现金大于 50000 时，每次交本金的 5%;
			2.当现金小于等于 50000 时，每次交 1000;

		（1）计算该人可以经过几次路口，使用 for - break 完成。

			分析：
				1.变量 money := 100000
				2.当 money > 50000 时 ，缴纳 --> money -= money * 0.05
				3.当 money <= 50000 时 ，缴纳 --> money -= 1000
				4.当 money < 1000 时 ， 无法缴纳和通过
				5.变量统计经过路口 passCount++
	*/

	/*
		var passCount int

		for {
			if money > 50000 {
				money -= money * 0.05
				passCount++
				fmt.Printf("money=%.3f\t\tpassCount=%v\n", money, passCount)
				continue
			}

			money -= 1000
			passCount++

			fmt.Printf("money=%.3f\t\tpassCount=%v\n", money, passCount)

			if money < 1000 {
				break
			}

		}
		fmt.Printf("经过了 %v 个路口,还剩 %.4f 元", passCount, money)
	*/

	// 优化代码
	var passCount int

	for {

		if money > 50000 {
			money -= money * 0.05
		} else {
			money -= 1000
		}

		fmt.Printf("money=%.3f\t\tpassCount=%v\n", money, passCount)

		if money < 1000 {
			break
		}
		passCount++

	}
	fmt.Printf("经过了 %v 个路口,还剩 %.4f 元", passCount, money)

}

func main() {
	//Rand()
	PassCount(100000)
}
