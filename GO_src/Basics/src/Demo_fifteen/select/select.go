package main

import (
	"fmt"
)

// 斐波那契数列数列：1,1,2,3,5,8,13，...
func fibonacci(c, quit chan int) {
	x, y := 1, 1

	for {
		select {
		// 因为一直重新赋值所以才可以得出数列
		case c <- x:
			y = x + y
			x = y - x
			//x, y = y, x+y

		// 当接收到匿名函数传入的值时，结束——>意味着主程序结束
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		// 当50次循环完毕，向quit管道传入一个值
		for i := 0; i < 50; i++ {
			fmt.Println(<-c)
		}

		quit <- 0
	}()

	fibonacci(c, quit)

}
