/* append */

package main

import "fmt"

// len：表示左指针至右指针之间的距离
// cap：表示左指针至底层数组末尾的距离
// append：本质上就是一个底层数组的值拷贝，对底层数组的扩容，再赋值，地址不变。（因为数组一旦确定，无法改变，只能通过改变len和cap后再改变数组，切片的本质就是一个动态数组，可以通过append的方式重新赋值，以达到扩容的效果。）

/*
	append的工作流程：
		1.在底层创建一个长度为（原切片数组长度+扩容元素长度）的数组。
		2.再将原切片数组的值和扩容元素的值拷贝到该新数组中。
		3.最后再重新赋值给原切片。（也可以赋值到其他切片，但是没有意义，原来的切片并没有变化）
		4.原切片的底层数组销毁。
*/

func main() {
	// make(type,len,cap)
	// len = 3,cap = 5,[0,0,0]
	var numbers = make([]int, 3, 5)
	fmt.Printf("len = %d,cap = %d,slice = %v\n", len(numbers), cap(numbers), numbers)

	// 向切片追加一个元素，len = 4 ,cap = 5 ,[0,0,0,1]
	numbers = append(numbers, 1)
	fmt.Printf("len = %d,cap = %d,slice = %v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 2) //len = 5 ,cap = 5 ,[0,0,0,1,2]
	fmt.Printf("len = %d,cap = %d,slice = %v\n", len(numbers), cap(numbers), numbers)

	// 继续追加时，当超过容量 cap 时 ，cap 自动 * 2，len = 6,cap = 10,[0 0 0 1 2 3]
	numbers = append(numbers, 3)
	fmt.Printf("len = %d,cap = %d,slice = %v\n", len(numbers), cap(numbers), numbers)

	// make（）声明时没有初始化cap时，默认 2cap = len ,当超出cap时，len*2 = cap*2
	var numbers2 = make([]int, 3) // len = 3,cap = 3,slice = [0 0 0]
	fmt.Printf("len = %d,cap = %d,slice = %v\n", len(numbers2), cap(numbers2), numbers2)

}
