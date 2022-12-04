package main

import (
	"fmt"
	"time"
)

// 会一直输出1，主程序不会结束
func main() {
	bufChan := make(chan int)

	go func() {
		for {
			bufChan <- 1
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			fmt.Println(<-bufChan)
		}
	}()

	select {}
}
