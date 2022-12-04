/*
*****************************************************

	Author       : Mr.Chan
	Date         : 2022-11-11 15:39:58
	LastEditors  : Mr.Chan
	LastEditTime : 2022-11-11 15:41:57
	FilePath     : /vivianchan.cn/Basics/src/sync/race/raceCase.go

*****************************************************
*/
package main

import "fmt"

// 同时访问一个映射
func main() {
	exit := make(chan int)
	m := make(map[string]interface{})

	go func() {
		m["vian"] = 18 // 第一个访问
		exit <- -1
	}()

	m["jack"] = 45 // 第二个访问
	<-exit         // 等待协程

	for i, v := range m {
		fmt.Printf("%s:%v", i, v)
	}
}
