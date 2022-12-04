/* 类的继承 */

package main

import "fmt"

// 定义一个父类
type Human struct {
	Name string
	age  int
	sex  string
}

// 绑定结构体 Human 方法
func (this *Human) Eat() {
	fmt.Println("Human Eat")
}

func (this *Human) Walk() {
	fmt.Println("Human Walk")
}

// 定义一个子类（添加父类缺少的数据或者方法）
type SuperMan struct {
	// 把父类类名放入子类里面表示继承
	Human
	Level int
}

// 子类可以重写父类方法
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan Eat")
}

// 子类新方法
func (this *SuperMan) Fly() {
	fmt.Println("SuperMan Fly")
}

func main() {
	// 定义一个 Human 对象
	human := Human{"张三", 21, "男"}
	// 调用 类方法
	human.Eat()
	human.Walk()

	fmt.Println("===================================")

	// 定义一个子对象

	// 一：
	superman := SuperMan{Human{"李四", 25, "男"}, 5}

	superman.Walk() //父
	superman.Eat()  //子
	superman.Fly()  //子

	fmt.Println("===================================")

	// 二：
	var superman2 SuperMan
	superman2.Name = "王五"
	superman2.age = 30
	superman2.Level = 10

	superman.Walk() //父
	superman.Eat()  //子
	superman.Fly()  //子

}
