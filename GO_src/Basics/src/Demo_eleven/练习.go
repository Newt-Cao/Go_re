package main

import "fmt"

func userSet(user map[string]map[string]string, userName string, nickName string) {
	// 判断是否存在该用户,存在则改密码，不存在就创建
	if user[userName] != nil {
		user[userName]["passWord"] = "88888888"
		user[userName]["nickName"] = nickName
	} else {
		user[userName] = make(map[string]string)
		user[userName]["nickName"] = nickName
		user[userName]["passWord"] = "888888"
	}
}

func main() {
	/*
		map的简单练习：
			存储用户信息，若有该用户，则把密码改成 “888888”，若不存在该用户则创建该用户，密码一样。
	*/
	user := make(map[string]map[string]string)
	user["Niko"] = make(map[string]string)
	user["Niko"]["nickName"] = "沙鹰"
	user["Niko"]["passWord"] = "888888"
	userSet(user, "Niko", "沙鹰")
	fmt.Println(user)
}
