package personModule

import (
	"fmt"
)

/*
练习二：
	设计一个工资程序，不可以查看别人的年龄、工资等隐私，并对输入的年龄进行判断。
*/

type person struct {
	Name   string
	age    int
	salary float64
}

func PersonCreat(name string) *person {
	return &person{
		Name: name,
	}
}

func (per *person) SetAge(age int) {
	if age < 18 || age >= 150 {
		fmt.Println("年龄不符合！")
	} else {
		(*per).age = age
		fmt.Println("年龄修改成功！")
	}
}

func (per *person) GetAge() {
	fmt.Println("您的年龄是：", (*per).age)
}

func (per *person) SetPalary(salary float64) {
	(*per).salary = salary
	fmt.Println("工资修改成功！")
}

func (per *person) GetPalary() {
	fmt.Println("您的工资是:", (*per).salary)
}
