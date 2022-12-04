package main

import "fmt"

func test() {

	// 错误处理：defer + recover()
	defer func() {
		err := recover() //捕获异常
		if err != nil {
			fmt.Println("发现错误！", err)
		}
	}()

	// 0不可以作为分母
	var (
		num  = 10
		num2 = 0
		res  = num / num2
	)
	fmt.Printf("res: %v\n", res)
}

func main() {
	test()
	fmt.Println("捕获异常成功会打印这行字.")
}
