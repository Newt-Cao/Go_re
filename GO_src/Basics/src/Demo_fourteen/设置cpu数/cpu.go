package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 查看本机逻辑CPU个数（线程数）
	cpuNum := runtime.NumCPU()
	fmt.Println("Cpu num = ", cpuNum) // Cpu num =  20
	// 设置程序调用的CPU个数，不设置则使用默认的值
	runtime.GOMAXPROCS(1)
}
