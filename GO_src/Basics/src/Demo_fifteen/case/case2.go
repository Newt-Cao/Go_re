/*
	案例二：
		1.启动一个协程，把1-2000的数放入到channel（numChan）中。
		2.启动八个协程，从numChan中取出数n，并计算1+...+n的值，然后存放到resChan中。
		3.等八个协程完成工作后，遍历resChan，显示结果如：res[1]=1...res[x]=y。
*/

package main

import (
	"fmt"
)

// 函数生成1-2000
func numCreate(numChan chan int) {
	for i := 1; i <= 2000; i++ {
		numChan <- i
	}
	// 存完关闭管道
	close(numChan)
}

// 取出数函数
func numGet(numChan chan int, resChan chan map[int]int, judgeChan chan int) {
	for {
		n, ok := <-numChan
		if !ok {
			break
		}
		res := 0
		for i := 1; i <= n; i++ {
			res += i
		}
		resChan <- map[int]int{n: res}
	}
	judgeChan <- -1
}

func main() {
	// 声明三个管道，存放、接收、判断
	numChan := make(chan int, 2000)
	resChan := make(chan map[int]int, 2000)
	judgeChan := make(chan int, 8)
	// 启动一个协程，生成1-2000的数
	go numCreate(numChan)
	// 启动八个协程
	for i := 0; i < 8; i++ {
		go numGet(numChan, resChan, judgeChan)
	}
	// 使主程序等待，当子程序结束后结束
	for i := 0; i < 8; i++ {
		<-judgeChan
	}
	// 子程序结束后关闭接收
	close(resChan)
	// 遍历
	for date := range resChan {
		for i, v := range date {
			fmt.Printf("res[%v] = %v\n", i, v)
		}
	}
	// 不再使用判断管道后关闭
	defer close(judgeChan)
}
