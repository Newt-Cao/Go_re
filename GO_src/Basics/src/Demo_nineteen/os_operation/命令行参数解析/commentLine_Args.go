package main

import (
	"fmt"
	"os"
)

func main() {
	// 自动获取命令行参数
	for i, v := range os.Args {
		fmt.Printf("arg[%v]=%v\n", i, v)
	}
}
