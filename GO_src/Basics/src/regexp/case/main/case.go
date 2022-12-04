/*
*****************************************************

	Author       : Mr.Chan
	Date         : 2022-11-05 10:17:41
	LastEditors  : Mr.Chan
	LastEditTime : 2022-11-05 10:17:42
	FilePath     : /vivianchan.cn/Basics/src/正则表达式/case.go

*****************************************************
*/
package main

import (
	"controllers"
	"encoding/csv"
	"log"
	"models"
	"os"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 初始化数据库连接
func init() {
	orm.RegisterDataBase("default", "mysql", "root:112233@tcp(localhost:3306)/vian?charset=utf8mb4&parseTime=true&loc=Local") // 注册并连接数据库

	orm.RunSyncdb("default", false, true) // 创建表，不强制更新，且输出sql语句
	orm.SetMaxIdleConns("default", 30)    // 最大空闲连接
	orm.SetMaxOpenConns("default", 30)    // 最大连接数
	orm.Debug = true                      // 打开调试，输出sql语句
}

// 编写导出 csv 文件的方法，二维数组用于导出文件的格式
func Export_Csv(filepath string, data [][]string) {
	// 创建文件
	file, err := os.Create(filepath)
	if err != nil {
		log.Fatal("File created error:", err)
	}

	defer file.Close() // 释放资源

	file.WriteString("\xEF\xBB\xBF") // 写入 UTF-8 BOM，否则打开是乱码
	writer := csv.NewWriter(file)    // 创建写入文件流
	writer.WriteAll(data)            // 写入多条数据到缓存，传入的是二维数组
	writer.Flush()                   // 刷新数据到文件输出流，清空缓存
}

func main() {
	// 赋值将要插入的数据
	userInfo := []models.User_csv{
		{Uid: 1001, Name: "张三", Password: "123456"},
		{Uid: 1002, Name: "李四", Password: "1234567"},
	}
	controllers.Cl.Insert_Data(userInfo) // 插入数据

	Query_Info := controllers.Cl.Query_Data() // 查询多条数据

	exportFormat := [][]string{{"Uid", "Name", "Password"}} // 输出格式

	// 转换类型
	for _, v := range Query_Info {
		str := []string{}
		str = append(str, strconv.Itoa(v.Uid))
		str = append(str, v.Name)
		str = append(str, v.Password)
		exportFormat = append(exportFormat, str)
	}

	filepath := "./exportUsers.csv" // 导出文件的路径

	Export_Csv(filepath, exportFormat) // 导出csv文件

	beego.Run() // 运行
}
