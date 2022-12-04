/*
*****************************************************

	Author       : Mr.Chan
	Date         : 2022-11-06 08:28:19
	LastEditors  : Mr.Chan
	LastEditTime : 2022-11-06 08:28:19
	FilePath     : /vivianchan.cn/Basics/src/log/new.go

*****************************************************
*/
package main

import (
	"log"
	"os"
)

func main() {
	filePath := "./new.log"
	logFile, err := os.Create(filePath) // 创建文件
	if err != nil {
		log.Fatal("Could not create log file , error:", err)
	}

	debugLogger := log.New(logFile, "[Logger]", log.Lshortfile) // 设置日志对象模式，并指定输出文件
	debugLogger.Println("Logger test.")                         // 输出的文本
	debugLogger.SetPrefix("[Info]")                             // 更改输出时的 Prefix，即上面的 [Logger]
	debugLogger.SetFlags(log.LstdFlags)                         // 更改输出的 flags
	debugLogger.Println("Info test.")                           // 输出文本
}
