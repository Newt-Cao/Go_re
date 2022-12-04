/*
*****************************************************

	Author       : Mr.Chan
	Date         : 2022-11-04 23:28:50
	LastEditors  : Mr.Chan
	LastEditTime : 2022-11-04 23:28:51
	FilePath     : /vivianchan.cn/Basics/src/正则表达式/匹配Email.go

*****************************************************
*/
package main

import (
	"fmt"
	"regexp"
)

func emailMatch(email string) bool {
	re := regexp.MustCompile(`^\w+@\w+\.\w+`)
	result := re.MatchString(email)
	return result
}

func main() {
	fmt.Println(emailMatch("1233@qq.com"))    // true
	fmt.Println(emailMatch("123.@qq.com"))    // false
	fmt.Println(emailMatch("123@qqcom"))      // false
	fmt.Println(emailMatch("123.232@qq.com")) // false
	fmt.Println(emailMatch("123232qq.com"))   // false
}
