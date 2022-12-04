package main

import (
	"fmt"
	"os"
)

func main() {
	// 打开一个文件
	/*
		filename的几种叫法：
			filename对象
			filename指针
			filename文件句柄
	*/
	filename, err := os.Open("D:/Study/Programming/GoCode/src/vivianchan.cn/Basics/src/Demo_nineteen/os/文件的读写/test.txt")
	if err != nil {
		fmt.Println("Open file err = ", err)
	}

	// 输出查看该文件，文件是一个指针
	fmt.Println(filename) // &{0xc0000d0780}

	// 关闭文件
	err = filename.Close()
	if err != nil {
		fmt.Println("Close file err = ", err)
	}
}
