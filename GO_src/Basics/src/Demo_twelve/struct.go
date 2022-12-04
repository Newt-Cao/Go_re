/* 结构体的定义与参数传递 */

package main

import "fmt"

// 声明一种新的数据类型 myint ,可以理解是 int 数据的别名。
type myint int

// 定义一个结构体，把多种数据类型结合变成新的复杂数据类型。
type People struct {
	name string
	age  int
}

// 传参方式类似于变量
func changePeople(person People) {
	// 值传递
	person.age = 35
	fmt.Println("姓名：", person.name, "年龄：", person.age)
}

func changePeople2(person2 *People) {
	// 指针传递
	person2.age = 45
	fmt.Println("姓名：", person2.name, "年龄：", person2.age)
}

func main() {
	var a myint = 10
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a) //type of a = main.myint

	// 使用结构体
	var person People
	person.name = "张三"
	person.age = 30

	fmt.Printf("%v\n", person)

	// 调用函数
	changePeople(person)       //姓名： 张三 年龄： 35
	fmt.Printf("%v\n", person) //修改不了原来定义的 person {张三 30}

	changePeople2(&person)     //姓名： 张三 年龄： 45
	fmt.Printf("%v\n", person) //修改了原来定义的 person {张三 45}
}
