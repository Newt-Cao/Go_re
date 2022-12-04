/* 协程 goroutine */

package main

import (
	"fmt"
	"time"
)

// 从 goroutine
func GoroutineTest() {
	// 死循环
	for i := 0; ; i++ {

		fmt.Printf("Son run i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

// 主 goroutine
func main() {
	go GoroutineTest()

	for i := 0; ; i++ {

		fmt.Printf("Dad run i = %d\n", i)
		time.Sleep(1 * time.Second)

	}
}

// 当主线程结束后子线程一起结束
