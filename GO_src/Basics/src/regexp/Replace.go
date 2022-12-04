/*
*****************************************************

	Author       : Mr.Chan
	Date         : 2022-11-04 20:52:49
	LastEditors  : Mr.Chan
	LastEditTime : 2022-11-04 20:52:49
	FilePath     : /vivianchan.cn/Basics/src/正则表达式/Replace.go

*****************************************************
*/
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`w(a*)i`)

	/*
		$n 匹配的解析：匹配开始后被替换成匹配第 n 个分组 (a*) 匹配的子串 wi ，wi 中间为空，所以替换成空，然后匹配到末尾将
					 所有符合 w(a*)i 表达式的 wi 全部替换成空。
	*/
	fmt.Printf("%s", re.ReplaceAll([]byte("-wi-waaaaai-wai-wcci-\n"), []byte("$1")))    // --aaaaa-a-wcci-
	fmt.Printf("%s", re.ReplaceAll([]byte("-wi-waaaaai-wai-wcci-\n"), []byte("$1W")))   // ----wcci-，将满足条件的全部替换为空
	fmt.Printf("%s", re.ReplaceAll([]byte("-wi-waaaaai-wai-wcci-\n"), []byte("${1}W"))) // -W-aaaaaW-aW-wcci-
	fmt.Printf("%s", re.ReplaceAll([]byte("-wi-waaaaai-wai-wcci-\n"), []byte("${1}")))  // --aaaaa-a-wcci-，${1}匹配第一个(a*)
}
