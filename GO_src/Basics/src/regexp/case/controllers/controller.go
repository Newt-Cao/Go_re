/*
*****************************************************

	Author       : Mr.Chan
	Date         : 2022-11-05 10:45:03
	LastEditors  : Mr.Chan
	LastEditTime : 2022-11-05 10:45:04
	FilePath     : /vivianchan.cn/Basics/src/regexp/case/controllers/controller.go

*****************************************************
*/
package controllers

import (
	"fmt"
	"log"
	"models"

	"github.com/astaxie/beego/orm"
)

type controller struct{}

var Cl *controller

// 批量插入数据
func (c *controller) Insert_Data(data []models.User_csv) string {
	orm_Obj := orm.NewOrm()                 // 创建 orm 对象，操作数据库
	id, err := orm_Obj.InsertMulti(1, data) // 插入数据
	if err != nil {
		log.Fatal("Insert error: ", err)
	}
	return fmt.Sprintln("插入的数据：", id)
}

// 查询多条数据
func (c *controller) Query_Data() []models.User_csv {
	orm_Obj := orm.NewOrm()
	var userInfo []models.User_csv                                            // 接收数据
	result, err := orm_Obj.Raw("select * from user_csv").QueryRows(&userInfo) // 查询数据

	if err == orm.ErrNoRows {
		fmt.Println("数据不存在。")
		return nil
	} else if err == orm.ErrMissPK {
		fmt.Println("主键不存在")
		return nil
	} else if err != nil {
		fmt.Println(err)
		return nil
	} else {
		fmt.Println("查询总数：", result)
	}

	return userInfo
}
