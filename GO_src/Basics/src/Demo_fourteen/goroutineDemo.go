package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	// 定义一个匿名函数，无参数。协程go无返回值。
	go func() {

		defer fmt.Println("Print A")

		func() {

			defer fmt.Println("Print B")

			runtime.Goexit() //退出当前程序，结束协程，只运行两个defer

			fmt.Println("Print b")

		}() //函数调用，自调。

		fmt.Println("Print a")

	}()

	/* 带参数的函数，go协程无返回值，不可以直接接收。

	 go func(a,b int)bool{
		fmt.Println("a = ",a,"b = ",b)
		return true
	}(10,20)  //传入参数

	*/

	// 死循环
	for {
		time.Sleep(1 * time.Second)
	}

}
