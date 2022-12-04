package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name string
	Age  int
}

func varChange(x interface{}) {
	// 转成 reflect.Value 类型，这里接收的是一个地址
	changeLate := reflect.ValueOf(x)
	// 获取类型对应的值后更改，这里的类型是地址
	changeLate.Elem().SetString("Jack")
}

func main() {
	// 更改变量值
	// 变量
	Name := "Vian"
	fmt.Println(Name) // Vian
	// 函数调用更改，因为需要更改原来的值，所以传递一个地址
	varChange(&Name)
	fmt.Println(Name) // Jack
	// 直接调用更改
	NameChange := reflect.ValueOf(&Name)
	NameChange.Elem().SetString("Niko")
	fmt.Println(Name) // Niko

	// 更改结构体字段值
	// 实例
	animal := Animal{"大笨熊", 20}
	fmt.Println(animal) // {大笨熊 20}

	// 方法一：
	animalChange := reflect.ValueOf(&animal)
	animalChange.Elem().Field(0).SetString("大蜘蛛")
	fmt.Println(animal) // {大蜘蛛 20}

	// 方法二：
	Animal := &Animal{}
	AnimalChange := reflect.ValueOf(Animal)
	// 指针类型
	fmt.Println(AnimalChange.Kind().String()) // ptr
	// 对应的是一个结构体
	fmt.Println(AnimalChange.Elem().Kind().String()) // struct
	// 第一个字段
	AnimalChange.Elem().FieldByName("Name").SetString("小海豚")
	// 第二个字段
	AnimalChange.Elem().FieldByName("Age").SetInt(18)
	fmt.Println(Animal) // &{小海豚 18}
}
