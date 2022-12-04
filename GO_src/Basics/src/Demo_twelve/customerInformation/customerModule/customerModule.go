package customerModule

import "fmt"

type Customers struct {
	Id        int
	name      string
	age       int
	gender    string
	telephone string
}

// 用户信息添加(相当于Module工厂模式)，实例化一个Customers对象
func (cus *Customers) UserAdd() Customers {
	for {
		fmt.Print("\t姓名：")
		fmt.Scanln(&cus.name)
		if len(cus.name) < 4 || len(cus.name) > 30 {
			fmt.Println("\t姓名长度错误！请重新输入！")
			continue
		}

	Age:
		fmt.Print("\t年龄：")
		fmt.Scanln(&cus.age)
		if cus.age <= 0 || cus.age >= 200 {
			fmt.Println("\t年龄输入有误！请重新输入！")
			goto Age
		}

	Gender:
		fmt.Print("\t性别：")
		fmt.Scanln(&cus.gender)
		if cus.gender == "男" || cus.gender == "女" {
			goto Tele
		} else {
			fmt.Println("\t性别输入有误！请重新输入！")
			goto Gender
		}

	Tele:
		fmt.Print("\t电话：")
		fmt.Scanln(&cus.telephone)
		if len(cus.telephone) <= 0 || len(cus.telephone) > 11 {
			fmt.Println("\t电话输入有误！请重新输入！")
			goto Tele
		}

		cus.Id++

		return Customers{
			Id:        cus.Id,
			name:      cus.name,
			age:       cus.age,
			gender:    cus.gender,
			telephone: cus.telephone,
		}
	}
}

// 用户信息展示
func (cus *Customers) UserShow() string {
	return fmt.Sprintf("%v\t%v\t%v\t%v\t%v\n", cus.Id, cus.name, cus.age, cus.gender, cus.telephone)
}

// 用户信息修改
func (cus *Customers) UserRevamp() Customers {
	for {
		fmt.Print("\t姓名：")
		fmt.Scanln(&cus.name)
		if len(cus.name) < 4 || len(cus.name) > 30 {
			fmt.Println("\t姓名长度错误！请重新输入！")
			continue
		}

	Age:
		fmt.Print("\t年龄：")
		fmt.Scanln(&cus.age)
		if cus.age <= 0 || cus.age >= 200 {
			fmt.Println("\t年龄输入有误！请重新输入！")
			goto Age
		}

	Gender:
		fmt.Print("\t性别：")
		fmt.Scanln(&cus.gender)
		if cus.gender == "男" || cus.gender == "女" {
			goto Tele
		} else {
			fmt.Println("\t性别输入有误！请重新输入！")
			goto Gender
		}

	Tele:
		fmt.Print("\t电话：")
		fmt.Scanln(&cus.telephone)
		if len(cus.telephone) <= 0 || len(cus.telephone) > 11 {
			fmt.Println("\t电话输入有误！请重新输入！")
			goto Tele
		}

		return Customers{
			name:      cus.name,
			age:       cus.age,
			gender:    cus.gender,
			telephone: cus.telephone,
		}
	}
}
