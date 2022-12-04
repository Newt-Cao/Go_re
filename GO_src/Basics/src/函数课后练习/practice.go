package main

import (
	"fmt"
	"math/rand"
	"time"
)

func DatePrint() {
	/*
		练习一：
			循环打印输入月份的天数（使用continue实现），要有判断输入月份是否错误的语句。

		分析：
			1.定义3个变量接收年月日。
			2.Scanln接收用户输入信息。
			4.当执行完和输入错误时，使用continue跳到下一次循环重新输入。
			3.先简单判断是否存在逻辑错误，比如: 1 <= month <= 12 ，1 <= day <= 31，使用标签优化体验，返回到当前输入。
			4.判断闰年还是平年，闰年：(year % 4 == 0 && year % 100 != 0) || year % 400 == 0 。
			5.判断月份：
				2月：闰年 1<= day <= 29 ，平年 1<= day <= 28。
				1,3,5,7,8,10,12月：有31天。
				2,4,6,9,11月：有30天。
			6.年月日输入正确则Printf打印，然后跳到下一次循环。
			7.如输入错误则直接跳出当次循环，执行下次循环，并提示输入错误。
			8.使用死循环。
	*/
	var (
		year  int
		month int
		day   int
	)
	for {
		fmt.Print("请输入年：")
		fmt.Scanln(&year)
		fmt.Print("请输入月：")
	Mon:
		fmt.Scanln(&month)
		if month < 1 || month > 12 {
			fmt.Println("月份输入有误！请重新输入。")
			goto Mon
		}
	Day:
		fmt.Print("请输入日：")
		fmt.Scanln(&day)
		if day < 1 || day > 31 {
			fmt.Println("天数输入有误！请重新输入。")
			goto Day
		}

		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			switch month {
			case 2:
				if day > 29 {
					fmt.Println("天数输入有误！请重新输入。")
					continue
				} else {
					fmt.Printf("输入的日期是：%v年%v月%v日\n", year, month, day)
					continue
				}
			case 1, 3, 5, 7, 8, 10, 12:
				fmt.Printf("输入的日期是：%v年%v月%v日\n", year, month, day)
				continue
			default:
				if day > 30 {
					fmt.Println("天数输入有误！请重新输入。")
					continue
				} else {
					fmt.Printf("输入的日期是：%v年%v月%v日\n", year, month, day)
					continue
				}
			}
		} else {
			switch month {
			case 2:
				if day > 28 {
					fmt.Println("天数输入有误！请重新输入。")
					continue
				} else {
					fmt.Printf("输入的日期是：%v年%v月%v日\n", year, month, day)
					continue
				}
			case 1, 3, 5, 7, 8, 10, 12:
				fmt.Printf("输入的日期是：%v年%v月%v日\n", year, month, day)
				continue
			default:
				if day > 30 {
					fmt.Println("天数输入有误！请重新输入。")
					continue
				} else {
					fmt.Printf("输入的日期是：%v年%v月%v日\n", year, month, day)
					continue
				}
			}
		}
	}
}

func GuessNum() {
	/*
		练习二：
			随机生成一个 1-100 的整数，让用户进行猜数游戏，一共有十次机会，统计猜中所用的次数。
			1.猜 1 次猜中打印 “你真是个天才！”
			2.猜 2 - 3 次猜中打印 “你很聪明，快赶上我了！”
			3.猜 4 - 9 次猜中打印 “一般般！”
			4.猜 10 次猜中打印 “可算猜对了！”
			5.10次都没猜中最后打印 “Stupid！”

		分析：
			1.死循环，统计循环次数作为猜的次数，也作为判断剩余几次机会。
			2.rand生成随机数。
			3.用户输入一个数。
			4.判断是否存在逻辑错误。
			4.用户输入的数和生成随机数判断。
			5.猜大了和猜小了，做出适当提示。
			6.定义两个变量：生成的随机数、用户输入的数。
	*/
	var (
		randNum int
		userNum int
	)

	rand.Seed(time.Now().Unix())
	randNum = rand.Intn(100) + 1

guess:
	for i := 1; ; i++ {
		fmt.Print("请输入 1-100 内的一个整数：")
		fmt.Scanln(&userNum)

		if userNum < 1 || userNum > 100 {
			fmt.Println("输入错误！重新输入！")
			goto guess
		} else {
			if i < 10 {
				if userNum == randNum {
					switch i {
					case 1:
						fmt.Println("你真是个天才！")
						break
					case 2, 3:
						fmt.Println("你很聪明，快赶上我了！")
						break
					case 4, 5, 6, 7, 8, 9:
						fmt.Println("一般般！")
						break
					}
				} else if userNum < randNum {
					fmt.Printf("小了小了，往上猜！剩余 %v 次机会。\n", 10-i)
					continue
				} else if userNum > randNum {
					fmt.Printf("大了大了，往下猜！剩余 %v 次机会。\n", 10-i)
					continue
				}
			} else if i == 10 {
				if userNum == randNum {
					fmt.Println("可算猜对了！")
					break
				} else {
					fmt.Println("Stupid!")
					break
				}
			}
		}
	}
}

func PrimeNum() {
	/*
		练习三：
			输出100以内的所有素数，每行显示5个，并求和。

		分析：
			1.素数：大于1的自然数中，除了1和它本身以外不再有其他因数的自然数，因数小于或等于最大公因数。
			2.遍历：两个循环体判断 2-100 。
			3.判断：声明 i 为第一个循环体，声明 j 为第二个循环体，声明 Count 为得出素数时的次数。
				1) 当为素数时，有以下关系：
					1 * j == i --> j == i

				2) 当为积数时，有以下关系：
					i % j == 0 且 j < i

				3) 当得出素数时打印，积数时跳出当前循环体，继续下一个循环体。
				4) 当打印素数时再判断 Count ，如果 Count % 5 == 0 时，换行，否则不换行。
			4.Print打印时不换行。
			5.声明一个 sum 变量，sum += i 累加素数。
	*/
	var sum int
	var Count int

	for i := 2; i <= 100; i++ {

		for j := 2; j <= i; j++ {

			if j == i {
				fmt.Print(i)
				sum += i
				Count++

				if Count%5 == 0 {
					fmt.Print("\n")
				} else {
					fmt.Print("\t")
				}
			}

			if i%j == 0 && j < i {
				break
			}
		}
	}
	fmt.Println("100以内所有素数之和：", sum)
}

func FishingNet() {
	/*
		练习四：
			判断是打鱼还是晒网。
			1.从1990年1月1日起开始执行“三天打鱼两天晒网”
			2.判断在以后的某一天中，是打鱼还是晒网。

		分析：
			1.请用户输入年月日，分别用 userYear、userMonth、userDay 变量接收，判断基本逻辑：
				userYear >= 1990
				userMonth >= 1 && userMonth <= 12
				userDay >= 1 && userDay <= 31
				PS:如果错误跳回到当前标签重新输入。

			2.分别声明，闰年 leapYear 和平年 commonYear，从 1990 遍历到 userYear - 1：
				闰年：(i%4 == 0 && i%100 != 0) || i%400 == 0 --> leapYear++
				平年：else --> commonYear++
				得出：总的闰年和平年有多少个。
				PS:用于计算两个年份间相差多少天。

			3.判断 userYear :
				闰年：判断用户输入的二月份是否大于29，如果不是，从 1 遍历到 userMonth - 1，是则返回到day标签。
				平年：判断用户输入的二月份是否大于28，如果不是，从 1 遍历到 userMonth - 1，是则返回到day标签。
				使用switch分支，当 i 符合下面要求用 continue 终止本次循环，对应声明的变量++：
					1,3,5,7,8,10,12 --> 31天 leapMonth++/commonMonth++
					2               --> 闰年：29天 平年：28天 leapMonth2++/commonMonth2++
					4,6,9,11        --> 30天 leapMonth3++/commonMonth3++

			4.然后加上 userDay 计算总天数：
			Day = (leapYear * 366) + (commonYear * 365) + (leapMonth * 31) + (leapMonth2 * 29) + (leapMonth3 * 30) +(commonMonth * 31) + (commonMonth2 * 28) + (commonMonth3 * 30) + userDay

			5.最后分析发现，使用switch实现：
			Day % 5 == 1,2,3 --> 打鱼
			day % 5 == 0,4	 --> 晒网

		PS:遍历到 userYear 和 userMonth ，是因为最后一年是用户输入的，不满一年，按用户输入的计算，同理，最后一年的最后一个月也是用户输入的，不满12个月，也按用户输入的计算。

	*/
ALL:

	var (
		// 用户输入
		userYear  int
		userMonth int
		userDay   int
		// 闰年和平年
		leapYear   int
		commonYear int
		// 闰年的每个月
		leapMonth  int
		leapMonth2 int
		leapMonth3 int
		// 平年的每个月
		commonMonth  int
		commonMonth2 int
		commonMonth3 int
		// 最后计算 1990 到 userYear 的天数
		Day int
	)

	fmt.Print("请输入年：")
	fmt.Scanln(&userYear)
	if userYear < 1990 {
		fmt.Println("输入有误！请重新输入！")
		goto ALL
	}

mon:
	fmt.Print("请输入月份：")
	fmt.Scanln(&userMonth)
	if userMonth < 1 || userMonth > 12 {
		fmt.Println("输入有误！请重新输入！")
		goto mon
	}

day:
	fmt.Print("请输入日份：")
	fmt.Scanln(&userDay)
	if userDay < 1 || userDay > 31 {
		fmt.Println("输入有误！请重新输入！")
		goto day
	}

	for i := 1990; i <= userYear-1; i++ {
		if (i%100 != 0 && i%4 == 0) || i%400 == 0 {
			leapYear++
		} else {
			commonYear++
		}
	}

	if (userYear%100 != 0 && userYear%4 == 0) || userYear%400 == 0 {
		if userMonth == 2 {
			if userDay > 29 {
				fmt.Println("输入错误！请重新输入")
				goto day
			} else {
				for i := 1; i <= userMonth-1; i++ {
					switch i {
					case 1, 3, 5, 7, 8, 10, 12:
						leapMonth++
						continue
					case 2:
						leapMonth2++
						continue
					case 4, 6, 9, 11:
						leapMonth3++
						continue
					}
				}
			}
		}
	} else {
		if userMonth == 2 {
			if userDay > 28 {
				fmt.Println("输入错误！请重新输入")
				goto day
			} else {
				for i := 1; i <= userMonth-1; i++ {

					switch i {
					case 1, 3, 5, 7, 8, 10, 12:
						commonMonth++
						continue
					case 2:
						commonMonth2++
						continue
					case 4, 6, 9, 11:
						commonMonth3++
						continue
					}
				}
			}
		}
	}

	Day = (leapYear * 366) + (commonYear * 365) + (leapMonth * 31) + (leapMonth2 * 29) + (leapMonth3 * 30) +
		(commonMonth * 31) + (commonMonth2 * 28) + (commonMonth3 * 30) + userDay

	switch Day % 5 {
	case 1, 2, 3:
		fmt.Printf("从 1990.1.1 到 %v.%v.%v ，一共过了 %v 天，%v.%v.%v 这一天在打鱼!\n", userYear, userMonth, userDay, Day, userYear, userMonth, userDay)
		goto ALL
	case 4, 0:
		fmt.Printf("从 1990.1.1 到 %v.%v.%v ，一共过了 %v 天，%v.%v.%v 这一天在晒网\n!", userYear, userMonth, userDay, Day, userYear, userMonth, userDay)
		goto ALL
	}

}

func SmallCalc() {
	/*
		练习五：
			打印并实现其功能。

				----------小小计算器----------

				1.加法
				2.减法
				3.乘法
				4.除法
				0.退出

				----------小小计算器----------

				请输入：

		分析：
			1.使用``打印框架。
			2.判断用户输入。
			3.跳转对应功能实现计算。
	*/
	var num int8

all:

	for {

		fmt.Println(`
				  ----------小小计算器----------
				|				|
				|	  1.加法		|
				|	  2.减法		|
				|	  3.乘法		|
				|	  4.除法		|
				|	  0.退出		|
				|				|
				  ----------小小计算器----------`)

		fmt.Print("请输入功能（0 - 4）：")
		fmt.Scanln(&num)

		switch num {

		case 1:
			var a, b int
			fmt.Print("输入第一个数：")
			fmt.Scanln(&a)
			fmt.Print("输入第二个数：")
			fmt.Scanln(&b)
			fmt.Println(a, "+", b, "=", a+b)
			time.Sleep(1 * time.Second)

		case 2:
			var a, b int
			fmt.Print("输入第一个数：")
			fmt.Scanln(&a)
			fmt.Print("输入第二个数：")
			fmt.Scanln(&b)
			fmt.Println(a, "-", b, "=", a-b)
			time.Sleep(1 * time.Second)

		case 3:
			var a, b int
			fmt.Print("输入第一个数：")
			fmt.Scanln(&a)
			fmt.Print("输入第二个数：")
			fmt.Scanln(&b)
			fmt.Println(a, "*", b, "=", a*b)
			time.Sleep(1 * time.Second)

		case 4:
			var a, b float32
			fmt.Print("输入第一个数：")
			fmt.Scanln(&a)
			fmt.Print("输入第二个数：")
			fmt.Scanln(&b)
			fmt.Println(a, "/", b, "=", a/b)
			time.Sleep(1 * time.Second)

		case 0:
			var bool string
			fmt.Print("是否退出(t or f)？")
			fmt.Scanln(&bool)

			if bool == "t" {
				break all
			} else {
				continue all
			}

		default:
			fmt.Println("输入错误！请重新输入！")
			continue all

		}
	}
}

func English() {
	/*
		练习六：
			输出小写的 a-z 以及大写的 Z-A，要求使用for循环。

		分析：
			1.使用 ascii 码。
			2.使用 for 循环。
			3.分析 ascii 码发现 a-z 对应的是 97-122，A-Z 对应的是 65-91。
			4.Z-A可以使用 i-- 实现。
	*/
	for i := 97; i <= 122; i++ {
		fmt.Printf("%c ", i)
	}

	fmt.Println()

	for i := 90; i >= 65; i-- {
		fmt.Printf("%c ", i)
	}

}

func main() {

	//DatePrint()

	//GuessNum()

	//PrimeNum()

	//FishingNet()

	//SmallCalc()

	//English()

}
