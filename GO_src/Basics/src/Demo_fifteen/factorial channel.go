package main

import (
	"fmt"
)

var (
	// 用于接收计算的值并保存为map类型
	factorialChan = make(chan map[int]uint64, 20)
	// 用于主程序等待
	exitChan = make(chan int, 20)
)

func factorial(n int, factorialChan chan<- map[int]uint64, exitChan chan int) {
	// 阶乘计算
	res := 1
	i := 1
	for ; i <= n; i++ {
		res *= i
	}

	// 将结果放入到map中，转换为uint64防止溢出
	factorialChan <- map[int]uint64{i - 1: uint64(res)}
	// 每调用一个协程结束，发送一次
	exitChan <- 1
}

func main() {
	// 并发或并行，计算1-10阶乘
	for i := 1; i <= 20; i++ {
		go factorial(i, factorialChan, exitChan)
	}
	// 接收上面协程完成后发送的信号
	for i := 0; i < 20; i++ {
		<-exitChan
	}
	// 当所有协程结束后，关闭
	close(factorialChan)
	// 遍历打印
	for factMap := range factorialChan {
		for i, v := range factMap {
			fmt.Printf("%v！ = %v\n", i, v)
		}
	}
}
