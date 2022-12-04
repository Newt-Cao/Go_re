package main

import (
	"errors"
	"fmt"
)

// 判断文件名是否传入正确，是则执行下面代码，否则终止，并抛出指定自定义错误。
func readConfig(name string) error {
	if name == "Config.ini" {
		return nil
	} else {
		return errors.New("文件读取错误...")
	}
}

// 判断是否终止程序。
func stopRead(err error) {

	// 使用处理机制
	/* defer func() {
		errs := recover()
		fmt.Println(errs)
	}() */

	if err != nil {
		panic(err)
	} else {
		fmt.Println("文件读取正确，继续执行代码...")
	}
}

func main() {
	stopRead(readConfig("2Config.ini"))
	fmt.Println("代码")
}
