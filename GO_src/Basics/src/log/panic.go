/*
*****************************************************

	Author       : Mr.Chan
	Date         : 2022-11-06 08:24:42
	LastEditors  : Mr.Chan
	LastEditTime : 2022-11-06 08:24:43
	FilePath     : /vivianchan.cn/Basics/src/log/panic.go

*****************************************************
*/
package main

import "log"

func main() {
	data := []int{1, 2}
	log.Panicln(data)
}
