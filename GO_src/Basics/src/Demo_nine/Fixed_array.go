/* 固定数组 */

package main

import "fmt"

/*
数组的首地址为第一个元素的地址，尾地址为最后一个元素的地址，每个元素的地址是对应数组类型的字节递增。
比如 [...]int{1,2} 的第二个元素地址，为第一个元素地址加int所占的字节。int的字节为8。
*/

// 数组的参数传递(严格匹配数组类型)
func printArray(myArray [4]int) {
	// 值拷贝
	for index, value := range myArray {
		fmt.Println("index = ", index, "value = ", value)
	}

	// 不起作用
	myArray[0] = 111

}

func main() {

	// 定义一个固定长度的数组，不同长度的数组的数据类型不同，传参时要指定对应的长度的数据类型。
	var myArray [10]int //定义一个长度为10的数组，默认值为0。

	myArray1 := [10]int{1, 2, 3, 4} //定义一个长度为10的数组，传入前四个元素的值。

	myArray2 := [4]int{11, 22, 33, 44} //定义一个长度为4的数组，传入前四个元素的值。

	// 自动推断数组的定义方法（可以定义索引）
	myArray3 := [...]int{1, 2, 3: 10, 4, 5}

	// 修改数组元素
	myArray3[0] = 10

	// 遍历数组的两种方法：
	// for i := 0;i < 10;i++ {}
	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}

	// 使用range遍历时返回两个值： 索引，值
	for index, value := range myArray1 {
		fmt.Println("index = ", index, "value = ", value)
	}
	// 类型比较: myArray、myArray1 类型为 [10]int ， myArray2 类型为[4]int.
	fmt.Printf("type of myArray = %T\n", myArray)
	fmt.Printf("type of myArray1 = %T\n", myArray1)
	fmt.Printf("type of myArray2 = %T\n", myArray2)
	fmt.Printf("type of myArray3 = %T\n", myArray3)
	fmt.Printf("type of myArray3 = %v\n", myArray3)

	// 传递参数函数调用
	printArray(myArray2)

	array := [2]int{1, 2}
	fmt.Printf("array的首地址：%p\narray[0]第一个元素的地址：%p\n", &array, &array[0])
	/*
		array的首地址：0xc00001a0d0
		array[0]第一个元素的地址：0xc00001a0d0
	*/
	fmt.Printf("array[0]第一个元素的地址：%p\narray[1]第二个元素的地址：%p", &array[0], &array[1])
	/*
		array[0]第一个元素的地址：0xc00001a0d0 + 8 --> 第二个元素的地址
		array[1]第二个元素的地址：0xc00001a0d8
	*/
}
