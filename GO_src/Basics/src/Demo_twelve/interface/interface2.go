package main

import "fmt"

// interface{}表示万能类型
func MyFunc(arg interface{}) {
	fmt.Println("MyFunc is called...")
	fmt.Println(arg)

	// 如何区分数据类型（类型断言）机制，必须是interface{}万能类型才可以使用
	// 变量名.(type) 返回两个值: value bool
	value, ok := arg.(string)
	fmt.Println(value, ok)
	// ok == false
	if !ok {
		fmt.Printf("%v is not string\n", arg)
	} else {
		fmt.Printf("%v is string\n", arg)
	}
}

type Human struct {
	Name string
}

func main() {
	human := Human{"张三"}
	MyFunc(human)
	MyFunc("张三")
	MyFunc(100)
	MyFunc(3.14)
	MyFunc(true)
}
