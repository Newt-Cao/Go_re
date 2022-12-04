package main

import (
	"fmt"
)

func main() {
	// 声明一个缓冲为3的管道
	testChannel := make(chan int, 3)
	// 传入数据
	testChannel <- 10
	// 关闭管道
	close(testChannel)
	// 从管道中读取数据，判断管道是否关闭
	num := <-testChannel
	num2, ok := <-testChannel
	// ok判断读取数据是否成功，如果读取管道最后的数据后，无法再读取数据，则返回false管道已关闭，反之则返回true
	if ok {
		fmt.Println("管道未关闭！")
	} else {
		fmt.Println("管道已关闭！")
	}
	// 打印
	fmt.Println(num, num2, ok)
	/*
		管道已关闭！
		10 0 false
	*/
}
