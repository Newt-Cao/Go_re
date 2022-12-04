package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 读写方式打开，先读取文件内容，再追加
	filename, err := os.OpenFile("./abc.txt", os.O_RDWR, 4)
	if err != nil {
		fmt.Println("Open err = ", err)
	}

	// 最后执行关闭文件
	defer filename.Close()

	// 读取文件缓存
	reader := bufio.NewReader(filename)
	for {
		content, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(content)
	}

	// 写入5句 Hello World！到 abc.txt 文件里
	str := "Hello S1mple!\n"
	// 创建带缓存的对象
	writer := bufio.NewWriter(filename)
	for i := 0; i < 5; i++ {
		// 先写入到缓存里，默认缓存4096个字节
		writer.WriteString(str)
	}
	// 再写入到文件
	writer.Flush()
}
