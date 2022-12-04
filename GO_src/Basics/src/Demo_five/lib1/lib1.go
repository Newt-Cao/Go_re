/* 模拟第三方包 */

package lib1

import "fmt"

// lib包提供的API接口（函数的第一个字母大写才可以被外部调用，小写只能在包内调用。）
func Lib1Test() {
	fmt.Println("Lib1Test()")
}

func init() {
	fmt.Println("lib1.init()")
}
