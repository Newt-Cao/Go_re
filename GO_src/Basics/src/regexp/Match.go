/*
*****************************************************

	Author       : Mr.Chan
	Date         : 2022-11-03 22:01:55
	LastEditors  : Mr.Chan
	LastEditTime : 2022-11-03 22:01:56
	FilePath     : /vivianchan.cn/Basics/src/正则表达式/Match.go

*****************************************************
*/
package main

import (
	"fmt"
	re "regexp"
)

/*
	func MatchReader(pattern string, r io.RuneReader) (matched bool, err error)
	func MatchString(pattern string, s string) (matched bool, err error)
	func Match(pattern string, b []byte) (matched bool, err error)
*/

func main() {
	text := "Hello GoLang."                // 被检测对象
	regexp2 := re.MustCompile(`\w+`)       // 生成正则对象，匹配非标点符号以外的所有单个字符
	fmt.Println(regexp2.MatchString(text)) // true

	fmt.Println(re.MatchString("H(.*)d!", "Hello World!"))   // 链式检测，true <nil>
	fmt.Println(re.Match("H(.*)d!", []byte("Hello World!"))) // 链式检测，true <nil>
}
