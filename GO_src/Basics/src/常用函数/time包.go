package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	// 一、查看当前时间
	fmt.Println("时间：", time.Now()) //时间： 2022-05-18 10:22:10.4621978 +0800 CST m=+0.002277501

	// 二、格式化时间日期
	// 第一种方式：
	fmt.Printf("%d-%d-%d %d:%d:%d\n", time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second()) //2022-5-18 10:22:10
	// Sprintf()获取Printf的string，可以复制给一个变量，存入数据库
	dateStr := fmt.Sprintf("%d-%d-%d %d:%d:%d\n", time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second())
	fmt.Println(dateStr) //2022-5-18 10:22:10
	// 第二种方式（日期和时间可以分开获取，也可以单独获取月份或年：01/2006。数字不能改，格式可以改）：
	fmt.Println(time.Now().Format("2006-01-02 15:04:05")) //2022-05-18 10:22:10

	// 三、时间常量
	/*
		const(
			Nanosecond Duration = 1
			Microsecond = 1000 * Nanosecond
			Millisecond = 1000 * Microsecond
			Second = 1000 * Millisecond
			Minute = 60 * Second
			Hour = 60 * Minute
		)
	*/
	// 休眠，每隔100毫秒输出一个数
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Println(i)
	}

	// 四、时间戳,Unix和UnixNano（用于获取伪随机数字）
	fmt.Printf("Unix:%d UnixNano:%d", time.Now().Unix(), time.Now().UnixNano()) //Unix:1652840531 UnixNano:165284053154550570040
	// (1)、获取随机数
	rand.Seed(time.Now().UnixNano())
	for {
		var Count int
		r := rand.Intn(100)
		fmt.Println(r)
		if r == 99 {
			break
		}
		Count++
	}
	// (2)、获取代码执行时间(爬虫)
	start := time.Now().Unix()

	func() {
		var str string
		for i := 0; i <= 100000; i++ {
			str += "golang" + strconv.Itoa(i)
		}
	}()

	end := time.Now().Unix()

	fmt.Printf("时间：%v 秒", end-start) //时间：4 秒
}
