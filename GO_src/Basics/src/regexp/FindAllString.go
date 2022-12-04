/*
*****************************************************

	Author       : Mr.Chan
	Date         : 2022-11-04 14:18:20
	LastEditors  : Mr.Chan
	LastEditTime : 2022-11-04 14:18:21
	FilePath     : /vivianchan.cn/Basics/src/正则表达式/FindAllString.go

*****************************************************
*/
package main

import (
	"fmt"
	re "regexp"
)

func main() {
	str := "112233 vian golang python"
	re_expression, _ := re.Compile(`\D*`)                         // 匹配所有非数字字符
	re_result := re_expression.FindAllString(str, -1)             // 第二个参数是限定查找的数量，-1 表示不限制
	fmt.Printf("[]string = %v , type = %T", re_result, re_result) // []string = [       vian golang python] , type = []string
}
