/*
*****************************************************

	Author       : Mr.Chan
	Date         : 2022-11-06 00:43:30
	LastEditors  : Mr.Chan
	LastEditTime : 2022-11-06 00:43:30
	FilePath     : /vivianchan.cn/Basics/src/log/print.go

*****************************************************
*/
package main

import "log"

func main() {
	data := []int{1, 2}
	log.Print("Print", data, "\n")
	log.Printf("Printf [%d:%d]", data[0], data[1])
	log.Println("Println", data)
	/*
		2022/11/06 00:46:20 Print[1 2]
		2022/11/06 00:46:20 Printf [1:2]
		2022/11/06 00:46:20 Println [1 2]
	*/
}
