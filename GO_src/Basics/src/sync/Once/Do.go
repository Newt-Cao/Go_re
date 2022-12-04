/*
*****************************************************

	Author       : Mr.Chan
	Date         : 2022-11-10 17:04:30
	LastEditors  : Mr.Chan
	LastEditTime : 2022-11-10 17:04:31
	FilePath     : /vivianchan.cn/Basics/src/sync/Once/Do.go

*****************************************************
*/

/*
	在关闭管道时，如果关闭一个已经关闭的管道会报错，可用 sync.Once.Do() 来确保程序运行过程中管道只被关闭一次
*/

package main

import (
	"fmt"
	"sync"
)

var (
	once     sync.Once
	wg       sync.WaitGroup
	intChan  = make(chan int, 10)
	intChan2 = make(chan int, 10)
)

// 写入管道1
func test1(intChan chan<- int) {
	defer wg.Done() // 每当一个协程完成后减1
	for i := 0; i < 10; i++ {
		intChan <- i // 传入
	}
	close(intChan) // 传入完毕后关闭管道
}

// 读取管道1的数据，写入管道2
func test2(intChan <-chan int, intChan2 chan<- int) {
	defer wg.Done()
	for {
		n, ok := <-intChan // 读取管道1的数据
		if !ok {
			break // ok = false表示读取不到数据，管道已关闭
		}

		intChan2 <- n * 2 // 将读取的数据乘以2后写入管道2
	}
	once.Do(func() { close(intChan2) }) // 匿名函数不需要加 () ，Do函数会直接调用，且确保 close(intChan2) 只执行一次
}

func main() {

	// 执行一次 test1 ，两次 test2 ，但是 close(intChan2) 只执行一次
	go func() {
		wg.Add(1)
		test1(intChan)
	}()
	go func() {
		wg.Add(1)
		test2(intChan, intChan2)
	}()
	go func() {
		wg.Add(1)
		test2(intChan, intChan2)
	}()

	wg.Wait() // 等待协程结束

	// 遍历管道2
	for result := range intChan2 {
		fmt.Println(result)
	}
}
