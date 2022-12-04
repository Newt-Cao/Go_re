package main

import "fmt"

func main() {

	// 符合条件完成自动 break ,只会执行一个case
	Grade()
	Day()

	// 穿透，一直往下执行，可使用break 终止
	Fallthrough()
}

func Grade() {

	/*
		1.switch和case后面可以连接变量、常量、有返回值的函数、表达式
		2.switch和case的数据类型必须相同
		3.case后面可以跟多个表达式，使用 "," 隔开
		4.case后面如果是常量或者字面量，要求不可以重复
		5.default不是必须的
		6.switch后面可以不接任何表达式，当做if-else来用
		7.switch后面可以声明一个变量，使用分号结束（不推荐）
		8.case后增加fallthrough进行穿透，匹配符合条件时不结束，继续下一个，直到没有fallthrough为止
		9.数据类型判断
	*/

	var x interface{}
	var y = 10.0
	x = y

	switch i := x.(type) {
	case nil:
		fmt.Printf("%T", i)
	case bool:
		fmt.Printf("%T", i)
	case float64:
		fmt.Printf("%T", i)
	case string:
		fmt.Printf("%T", i)
	case func(int) float64:
		fmt.Printf("%T", i)
	case int:
		fmt.Printf("%T", i)
	}

	
	var grade string
	fmt.Println("请输入等级：")
	fmt.Scanln(&grade)

	// 判断单个条件
	switch grade {
	case "A":
		fmt.Println("优秀")
	case "B":
		fmt.Println("良好")
	case "C":
		fmt.Println("合格")
	default:
		fmt.Println("错误")
	}
}

func Day() {
	// 判断多个条件
	var day int
	fmt.Println("输入星期几：")
	fmt.Scan(&day)

	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("工作日")
	case 6, 7:
		fmt.Println("休息日")
	default:
		fmt.Println("错误")
	}
}

func Fallthrough() {
	num := 100

	switch num {
	case 100:
		fmt.Println("100")
		fallthrough

	case 200:
		fmt.Println("200")
	case 300:
		fmt.Println("300")
	}
} // 100 200
