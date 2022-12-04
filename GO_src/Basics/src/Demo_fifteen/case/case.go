/*
	案例一：
		编写两个函数，使用协程分别是向通道写入和向通道读取，要求操作同一个通道，主程序需要等待子程序执行完毕后结束。
*/

package main

import (
	"fmt"
)

// 写入
func chanWriter(chanwriter chan int) {
	for i := 1; i <= 20; i++ {
		chanwriter <- i
	}
	// 写入完毕后，关闭通道
	close(chanwriter)
}

// 读取
func chanReader(chanwriter chan int, chanjudge chan bool) {
	// 不断向管道读取内容
	for {
		v, ok := <-chanwriter
		// 判断是否读取完毕
		if !ok {
			fmt.Println("读取完毕！")
			// 读取完毕表示子程序结束，将false传入退出判断管道
			chanjudge <- ok
			break
		}
		fmt.Println("读取数据：", v)
	}
}

func main() {
	// 声明读写操作管道，退出判断管道
	chanwriter := make(chan int, 20)
	chanjudge := make(chan bool)
	// 调用协程和函数
	go chanWriter(chanwriter)
	go chanReader(chanwriter, chanjudge)
	// 判断chanjudge管道接收值
	for {
		// 读取子程序是否已经结束，false结束则退出主程序，整体结束
		v, _ := <-chanjudge
		if !v {
			fmt.Println("主程序结束！")
			break
		}
	}
}

/*
读取数据： 1
读取数据： 2
读取数据： 3
读取数据： 4
读取数据： 5
读取数据： 6
读取数据： 7
读取数据： 8
读取数据： 9
读取数据： 10
读取数据： 11
读取数据： 12
读取数据： 13
读取数据： 14
读取数据： 15
读取数据： 16
读取数据： 17
读取数据： 18
读取数据： 19
读取数据： 20
读取完毕！
主程序结束！
*/
