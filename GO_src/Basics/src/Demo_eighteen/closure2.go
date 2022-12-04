package main

import (
	"fmt"
	"strings"
)

// 修改文件名后缀

func makeSuffix(suffix string) func(string) string {
	// 匿名函数引用到了makeSuffix函数的形参suffix，整体构成了闭包
	return func(name string) string {
		// strings.HasSuffix(后缀,文件名)判断后缀是否是指定后缀
		if !strings.HasSuffix(suffix, name) {
			// 字符串拼接
			return name + suffix
		} else {
			return name
		}
	}
}

func main() {
	m := makeSuffix(".go")
	fmt.Printf("文件修改后：%s\n", m("golang"))
	fmt.Printf("文件修改后：%s\n", m("golang.go"))
	fmt.Printf("文件修改后：%s", m("golang.jpg"))
}
