package customerView

import (
	"customerService"
	"fmt"
)

type userControl struct {
	// 控制用户输入数字功能
	keyNum string
	// 控制程序是否退出
	judge bool
	// 控制用户是否添加成功
	useradd bool
	// 控制用户列表展示
	usershow bool
	// 控制用户信息修改
	userrevamp bool
	// 控制用户信息删除
	userdelete bool
}

// 工厂模式
func FactoryView() *userControl {
	return &userControl{}
}

func (con *userControl) View() {
	// 实例化一个Service对象
	var customerservice = customerService.FactoryService()

	for {
		fmt.Println(`
	------------用户信息管理软件-----------

	1.添加用户信息
	2.修改用户信息
	3.删除用户信息
	4.用户信息列表
	5.管理软件退出
	`)
		fmt.Print("\t请输入（1-5）：")
		fmt.Scanln(&con.keyNum)
		switch con.keyNum {
		case "1":
			fmt.Println("\n\t------------用户信息添加------------")
			// 返回布尔值
			con.useradd = customerservice.Add()
			// 判断布尔值true则成功
			if con.useradd {
				fmt.Println("\n\t------------信息添加成功------------")
			}
		case "2":
			fmt.Println("\n\t------------用户信息修改------------\n")
			con.userrevamp = customerservice.Update()
			if con.userrevamp {
				fmt.Println("\n\t------------信息修改成功------------")
			} else {
				fmt.Println("\n\t------------信息修改失败------------")
			}
		case "3":
			fmt.Println("\n\t------------用户信息删除------------\n")
			con.userdelete = customerservice.Delete()
			if con.userdelete {
				fmt.Println("\n\t------------信息删除成功------------\n")
			} else {
				fmt.Println("\n\t------------信息删除失败------------\n")
			}
		case "4":
			fmt.Println("\n\t------------用户信息列表------------")
			// 返回布尔值
			con.usershow = customerservice.List()
			// fales失败跳出switch判断
			if con.usershow {
				fmt.Println("\n\t------------信息展示成功------------\n")
			} else {
				fmt.Println("\n\t------------信息展示失败------------\n")
			}
		case "5":
			con.judge = customerservice.Exit()
		default:
			fmt.Print("\n\t输入有误请重新输入!\n")
		}
		if con.judge {
			break
		}
	}
	fmt.Println("\n\t程序已退出！")
}

func main() {
	var contorl = userControl{}
	contorl.View()
}
