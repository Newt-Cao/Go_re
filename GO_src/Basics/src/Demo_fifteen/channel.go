//  channel 协程之间的通信。管道！ x := make(chan type ,capacity)

package main

import "fmt"

// 实现 main 函数与 X 函数的交互
func main() {

	// 定义一个channel管道,接收,数据类型为chan int。（发送与接收默认是阻塞的。这是什么意思？当把数据发送到信道时，程序控制会在发送数据的语句处发生阻塞，直到有其它 Go 协程从信道读取到数据，才会解除阻塞。与此类似，当读取信道的数据时，如果没有其它的协程把数据写入到这个信道，那么读取过程就会一直阻塞着。意味着发送和接收具有同步功能。）
	C := make(chan int) //无缓冲

	// 创建一个协程
	go func() {
		defer fmt.Println("goroutine 结束")

		fmt.Println("goroutine 运行")

		C <- 18002094 //将 18002094 发送给通道 C
	}()

	// 从管道读取一个数据，<-C 是读取管道信息后丢弃。
	ID := <-C // ID接收匿名函数发送的 18002094

	fmt.Println("ID = ", ID)
	fmt.Println("main goroutine 结束")
}
