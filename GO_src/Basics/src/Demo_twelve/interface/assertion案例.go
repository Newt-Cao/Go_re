package main

import "fmt"

type Usb interface {
	Start()
}

type Phone struct {
}

type Camera struct {
}

func (phone Phone) Start() {
	fmt.Println("手机开机中...")
}

func (phone Phone) Call() {
	fmt.Println("正在通话...")
}

func (camera Camera) Start() {
	fmt.Println("相机开机中...")
}

func Working(usb Usb) {
	usb.Start()
	// 断言转换成功直接赋值给phone后再调用Call()方法
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
}

func main() {
	var usb [2]Usb = [2]Usb{
		Phone{},
		Camera{},
	}
	for _, v := range usb {
		Working(v)
	}
	/*
		手机开机中...
		正在通话...
		相机开机中...
	*/
}
