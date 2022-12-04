/*
*****************************************************

	Author       : Mr.Chan
	Date         : 2022-11-10 09:51:20
	LastEditors  : Mr.Chan
	LastEditTime : 2022-11-10 09:51:20
	FilePath     : /vivianchan.cn/Basics/src/sync/RWMutex/case.go

*****************************************************
*/

// 读写锁使用

package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// 全局变量
var (
	rwm   sync.RWMutex
	count int
)

// 读锁使用
func Rmutex(n int, exit chan int) {
	rwm.RLock() // 读锁，准备读取
	fmt.Printf("机器 %d 读取中。。。", n)
	reader := count // 模拟读取到其他地方
	fmt.Println("读取成功，读取的值为：", reader)
	rwm.RUnlock() // 解读锁
	exit <- -1    // 读取结束发送退出码
}

// 写锁使用
func Wmutex(n int, exit chan int) {
	rwm.Lock() // 写锁，此时的count由当前的go协程独享
	fmt.Printf("机器 %d 写入中。。。", n)
	count = rand.Intn(100) // 随机写入一个 0-9 范围的数
	fmt.Println("写入成功，写入的值为：", count)
	rwm.Unlock() // 解写锁后其他go协程才可以使用count
	exit <- -1   // 写入结束发送退出码
}

func main() {
	exit := make(chan int, 6)

	// 开启读协程
	for i := 0; i < 3; i++ {
		go Rmutex(i, exit) // 同时开启多个读
	}

	// 开启写
	for i := 0; i < 3; i++ {
		go Wmutex(i, exit) // 一次只有一个函数执行，执行完后才可以进行下一次的读或写
	}

	// 等待协程结束
	for i := 0; i < 6; i++ {
		<-exit
	}
}
