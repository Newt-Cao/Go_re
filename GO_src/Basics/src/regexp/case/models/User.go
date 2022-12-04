/*
*****************************************************

	Author       : Mr.Chan
	Date         : 2022-11-05 10:38:24
	LastEditors  : Mr.Chan
	LastEditTime : 2022-11-05 10:38:25
	FilePath     : /vivianchan.cn/Basics/src/regexp/case/module/User.go

*****************************************************
*/
package models

import "github.com/astaxie/beego/orm"

// 接收数据库返回的数据和用于生成表
type User_csv struct {
	Uid      int    `json:"uid" orm:"auto;column(uid);size(10);pk"`
	Name     string `json:"name" orm:"size(30);column(name);null"`
	Password string `json:"password" orm:"size(100);column(password);null`
}

//注册模型
func init() {
	orm.RegisterModel(new(User_csv))
}
