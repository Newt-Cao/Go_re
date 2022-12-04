package main

import "fmt"

func killOrbask() {
	var (
		year  int
		month int
		//当日天数
		nowday int
		//月总天数
		monthDay int = 1

		//月天数叠加
		monthdays int
		//年天数叠加
		yeardays  int
		yeardayss int
		//天数叠加
		days int
	)
	fmt.Println("输入年月日")
	fmt.Scanln(&year)
	fmt.Scanln(&month)
	fmt.Scanln(&nowday)

	//从1990遍历到year ， 日子叠加
	for i := 1990; i <= year-1; i++ {
		// 判断是不是润年
		if i%4 == 0 && i%100 != 0 || i%400 == 0 {
			yeardays++ //2月有29号
		} else {
			yeardayss++
		}
	}

	yeardays = (yeardays * 366) + (yeardayss * 365)

	switch {
	case year%4 == 0 && year%100 != 0 || year%400 == 0:
		switch month - 1 {
		case 1, 3, 5, 7, 8, 10, 12:
			monthdays += (month - monthDay) * 31
		case 4, 6, 9, 11:
			monthdays += (month - monthDay) * 30
		case 2:
			monthdays += (month - monthDay) * 29 //2月有29号
		}
	default:
		switch month - 1 {
		case 1, 3, 5, 7, 8, 10, 12:
			monthdays += (month - monthDay) * 31
		case 4, 6, 9, 11:
			monthdays += (month - monthDay) * 30
		case 2:
			monthdays += (month - monthDay) * 28 //2月有28号
		}
	}

	days = yeardays + nowday + monthdays

	fmt.Printf("从1990年1月1日到现在一共过了%v天\n", days)

	//判断晒网 打鱼 总天数%5 == [1,2,3]打鱼  [4,0]晒网
	switch days % 5 {
	case 1, 2, 3:
		fmt.Printf("今天是%v - %v - %v 应该打鱼", year, month, nowday)
	case 4, 0:
		fmt.Printf("今天是%v - %v - %v 应该晒网", year, month, nowday)
	}
}

func main() {
	killOrbask()
}
