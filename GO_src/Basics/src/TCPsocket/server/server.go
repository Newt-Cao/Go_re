package main

import (
	"fmt"
	"net"
)

// 为每一个客户端开辟一个协程处理数据
func progress(conn net.Conn) {
	// 不用时要关闭连接，防止连接过多，服务器崩溃
	defer conn.Close()
	// 提示
	fmt.Printf("%s 请输入消息：\n", conn.RemoteAddr().String())
	// 等待conn写入的信息，并读取显示到终端
	for {
		// Conn.Read需要准备一个字节切片存储数据
		bufioConnRead := make([]byte, 4998)
		// 读取连接发送的数据
		n, err := conn.Read(bufioConnRead)
		if err != nil {
			fmt.Println("Client disconnent!")
			// 客户端退出结束协程
			return
		}
		// n是实际读到的内容长度，但是打印的bufioConnRead一般会大于这个内容，如果不选取范围打印，则会把整一个bufioConnRead打印
		fmt.Print(string(bufioConnRead[:n]))
	}
}

func main() {
	// 提示
	fmt.Println("服务器监听中...")
	// 地址
	laddr := "0.0.0.0"
	// 端口
	port := 8888
	// 监听端口
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", laddr, port))
	// 判断是否监听成功
	if err != nil {
		fmt.Println("Listen err = ", err)
		// 如果建立连接失败，需要重新连接
		return
	} else {
		fmt.Println("网络地址：", listener.Addr())
	}
	// 当收发数据完成时，关闭Accpet
	defer listener.Close()
	// 提示
	fmt.Println("客户端连接中...")
	// 循环等待客户端连接被监听的端口
	for {
		// 等待客户端连接（阻塞）
		conn, err := listener.Accept()
		if err != nil {
			// 一个客户端连接失败不需要把整个桥梁断开，还有别的客户端要连接
			fmt.Println("Client connection fail = ", err)
			continue
		} else {
			// 连接成功，同时可以打出连接方的IP地址
			fmt.Printf("Client connected successfully!\n%s\t已连入服务器！\n", conn.RemoteAddr().String())
		}
		// 给每一个客户端调用协程
		go progress(conn)
	}
}
