package main

import (
	"fmt"
	"time"
)

func main() {
	// 定义一个管道
	channel := make(chan int, 5)

	// 写入数据
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Println("do something...")
			time.Sleep(5 * time.Second)
			channel <- i
		}
	}()

	// 利用select处理超时问题，因为当select某一个case接收到值后，会退出
	// 当协程未在指定时间完成任务，会跳到下一个case，等待2秒后主程序退出，打印超时
	select {
	// 5秒后才会接收第一个值
	case x := <-channel:
		fmt.Println("receive！", x)

	// 2秒后超时
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout!")
	}

	/*
		do something...
		Timeout!
	*/
}
