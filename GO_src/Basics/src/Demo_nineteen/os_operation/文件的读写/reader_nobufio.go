package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 一次性读取文件吗，不需要打开和关闭
	content, err := ioutil.ReadFile("../test.txt")

	// 判断是否出错
	if err != nil {
		fmt.Println("err = ", err)
	}

	// 打印输出读取内容,因为返回的是一个字节切片 []byte ，元素为编码形式，所以需要转换
	fmt.Printf("%s", content)
	
	/*
		Study golang!
		go to school!
		mather fuck!
		aaaaaaaaa!

	*/
}
