/*
*****************************************************

	Author       : Mr.Chan
	Date         : 2022-11-03 21:41:45
	LastEditors  : Mr.Chan
	LastEditTime : 2022-11-03 21:41:49
	FilePath     : /vivianchan.cn/Basics/src/正则表达式/Compile.go

*****************************************************
*/
package main

import (
	"fmt"
	"log"
	re "regexp"
)

func main() {
	str := "123abc"                  // 被检测对象
	regexp, err := re.Compile(`\d+`) // 生成正则对象，匹配所有单个数字，非数字返回空
	if err != nil {
		log.Panic(err)
	}

	text := "Hello GoLang."                 // 被检测对象
	regexp2 := re.MustCompile(`\w+`)        // 生成正则对象，匹配非标点符号以外的所有单个字符
	fmt.Println(str, text, regexp, regexp2) // 123abc Hello GoLang. \d+ \w+
}
