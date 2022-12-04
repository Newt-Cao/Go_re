/* 类的多态 接口interface */

package main

import "fmt"

// 本质是一个指针，指向所有类型，函数列表，接口。
type Animal interface {
	Sleep()
	Getcolor() string //获取动物的颜色
	Gettype() string  //获取动物的种类
}

type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is Sleep")
}

func (this *Cat) Getcolor() string {
	return this.color
}

func (this *Cat) Gettype() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is Sleep")
}

func (this *Dog) Getcolor() string {
	return this.color
}

func (this *Dog) Gettype() string {
	return "Dog"
}

// 多态实现
func AnimalShow(animal2 Animal) {
	animal2.Sleep()
	fmt.Println("color = ", animal2.Getcolor(), "kind = ", animal2.Gettype())
}

func main() {
	// 定义一个接口类型，父类指针，指向子类的方法
	var animal Animal
	animal = &Cat{"blue"}
	animal.Sleep() //利用接口调用结构体 Cat 的方法

	cat := Cat{"蓝色"}
	dog := Dog{"黄色"}

	AnimalShow(&cat)
	AnimalShow(&dog)
	/*
		Cat is Sleep
		Cat is Sleep
		color =  蓝色 kind =  Cat
		Dog is Sleep
		color =  黄色 kind =  Dog
	*/
}
