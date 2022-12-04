package main

import (
	"fmt"
)

func main() {
	// map[type]struct类型
	// 声明对象结构体
	type stuDents struct {
		Name    string
		age     int
		address string
		Hobby   []string
	}

	// 实例化对象
	stuD1 := stuDents{
		Name:    "张三",
		age:     20,
		address: "上海",
	}
	stuD1.Hobby = make([]string, 3)
	stuD1.Hobby[0] = "读书"

	stuD2 := stuDents{
		Name:    "李四",
		age:     22,
		address: "广东",
	}
	stuD2.Hobby = make([]string, 3)
	stuD2.Hobby[0] = "看报"

	// 声明映射map
	StuDents := make(map[string]stuDents)
	// 添加
	StuDents["1001"] = stuD1
	StuDents["1002"] = stuD2
	// 遍历
	for key, value := range StuDents {
		fmt.Println("学号：", key)
		fmt.Println("姓名：", value.Name)
		fmt.Println("年龄：", value.age)
		fmt.Println("地址：", value.address)
		fmt.Println("爱好：", value.Hobby[0])
		fmt.Println()
	}
	/*
		学号： 1001
		姓名： 张三
		年龄： 20
		地址： 上海
		爱好： 读书

		学号： 1002
		姓名： 李四
		年龄： 22
		地址： 广东
		爱好： 看报
	*/
}
