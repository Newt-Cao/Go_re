/* 指针传递 */

package main

import "fmt"

// 定义一个变量时，系统会开辟：变量名，值（内存），内存地址
// value,值传递.

/* func changeValue(p int) //p默认是0。输入a后， p = a ，p = 1 修改了值，值传递 {
	p = 10
}

func main() {
	var a int = 1

	changeValue(a)

	fmt.Println("a = ", a) //a = 1
} */

// pointer,指针传递.参数的数据类型前加 * ，表示传递地址类型  *int -->整形地址数据类型。
//                p有一个0x0000000默认地址类型，传入地址&a后，p的值指向a的地址，p自己的有地址，涉及到层级指针（二级，三级）
// 可以通过 p 找到 a的内存，然后修改。
func changeValue2(p *int) {
	// 修改对应的传入地址，指针所指。
	*p = 10
}

func main() {
	var a int = 1

	// 在变量名前加 & ，表示传入该变量的地址
	changeValue2(&a)

	fmt.Println("ap = ", a) //ap = 10
}
