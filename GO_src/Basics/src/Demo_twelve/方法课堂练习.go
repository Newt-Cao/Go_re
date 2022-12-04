package main

import "fmt"

/*
练习一：
	编写结构体 MethodUils ，给结构体绑定给一个方法，方法不需要参数，方法执行打印一个 8*10 的矩形，最后在 main 方法中执行该方法。
*/
type MethodUtils struct {
	Array [3][3]int
}

func (met MethodUtils) graph() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 10; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
练习二：
	绑定一个方法，提供两个参数 m 和 n ，方法执行一个 m*n 的矩形。
*/
func (met MethodUtils) graphmn(m, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
练习三：
	绑定一个方法计算矩形的面积，接收两个参数长 len 和宽 width，将面积作为方法的返回值。在 main 中调用，接收返回值并打印。
*/
func (met MethodUtils) area(len, width float64) float64 {
	return len * width
}

/*
练习四：
	绑定一个方法，判断一个数是奇数还是偶数。
*/
func (met MethodUtils) JudgeNum(x int) {
	if x%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
}

/*
练习五：
	根据行、列、字符打印对应行数和列数的字符。
*/
func (met MethodUtils) charPrint(m, n int, c string) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(c)
		}
		fmt.Println()
	}
}

/*
练习六：
	定义一个计算机结构体，实现加减乘除四个功能。
*/
type Calcuator struct {
	oneNum   float64
	twoNum   float64
	operator string
}

func (calc *Calcuator) calcMethod() float64 {
	result := 0.0
	switch calc.operator {
	case "+":
		result = calc.oneNum + calc.twoNum
	case "-":
		result = calc.oneNum - calc.twoNum
	case "*":
		result = calc.oneNum * calc.twoNum
	case "/":
		result = calc.oneNum / calc.twoNum
	default:
		fmt.Println("输入有误！")
		return -1
	}
	return result
}

/*
练习七：
	用结构体 Merhodutils 绑定一个方法，从键盘接收整数（1-9），打印对应的乘法表。
*/
func (met MethodUtils) multiplicationTable() {
	var num int
	fmt.Print("输入整数（1-9）：")
	fmt.Scanln(&num)
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v\t", j, i, i*j)
		}
		fmt.Println()
	}
}

/*
练习八：
	绑定一个方法，使给定的一个二维数组（3 X 3）转置。
	1 2 3     1 4 7
	4 5 6 --> 2 5 8
	7 8 9     3 6 9
*/
func (met *MethodUtils) twoArray() {
	for _, v := range met.Array {
		for _, v1 := range v {
			fmt.Print(v1, " ")
		}
		fmt.Println()
	}

	var arr2 [3][3]int
	for i, v := range arr2 {
		for i1, _ := range v {
			arr2[i][i1] = met.Array[i1][i]
		}
	}
	fmt.Println()

	met.Array = arr2

	for _, v := range met.Array {
		for _, v1 := range v {
			fmt.Print(v1, " ")
		}
		fmt.Println()
	}
}

/*
	简单案例：
		门票售价，大于18的收费20，其他情况免费，声明结构体，根据年龄判断门票费用。
*/
type Visitor struct {
	Name string
	Age  int
}

func (visit *Visitor) price() {
	if visit.Age >= 99 || visit.Age <= 8 {
		fmt.Println("不符合该年龄群体游玩！")
		return
	}
	if visit.Age > 18 {
		fmt.Println("年龄大于18，需要交费20元。")
	} else {
		fmt.Println("您的年龄符合免费优惠。")
	}
}

func main() {
	// 练习一
	Met := MethodUtils{}
	Met.graph()

	fmt.Println()

	// 练习二
	Met.graphmn(8, 10)

	fmt.Println()

	// 练习三
	areaResult := Met.area(10.1, 8)
	fmt.Printf("矩形的面积为：%v\n", areaResult)

	fmt.Println()

	// 练习四
	Met.JudgeNum(11)

	fmt.Println()

	// 练习五
	Met.charPrint(10, 8, "&")

	fmt.Println()

	// 练习六
	calcResult := Calcuator{10, 20.1, "+"}
	fmt.Printf("计算结果：%.1f", calcResult.calcMethod())

	fmt.Println()

	// 练习七
	Met.multiplicationTable()

	fmt.Println()

	// 练习八
	Met.Array = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	(&Met).twoArray()

	// 简单案例
	visitors := Visitor{}
	for {
		fmt.Println("请输入姓名：")
		fmt.Scanln(&visitors.Name)
		fmt.Println("请输入年龄：")
		fmt.Scanln(&visitors.Age)
		visitors.price()
	}
}
