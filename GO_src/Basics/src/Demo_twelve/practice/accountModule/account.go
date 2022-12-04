package accountModule

import (
	"errors"
	"fmt"
)

/*
练习一：
	1.声明结构体account，包含字段：账号（长度在6-10之间）、余额（>20）、密码（6位数）。
	2.通过Setxxx方法给Account字段赋值。
	3.通过Getxxx方法查询字段值。
*/

type account struct {
	userAccount string
	balance     float64
	password    string
}

// 实例化对象方法
func AccountClass(useraccount string, userbalance float64, userpassword string) (*account, error) {
	if len(useraccount) < 6 || len(useraccount) > 10 {
		return nil, errors.New("账号长度错误！")
	}
	if userbalance <= 20 {
		return nil, errors.New("余额不足！")
	}
	if len(userpassword) < 6 {
		return nil, errors.New("密码位数错误！")
	}

	return &account{
		userAccount: useraccount,
		balance:     userbalance,
		password:    userpassword,
	}, nil
}

func (acc *account) SetUserAccount(user string) {
	fmt.Println("账号修改成功！")
	(*acc).userAccount = user
}

func (acc *account) SetBalance(userBalance float64) {
	fmt.Println("余额修改成功！")
	(*acc).balance = userBalance
}

func (acc *account) SetPassword(userPassword string) {
	fmt.Println("密码修改成功！")
	(*acc).password = userPassword
}

func (acc *account) GetUserAccount() {
	fmt.Println("您的账号是：",(*acc).userAccount)
}

func (acc *account) GetBalance() {
	fmt.Println("您的余额是：",(*acc).balance)
}

func (acc *account) GetPassword() {
	fmt.Println("您的密码是：",(*acc).password)
}
