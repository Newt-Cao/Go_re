package main

import (
	"fmt"
	"sync"
	"time"
)

// 声明全局map用于接收
var (
	factMap = make(map[int]uint64)
	// 声明全局变量互斥锁，Mutex是互斥的意思
	Lock sync.Mutex
)

func factorial(n int) {
	// 阶乘计算
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	// 加锁
	Lock.Lock()
	// 将结果放入到map中，转换为uint64防止溢出
	factMap[n] = uint64(res)
	// 解锁
	Lock.Unlock()
}

func main() {
	// 并发或并行，计算1-200阶乘
	for i := 1; i <= 10; i++ {
		go factorial(i)
	}

	// 主程序休眠，等待子程序执行
	time.Sleep(5 * time.Second)

	// 因为主程序和子程序之间无通信，所以也需要加锁，等子程序运行完后解锁
	Lock.Lock()
	// 打印
	for i, v := range factMap {
		fmt.Printf("%v! = %v\n", i, v)
	}
	Lock.Unlock()
}
