/* 指针案例 */

package main

import "fmt"

// 值传递,无论 swap里的 a和b如何改变，无法影响到 main函数里的a和b。
/* func swap(a, b int) {
	// 交换值写法
	var temp int //定义一个变量，默认为 0
	temp = a     //使该变量等于第一个变量的值
	a = b        //使第一个变量的值等于第二个变量的值
	b = temp     //再让第二个变量的值等于第一个变量的值

	// 方法二
	a , b = b , a
}

func main() {
	var a int = 10
	var b int = 20

	// 交换
	swap(a, b)

	fmt.Println("a = ", a, "b = ", b) //a =  10 b =  20
}
*/

//  地址传递（指针）
// swap里 *a和*b ，直接指向main函数里的 a和b ,实现地址交换。
func swap(a, b *int) {
	var temp int
	temp = *a //temp = main:a
	*a = *b   //main:a = main:b
	*b = temp //main:b = temp
}

func main() {
	var a int = 10
	var b int = 20

	// 交换
	swap(&a, &b)

	fmt.Println("a = ", a, "b = ", b) //a =  20 b =  10
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", &a)
	// 获取a指向的地址的值
	fmt.Printf("%v\n", *&a)

	// 层级指针

	//一级
	var p *int

	p = &a

	fmt.Println(p)
	fmt.Println(&a)
	fmt.Println(*p)
	// 二级
	var pp **int

	pp = &p

	fmt.Println(&p)
	fmt.Println(*pp)
	fmt.Println(&pp)

	// 可以通过二级指针找到一级指针，再通过一级指针找到对应的变量地址。
}
