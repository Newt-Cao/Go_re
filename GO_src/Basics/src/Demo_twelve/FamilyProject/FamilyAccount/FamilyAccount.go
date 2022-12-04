package FamilyAccount

import "fmt"

// 面向对象
// 家庭账户
type account struct {
	// 选择号码
	keyNum int
	// 标识是否退出和是否有记录
	close bool
	flag  bool
	// 余额
	balance float64
	// 详细页
	detail string
	// 账号
	name     []string
	password []string
}

// 外部调用工厂模式
func Factory() *account {
	return &account{
		name:     []string{"Vian"},
		password: []string{"88888888"},
	}
}

// 主程序
func (acc *account) Software_Core() {
	for {
		// 用户输入
		name := ""
		password := ""

		fmt.Print("请输入账号：")
		fmt.Scanln(&name)
		fmt.Print("请输入密码：")
		fmt.Scanln(&password)
		for _, v := range acc.name {
			for _, v1 := range acc.password {
				if name == v && password == v1 {
					for {
						fmt.Println()
						fmt.Println()

						fmt.Println(`---------------------家庭收支记账软件---------------------
										
		1.收入明细
		2.登记收入
		3.登记支出
		4.账号创建
		5.转账
		6.账号查询
		7.退出
		`)
						fmt.Print("请选择（1-7）：")
						fmt.Scanln(&acc.keyNum)

						fmt.Println()
						fmt.Println()

						switch acc.keyNum {
						case 1:
							acc.income_detail()
						case 2:
							acc.income_registered()
						case 3:
							acc.spend_registered()
						case 4:
							acc.creat_acc()
						case 5:
							acc.transfer_account()
						case 6:
							acc.acc_inquire()
						case 7:
							acc.exit()
						default:
							fmt.Println("请输入正确号码！")
							acc.close = false
						}

						if acc.close {
							break
						}
					}
				} else {
					fmt.Println("账号或密码错误！请重新输入！")
					continue
				}
			}
		}
	}
}

// 1.记录
func (acc *account) income_detail() {
	fmt.Println(`---------------------当前收支明细记录---------------------
收    支        账户金额        收支金额        说    明`)
	if acc.flag {
		fmt.Printf(acc.detail)
	} else {
		fmt.Println("目前无任何记录！")
	}
}

// 2.收入
func (acc *account) income_registered() {
	// 单次金额收入
	saveMoney := 0.0
	// 金额说明数据
	saveData := ""

	for {
		fmt.Print("输入本次收入金额：")
		fmt.Scanln(&saveMoney)
		// 收入的钱不能是负数
		if saveMoney < 0 {
			fmt.Println("收入记录错误！")
			continue
		} else {
			acc.balance += saveMoney
		}

		fmt.Print("输入本次收入说明：")
		fmt.Scanln(&saveData)
		acc.detail += fmt.Sprintf("\n收    入       %v              %v            %s", acc.balance, saveMoney, saveData)

		acc.flag = true
		break
	}
}

// 3.支出
func (acc *account) spend_registered() {
	//单次金额使用
	spendMoney := 0.0
	// 金额说明数据
	spendData := ""

	for {
		fmt.Print("输入本次支出金额：")
		fmt.Scanln(&spendMoney)
		// 支出的钱不可大余额
		if spendMoney > acc.balance {
			fmt.Println("余额不足！")
			break
		} else {
			acc.balance -= spendMoney
		}

		fmt.Print("输入本次支出说明：")
		fmt.Scanln(&spendData)
		acc.detail += fmt.Sprintf("\n支    出        %v               %v            %s", acc.balance, spendMoney, spendData)

		acc.flag = true
		break
	}
}

// 4.账号创建
func (acc *account) creat_acc() {
	// 用户输入
	newname := ""
	newpassword := ""

	for {
		fmt.Print("输入新账号：")
		fmt.Scanln(&newname)
		fmt.Print("输入新密码：")
		fmt.Scanln(&newpassword)

		if len(newname) < 4 || len(newpassword) < 8 {
			fmt.Println("账号或密码长度不够！请重新输入！")
			continue
		} else {
			fmt.Println("账号创建成功！")
			acc.name = append(acc.name, newname)
			acc.password = append(acc.password, newpassword)
			break
		}
	}
	acc.flag = true
}

// 5.转账
func (acc *account) transfer_account() {
	// 用户输入转账账号、金额
	transferAccount := ""
	transferMoney := 0.0
	flag := false

	fmt.Print("输入转账号码：")
	fmt.Scanln(&transferAccount)

	for _, v := range acc.name {
		if transferAccount == v {
			flag = true
			fmt.Print("请输入转账金额：")
			fmt.Scanln(&transferMoney)
			if transferMoney > acc.balance {
				fmt.Println("余额不足！")
				break
			} else {
				acc.balance -= transferMoney
				fmt.Println("转账成功！")
				acc.detail += fmt.Sprintf("\n转    账        %v               %v            ", acc.balance, transferMoney)
				break
			}
		}
	}
	if !flag {
		fmt.Println("查无此账号！")
	}
}

// 6.账号查询
func (acc *account) acc_inquire() {
	fmt.Println(" 账号有：", acc.name, "\n", "密码有：", acc.password, "\n")
}

// 7.退出
func (acc *account) exit() {
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
		acc.close = true
	} else if decide == "n" {
		acc.close = false
	}
}
