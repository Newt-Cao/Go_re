// 面向过程

package main

import "fmt"

func main() {
	// 用户输入数字
	keyNum := 0
	// 循环判断是否退出软件
	close := true
	// 判断是否有收入或支出行为
	flag := false
	// 用户余额
	balance := 0.0
	//收入/支出详情，结合功能 1 实现
	detail := ""

	for {
		fmt.Println()
		fmt.Println()

		fmt.Println(`---------------------家庭收支记账软件---------------------
										
		1.收入明细
		2.登记收入
		3.登记支出
		4.退出
		
		`)
		fmt.Print("请选择（1-4）：")
		fmt.Scanln(&keyNum)

		fmt.Println()
		fmt.Println()

		switch keyNum {
		case 1:
			fmt.Println(`---------------------当前收支明细记录---------------------
收    支        账户金额        收支金额        说    明`)
			if flag {
				fmt.Printf(detail)
			} else {
				fmt.Println("目前无任何记录！")
				break
			}

		case 2:
			// 单次金额收入
			saveMoney := 0.0
			// 金额说明数据
			saveData := ""

			fmt.Print("输入本次收入金额：")
			fmt.Scanln(&saveMoney)
			// 收入的钱不能是负数
			if saveMoney < 0 {
				fmt.Println("收入记录错误！")
				break
			} else {
				balance += saveMoney
			}

			fmt.Print("输入本次收入说明：")
			fmt.Scanln(&saveData)
			detail += fmt.Sprintf("\n收    入       %v              %v              %s", balance, saveMoney, saveData)

			flag = true

		case 3:
			//单次金额使用
			spendMoney := 0.0
			// 金额说明数据
			spendData := ""

			fmt.Print("输入本次支出金额：")
			fmt.Scanln(&spendMoney)
			// 支出的钱不可大余额
			if spendMoney > balance {
				fmt.Println("余额不足！")
				break
			} else {
				balance -= spendMoney
			}

			fmt.Print("输入本次支出说明：")
			fmt.Scanln(&spendData)
			detail += fmt.Sprintf("\n支    出        %v               %v              %s", balance, spendMoney, spendData)

			flag = true

		case 4:
			// y/n的输入
			decide := ""
			fmt.Print("确定退出软件吗？（y/n）")

			for {
				fmt.Scanln(&decide)
				if decide == "y" || decide == "n" {
					break
				} else {
					fmt.Println("输入有误，请重新输入：")
				}
			}

			if decide == "y" {
				fmt.Println("软件已退出！")
				close = false
			} else if decide == "n" {
				close = true
			}

		default:
			fmt.Println("请输入正确号码！")
			close = true
		}

		if !close {
			break
		}
	}
}
