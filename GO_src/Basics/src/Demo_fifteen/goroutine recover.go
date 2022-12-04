package main

import "fmt"

// 错误处理机制函数，在每个函数结束后调用，即使一个协程函数崩溃不会导致主程序崩溃
func errorRecover() {
	// 捕获错误并打印
	if err := recover(); err != nil {
		fmt.Println("err = ", err)
	}
}

// 函数一
func test01(quit chan int) {
	defer errorRecover()
	num := 0
	for {
		num++
		fmt.Println(num)
		if num >= 5 {
			break
		}
	}
	quit <- 1
}

// 函数二，内有错误时被errorRecover捕获，不会导致主程序崩溃
func test02(quit chan int) {
	defer errorRecover()
	var slice []int
	slice[0] = 1
	quit <- 1
}

func main() {
	quit := make(chan int)
	// 协程调用两个函数
	go test01(quit)
	go test02(quit)
	// 阻塞主程序
	select {
	case <-quit:
		fmt.Println("Main Over！")
	}
}

/*
err =  runtime error: index out of range [0] with length 0
1
2
3
4
5
Main Over！
*/
