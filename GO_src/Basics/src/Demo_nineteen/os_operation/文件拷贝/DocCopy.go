package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 封装成一个函数
func ImageCopy(dst_Image, src_Image string) (written int64, err error) {

	// dst Writer实现，打开拷贝后照片
	fileCopy, err := os.OpenFile(dst_Image, os.O_CREATE|os.O_WRONLY, 0660)
	if err != nil {
		fmt.Println("Image write err = ", err)
	}
	// 关闭文件资源
	defer fileCopy.Close()
	// 把文件写入到缓存
	Writer := bufio.NewWriter(fileCopy)

	// src Reader实现，打开被拷贝的照片
	file, err := os.Open(src_Image)
	if err != nil {
		fmt.Println("Image open err = ", err)
	}
	// 关闭文件资源
	defer file.Close()
	// 读取文件到缓存
	Reader := bufio.NewReader(file)

	// 调用 io 包的 Copy 函数
	return io.Copy(Writer, Reader)
}

func main() {
	// 原照片的路径 D:/Study/Resuource/Image/ins周慧敏图集/mmexport1634384853655.jpg
	imagePath := "D:/Study/Resuource/Image/ins周慧敏图集/mmexport1634384853655.jpg"
	// 拷贝照片后的路径 D:/Study/Programming/GoCode/src/vivianchan.cn/Basics/src/Demo_nineteen/os_operation/VivianChow.jpg
	imageCopyPath := "D:/Study/Programming/GoCode/src/vivianchan.cn/Basics/src/Demo_nineteen/os_operation/文件拷贝/VivianChow.jpg"
	// 函数调用
	_, err := ImageCopy(imageCopyPath, imagePath)
	if err != nil {
		fmt.Println("Copy err = ", err)
	} else {
		fmt.Println("Copy success!")
	}
}
