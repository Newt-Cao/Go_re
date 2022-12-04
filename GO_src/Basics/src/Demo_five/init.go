/* 导包顺序 */

package main

// 正常情况要加GOPATH路径才可以正常导包，使用 go module 后可管理模块，简化导包路径。
// 导包的时候会优先调用包里的init函数，直到没有包为止(可用于初始化，加载某些数据库).
// 一个go的优先级:全局var > init > main
// 引入go包优先级：包全局var > 包init > main var > main init> main
import (
	"fmt"
	// 先调用lib1里的初始化init函数
	"VivianChan.cn/Basics/src/Demo_five/lib1"
)

// 再调用main包里的init函数
func init() {
	fmt.Println("main.init()")
}

// 最后执行
func main() {
	// 调用lib1包里的函数
	lib1.Lib1Test()
}
