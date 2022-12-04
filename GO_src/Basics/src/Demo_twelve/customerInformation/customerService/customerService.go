package customerService

import (
	"customerModule"
	"fmt"
)

// 实例化一个Module结构体，用于储存单个信息
var customers customerModule.Customers

type customerSlice struct {
	// 储存多个用户信息
	customerSlices []customerModule.Customers
	// 读取切片有多少个用户
	customerSlicesNum int
}

// Service工厂模式
func FactoryService() *customerSlice {
	return &customerSlice{}
}

// 添加用户信息到切片中
func (cus *customerSlice) Add() bool {
	// 实例化一个Customers用户对象,调用了Module方法
	newUser := customers.UserAdd()
	// 然后添加到切片中，每实例一个对象就添加到切片中
	cus.customerSlices = append(cus.customerSlices, newUser)
	// 添加成功
	return true
}

// 用户信息修改
func (cus *customerSlice) Update() bool {
	// 判断切片中是否有用户信息
	if len(cus.customerSlices) == 0 {
		fmt.Println("\t目前无任何用户信息！赶快添加一个吧！")
		// 无用户信息返回false,用于customerView判断
		return false
	}

	// 用户输入
	userInput := 0
	fmt.Print("\t输入需要修改用户的ID（回车表示该项不修改）：")
	fmt.Scanln(&userInput)
	// 遍历查找id
	for i := 0; i < len(cus.customerSlices); i++ {
		if userInput == cus.customerSlices[i].Id {
			// 获取id对应切片的索引后调用方法修改
			cus.customerSlices[i].UserRevamp()
			return true
		} else {
			fmt.Println("\n\t查找不到该用户信息！")
			return false
		}
	}
	return false
}

// 用户信息删除
func (cus *customerSlice) Delete() bool {
	// 判断切片中是否有用户信息
	if len(cus.customerSlices) == 0 {
		fmt.Println("\t目前无任何用户信息！赶快添加一个吧！")
		// 无用户信息返回false,用于customerView判断
		return false
	}

	// 用户输入
	userInput := 0
	fmt.Print("\t请输入你要删除的用户信息的ID：")
	fmt.Scanln(&userInput)
	// 遍历查找
	for i := 0; i < len(cus.customerSlices); i++ {
		if userInput == cus.customerSlices[i].Id {
			// 判断用户选择
			userChoose := ""
			fmt.Print("\t是否删除Id为", userInput, "用户的信息（y/n）？")
			fmt.Scanln(&userChoose)
			if userChoose == "y" {
				// 把切片开头到用户输入id的索引截止，不包含用户输入，和用户输入后一位到最后添加到新切片中
				cus.customerSlices = append(cus.customerSlices[:i], cus.customerSlices[i+1:]...)
				return true
			} else if userChoose == "n" {
				return false
			} else {
				fmt.Println("\t输入有误！")
				break
			}
		}
	}
	fmt.Println("\n\t查找不到该用户信息！")
	return false
}

// 用户列表展示
func (cus *customerSlice) List() bool {
	// 判断切片中是否有用户信息
	if len(cus.customerSlices) == 0 {
		fmt.Println("\n\t目前无任何用户信息！赶快添加一个吧！")
		// 无用户信息返回false,用于customerView判断
		return false
	}

	fmt.Println("\tId\t姓名\t年龄\t性别\t电话")
	// 遍历切片列表
	for i := 0; i < len(cus.customerSlices); i++ {
		// cus.customerSlices[i].UserShow() 去切片的第 i 个索引的值（每一个元素是Module里的Customers结构体实例），然后调用UserShow()方法
		fmt.Printf("\n\t%s", cus.customerSlices[i].UserShow())
	}
	// 有用户信息返回true,用于customerView判断
	return true
}

// 软件退出
func (cus *customerSlice) Exit() bool {
	// 用户输入
	userChoose := ""
	fmt.Print("\n\t是否确定退出用户信息管理软件（y/n）？")
	for {
		fmt.Scanln(&userChoose)
		switch userChoose {
		case "y":
			return true
			break
		case "n":
			return false
			break
		default:
			fmt.Print("\n\t输入有误请重新输入：")
			continue
		}
	}
}
