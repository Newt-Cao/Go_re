package main

import (
	"accountModule"
	"fmt"
	"personModule"
)

func main() {
	// 练习一
	NewAcc, err := accountModule.AccountClass("VivianChow", 101, "88888888")
	if err == nil {
		fmt.Println(*NewAcc, "账号创建成功！")
	} else {
		fmt.Println(err)
	}

	NewAcc.GetUserAccount()
	NewAcc.GetBalance()
	NewAcc.GetPassword()

	NewAcc.SetBalance(1000)
	NewAcc.GetBalance()
	/*
		{VivianChow 101 88888888} 账号创建成功！
		您的账号是： VivianChow
		您的余额是： 101
		您的密码是： 88888888
		余额修改成功！
		您的余额是： 1000
	*/

	// 练习二
	NewPer := personModule.PersonCreat("Vian")
	fmt.Println(NewPer.Name)
	NewPer.SetAge(18)
	NewPer.SetPalary(3000)
	NewPer.GetAge()
	NewPer.GetPalary()
	/*
		Vian
		年龄修改成功！
		工资修改成功！
		您的年龄是： 18
		您的工资是: 3000
	*/
}
