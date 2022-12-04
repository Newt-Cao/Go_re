package main

import "fmt"

// 接口
type Airplane interface {
	Fly()
	Down()
}

// 结构体
type Warcraft struct {
}

// 实现接口方法
func (plane *Warcraft) Fly() {
	fmt.Println("A plane is flying...")
}

func (plane *Warcraft) Down() {
	fmt.Println("A plane is landing...")
}

func main() {
	var plane Warcraft
	// 将实现了接口的结构体实例化赋值给接口
	var airplane Airplane = &plane
	airplane.Fly()  // A plane is flying...
	airplane.Down() // A plane is landing...
}
