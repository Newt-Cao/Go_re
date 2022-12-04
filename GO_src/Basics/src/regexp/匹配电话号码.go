/*
*****************************************************

	Author       : Mr.Chan
	Date         : 2022-11-04 22:04:35
	LastEditors  : Mr.Chan
	LastEditTime : 2022-11-04 22:44:18
	FilePath     : /vivianchan.cn/Basics/src/正则表达式/匹配电话号码.go

*****************************************************
*/
package main

import (
	"fmt"
	"regexp"
)

func phoneNum(num string) bool {
	re, _ := regexp.Compile(`^1[1-9][0-9]{9}\b`) // 第一位 1 开头，第二位 1-9，第 3-9 位 0-9，的 11 位数字，\b为边界，生成表达式
	result := re.MatchString(num)                // 检索传入的号码是否符合要求
	return result
}

func main() {
	fmt.Println(phoneNum("1000"))           // 少于 11 位 false
	fmt.Println(phoneNum("1234-678000"))    // 特殊字符 false
	fmt.Println(phoneNum("02345678000"))    // 0 开头 false
	fmt.Println(phoneNum("12345678000000")) // 超过 11 位 false
	fmt.Println(phoneNum("10000000000"))    // 第二位不为 0 false
	fmt.Println(phoneNum("11000000000"))    // 正确 true
}
