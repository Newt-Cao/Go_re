/******************************************************
 Author       : Mr.Chan
 Date         : 2022-11-10 10:43:43
 LastEditors  : Mr.Chan
 LastEditTime : 2022-11-10 15:58:38
 FilePath     : /vivianchan.cn/Basics/src/sync/RWMutex/case2.go
******************************************************/

// 多个读锁使用

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

func main() {
	go Rmutex(1)
	go Rmutex(2)
	time.Sleep(2 * time.Second)
}
