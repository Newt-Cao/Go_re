/*******************************************************
 *	                       _oo0oo_
 *	                      o8888888o
 *	                      88" . "88
 *	                      (| -_- |)
 *	                      0\  =  /0
 *	                    ___/`---'\___
 *	                  .' \\|     |// '.
 *	                 / \\|||  :  |||// \
 *	                / _||||| -:- |||||- \
 *	               |   | \\\  - /// |   |
 *	               | \_|  ''\---/''  |_/ |
 *	               \  .-\__  '-'  ___/-. /
 *	             ___'. .'  /--.--\  `. .'___
 *	          ."" '<  `.___\_<|>_/___.' >' "".
 *	         | | :  `- \`.;`\ _ /`;.`/ - ` : | |
 *	         \  \ `_.   \_ __\ /__ _/   .-` /  /
 *	     =====`-.____`.___ \_____/___.-`___.-'=====
 *	                       `=---='
 *
 *
 *	     ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *
 *	           佛祖保佑     永不宕机     永无BUG
 *
 *	Author       : Mr.Chan
 *	Date         : 2022-08-30 09:03:33
 *	LastEditors  : Mr.Chan
 *	LastEditTime : 2022-08-30 09:34:30
 *	FilePath     : /vivianchan.cn/Basics/src/Demo_fourteen/timer.go
 *******************************************************/

package main

import (
	"fmt"
	"time"
)

// 几种定时器的使用：
func main() {
	timer := time.NewTimer(2 * time.Second) // 两秒
	fmt.Println(time.Now())                 // 2022-08-30 09:22:04.7571044 +0800 CST m=+0.001335401
	t1 := <-timer.C                         // timer.C 就是一个 chan Time 类型的管道，阻塞两秒
	fmt.Println(t1)                         // 2022-08-30 09:22:06.7719841 +0800 CST m=+2.016215101

	time.Sleep(2 * time.Second) // 睡眠两秒

	<-time.After(2 * time.Second) // 两秒之后写入管道，这里也是chan Time的C管道
	fmt.Println("after end...")

	// 中断计时器
	timer2 := time.NewTimer(2 * time.Second)
	go func() {
		<-timer2.C
		// 计时器中断后下面代码段不再执行
		fmt.Println("func end...")
	}()

	// 中断计时器返回布尔类型
	stop := timer2.Stop()
	if stop {
		fmt.Println("stop end...")
	}

	time.Sleep(3 * time.Second)
	fmt.Println("main end...")

	// 修改计时器时间
	timer3 := time.NewTimer(2 * time.Second) // 两秒
	timer3.Reset(time.Second)                // 修改成一秒
	t2 := <-timer3.C
	fmt.Println(t2)
}
