/* channel 缓冲的定义与用法 */

package main

import (
	"fmt"
	"time"
)

func main() {

	defer fmt.Println("父进程结束")

	// 定义一个有缓冲的channel
	C := make(chan int, 3) // 有缓冲
	// len() 元素个数 ，cap()容量
	fmt.Println("len(C) = ", len(C), ", cap(C) = ", cap(C))

	go func() {

		defer fmt.Println("子程序结束")

		// 遍历向 C 发送元素
		for i := 0; i < 5; i++ {
			C <- i //把遍历的i元素发送给 channel C
			fmt.Println("元素：", i, "len(C):", len(C), ", cap(C)", cap(C))
		}
	}()

	time.Sleep(2 * time.Second)

	//遍历接收
	for i := 0; i < 5; i++ {
		num := <-C

		fmt.Println("num = ", num)
	}
}

/*
1.当channel中容量满时再向里面写数据会发生阻塞
2.当channel中容量空时再向里面读数据会发生阻塞
*/
