package main

import (
	"fmt"
)

// 接口
type UsbInterface interface {
	Start()
	Stop()
}

type Phone struct {
	Name string
}

type Camera struct {
	Name string
}

// 两个结构体分别实现接口方法
func (phone *Phone) Start() {
	fmt.Println("手机连接中...")
}

func (phone *Phone) Stop() {
	fmt.Println("手机断开中...")
}

func (camera *Camera) Start() {
	fmt.Println("相机连接中...")
}

func (camera *Camera) Stop() {
	fmt.Println("相机断开中...")
}

func Working(usb UsbInterface) {
	usb.Start()
	usb.Stop()
}

func main() {
	var usb [2]UsbInterface = [2]UsbInterface{
		&Phone{"索尼"},
		&Camera{"三星"},
	}
	usb[0].Start()
	usb[1].Stop()
	fmt.Println(usb[0])
	/*
		手机连接中...
		相机断开中...
		&{索尼}
	*/
}
