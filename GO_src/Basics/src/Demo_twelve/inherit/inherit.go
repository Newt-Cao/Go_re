package main

import (
	"fmt"
)

type Student struct {
	Name string
}

type Subject struct {
	// 指针类型
	Scores *float64
}

type Credit_Hour struct {
	Student
	// 指针类型
	*Subject
}

func main() {
	stu := Credit_Hour{

		Student{
			Name: "Vian",
		},

		// 传入地址
		&Subject{
			// 需要开辟地址传入
			Scores: new(float64),
		},
	}
	// 赋值
	(*stu.Scores) = 100
	fmt.Println(stu.Name, *stu.Scores) // Vian 100
}
