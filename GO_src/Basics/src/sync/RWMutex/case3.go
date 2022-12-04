/******************************************************
 Author       : Mr.Chan
 Date         : 2022-11-10 10:47:01
 LastEditors  : Mr.Chan
 LastEditTime : 2022-11-10 16:07:43
 FilePath     : /vivianchan.cn/Basics/src/sync/RWMutex/case3.go
******************************************************/

// 读写锁互斥案例

package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	rwm sync.RWMutex
)

func Rmutex(i int) {
	fmt.Println(i, "开始读取。")
	rwm.RLock() // 读锁
	fmt.Println(i, "读取中...")
	time.Sleep(time.Second)
	rwm.RUnlock() // 解读锁
	fmt.Println(i, "读取结束。")
}

func Wmutex(i int) {
	fmt.Println(i, "开始写入。")
	rwm.Lock() // 读锁
	fmt.Println(i, "写入中...")
	time.Sleep(time.Second)
	fmt.Println(i, "写入结束。")
	rwm.Unlock() // 解读锁
}

func main() {
	Wmutex(0) // 写入时下面的goroutine都无法开始
	go Rmutex(1)
	go Wmutex(2)
	time.Sleep(3 * time.Second)
}
