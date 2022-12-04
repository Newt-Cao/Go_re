/* channel 的关闭 */

/*
1.不需要经常关闭，当确定没有任何信息或数据发送时、循环结束时，再去关闭channel
2.关闭channel后无法继续向其发送数据（引发panic错误后导致接收返回零值）
3.关闭channel后可继续从中接收数据
4.对于nil channel（只定义 var xxx chan），无论收发都会阻塞
*/

package main

import "fmt"

func main() {

	c := make(chan int, 5)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		//执行完上述发送命令到channel中后关闭channel。
		close(c) //如果不关闭则会报错，主程序会一直等待接收数据，然而子程序已经结束，无法继续写入数据，导致报错，死锁状态。

	}()

	//死循环
	for {
		// if的一种写法，先执行(data,ok := <-c)，后单独判断 ok 的返回值，ok如果为true表示channel未关闭，false表示channel已关闭。
		if data, ok := <-c; ok {
			fmt.Println(data) //未关闭打印data
		} else {
			break //关闭则结束
		}
	}

	fmt.Println("Finished!")

}
