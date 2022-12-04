/*
*****************************************************

	Author       : Mr.Chan
	Date         : 2022-11-04 20:26:57
	LastEditors  : Mr.Chan
	LastEditTime : 2022-11-04 20:26:58
	FilePath     : /vivianchan.cn/Basics/src/正则表达式/Index.go

*****************************************************
*/
package main

import (
	"fmt"
	re "regexp"
)

func main() {
	text := "abc oaabssj obj"
	reg_expression, _ := re.Compile(`obj?`)
	fmt.Println(reg_expression.FindIndex([]byte(text))) // [12 15]

	text2 := "abc oaabssj obj obj"
	fmt.Println(reg_expression.FindAllIndex([]byte(text2), -1))        // [[12 15] [16 19]]，第二个参数是匹配长度，-1 表示匹配到末尾
	fmt.Println(reg_expression.FindAllIndex([]byte("ggg"), -1) == nil) // 匹配不到返回 nil，true

	/*
		FindStringIndex()
		FindAllStringIndex()
		用法同上，只是类型 []byte() 换成 string
	*/
}
