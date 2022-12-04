package main

import (
	"flag"
	"fmt"
)

// 接收值声明
type CommentVar struct {
	// 用户名
	user string
	// 密码
	password string
	// host地址
	host string
	// 端口
	port int
}

func main() {
	// 创建接收值实例
	var commentvar CommentVar
	// 调用 StringVar()、IntVar()
	flag.StringVar(&commentvar.user, "u", "", "默认空值")
	flag.StringVar(&commentvar.password, "pwd", "", "默认空值")
	flag.StringVar(&commentvar.host, "h", "localhost", "默认localhost")
	flag.IntVar(&commentvar.port, "port", 3306, "默认端口：3306")
	// 转换
	flag.Parse()
	// 输出
	fmt.Printf("账号：%v 密码：%v 主机地址：%v 端口：%v", commentvar.user, commentvar.password, commentvar.host, commentvar.port)
}
