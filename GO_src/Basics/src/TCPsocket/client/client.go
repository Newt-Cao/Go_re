package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	// 先连接到服务器，指定服务器类型和IP:端口
	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("Server connected error = ", err)
		return
	}
	// 获取终端输入，os.Stdin代表标准输入，这里指键盘输入到终端
	Reader := bufio.NewReader(os.Stdin)
	// 循环读取终端输入
	for {
		content, err := Reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("Content gain over!")
		}
		// 输入exit时退出，这里去掉了\n用于判断比较，需要在下面发送时带上
		content = strings.Trim(content, " \n\r")
		if content == "exit" {
			fmt.Println("客户端已退出！")
			break
		}
		// 将数据发送给服务器，即把读到的数据写入连接里，n是[]byte的长度，error是写入失败的错误
		_, err = conn.Write([]byte(content + "\n"))
		if err != nil {
			fmt.Println("Date send err = ", err)
		}
	}
}
