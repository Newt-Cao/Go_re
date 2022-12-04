package main

import "fmt"

// 声明管道
var (
	// 双向
	testChan = make(chan int, 10)
	exitChan = make(chan int, 2)
	exit     int
)

// 只写管道
func send(testChan chan<- int, exitChan chan int) {
	for i := 1; i <= 10; i++ {
		testChan <- i
	}
	close(testChan)

	exitChan <- 1
}

// 只读管道
func receive(testChan <-chan int, exitChan chan int) {
	for {
		date, ok := <-testChan
		if !ok {
			break
		}
		fmt.Println(date)
	}
	exitChan <- 1
}

func main() {
	go send(testChan, exitChan)
	go receive(testChan, exitChan)
	// for-range遍历时，当使用匿名数据时，不需要加 ":" 去推导
	for _ = range exitChan {
		exit++
		if exit == 2 {
			break
		}
	}
	fmt.Println("Main Over！")
}
