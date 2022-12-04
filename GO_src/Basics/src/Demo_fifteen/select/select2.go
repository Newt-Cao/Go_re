package main

import "fmt"

func main() {
	// 定义两个管道
	channel01 := make(chan int, 2)
	channel02 := make(chan string, 2)

	// 分别写入数据
	for i := 1; i <= 2; i++ {
		channel01 <- i
	}
	for i := 1; i <= 2; i++ {
		channel02 <- "Vian" + fmt.Sprintf("%v", i)
	}

	for {
		// 结合for利用select分别读取管道数据并打印
		select {
		case x := <-channel01:
			fmt.Printf("channel01 = %v\n", x)
		case x := <-channel02:
			fmt.Printf("channel02 = %s\n", x)

		//当最后两个管道阻塞时，运行结束
		default:
			fmt.Println("Main Over")
			return
		}
	}
}

/*
channel01 = 1
channel02 = Vian1
channel02 = Vian2
channel01 = 2
Main Over
*/
