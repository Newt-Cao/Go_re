package main

import "fmt"

func fibonacci(fib int) int {
	// 斐波那契数列：第一项和第二项为 1 ，第三项开始是前两项之和。(n>3,n属于自然数)
	/*
		分析：
			1.n == 1 或者 n == 2时返回1
			2.n >= 3时，返回 f(n) = f(n-1) + f(n-2)
			3.获取每一项对应的的斐波那契数列
	*/
	if fib == 1 || fib == 2 {
		return 1
	} else {
		return fibonacci(fib-1) + fibonacci(fib-2)
	}
}

func monkeyEat(day int) int {
	/*
		猴子吃桃案例：
			有一堆桃子，猴子第一天吃了其中的一半，并再多吃一个。
			以后每天猴子都吃其中的一半，然后再多吃一个。当到第十天时（还没吃），发现只有1个桃子了。问：最初共有多少个桃子？

		分析：
			1.f(10)=1
			2.f(9)=(f(9+1)+1)*2
			3.f(8)=(f(9)+1)*2=(f((f(9+1)+1)*2)+1)*2
			4.f(1)=f(2)+1*2
			5.f(n)=(f(n+1)+1)*2
	*/

	// 代码
	if day > 10 || day < 1 {
		fmt.Println("输入天数有误！")
		return 0
	}
	if day == 10 {
		return 1
	} else {
		return (monkeyEat(day+1) + 1) * 2
	}
}

func main() {

	// 斐波那契数列
	var num int
	fmt.Printf("请输入一个整数：\n")
	fmt.Scanln(&num)
	for i := 1; i <= num; i++ {
		fmt.Printf("第 %v\t项的斐波那契数为：%v\n", i, fibonacci(i))
	}

	// 猴子吃桃问题
	var peach int
	fmt.Printf("请输入(1-10)的天数：\n")
	fmt.Scanln(&peach)
	m := monkeyEat(peach)
	fmt.Printf("第 %v 天的桃子数: %v\n", peach, m)
}
