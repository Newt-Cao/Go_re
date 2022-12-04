package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// 声明两个文件路径
	filePath := "./readA.txt"
	filePath2 := "./writerB.txt"
	// 打开第一个文件
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0660)
	if err != nil {
		fmt.Println("Open err = ", err)
		return
	}
	// 及时关闭文件
	defer file.Close()
	// 向第一个文件写入 Hello World！
	str := "Hello World!"
	writer := bufio.NewWriter(file)
	writer.WriteString(str)
	writer.Flush()
	// 一次性读取第一个文件 readA.txt
	content, err := ioutil.ReadFile(filePath)
	if err == io.EOF {
		return
	}
	// 一次性写入文件第二个文件 writerB.txt
	ioutil.WriteFile(filePath2, content, 0660)
	if err != nil {
		fmt.Println("Write err = ", err)
	}
}
