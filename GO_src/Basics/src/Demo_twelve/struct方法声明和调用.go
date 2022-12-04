package main

import "fmt"

// 声明结构体
type Person struct {
	Name string
}

// 方法声明和绑定
/*
	说明：
		1.语法是 func(变量 结构体或自定义类型)方法名( 参数列表 ) 返回值列表 {}
		2.括号里的变量名，相当于函数中的形参，是实例化对象的副本，可任意改变
		3.参数传递是值拷贝，只在方法里生效
		4.让某个方法和结构体绑定，括号里就使用对应的结构体类型，此处让test()和Person绑定
		5.绑定后test是Person的方法，只能通过Person的实例化对象来调用，不能直接调用（无意义），也不能使用其他类型变量来调用
		6.name可用可不用，是拷贝实例化对象的副本，不论用不用，都会拷贝一份
		7.自定义类型也可以绑定方法

*/
func (name Person) test() {
	name.Name = "Jack"
	fmt.Println("姓名：", name)
}

type Student struct {
	Name string
}

// 如果给结构体绑定 String 方法，那么在主函数输出实例化对象时，会自动调用
func (stu *Student) String() string {
	str := fmt.Sprintf("Name=%v", stu.Name)
	return str
}

func main() {
	// 声明一个Person对象
	person := Person{}
	person.Name = "Vian"
	fmt.Println(person) // {Vian}
	// 方法调用，只可以Person类型的对象调用
	person.test() // 姓名： {Jack}

	// 绑定 String 方法后，自动调用
	stu := Student{"Chow"}
	fmt.Println(&stu) // Name=Chow 自动调用了 String 方法

}
