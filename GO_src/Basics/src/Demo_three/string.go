// 字符串

package main

import (
	"bytes"
	"fmt"
	"strings"
)

type People struct {
	Name string
	age  string
}

func main() {

	h1 := People{"张三", "20"}

	h := strings.Join([]string{h1.Name, h1.age}, " ")

	fmt.Printf("%s\n", h)

	var buffer1 bytes.Buffer
	buffer1.WriteString("李四")
	buffer1.WriteString("|")
	buffer1.WriteString("30")
	fmt.Print(buffer1.String(), "\n")

	h2 := "王五 40"

	//分割成一个数组
	fmt.Printf("strings.Split(h2, \" \"): %v\n", strings.Split(h2, " "))

	//判断该字符串是否包含 王五 40
	fmt.Printf("strings.Contains(h2, \"王五\"): %v\n", strings.Contains(h2, "王五 40"))

	//判断字符串是否为王五开头
	fmt.Printf("strings.HasPrefix(h2, \"王五\"): %v\n", strings.HasPrefix(h2, "王五"))

	//判断字符串是否为40结尾
	fmt.Printf("strings.HasSuffix(h2, \"40\"): %v\n", strings.HasSuffix(h2, "40"))

	a := "HELLO"
	//把字符串转换成小写
	fmt.Printf("strings.ToLower(a): %v\n", strings.ToLower(a))

	b := "hello"
	//转换成大写
	fmt.Printf("strings.ToUpper(b): %v\n", strings.ToUpper(b))

	str := "VivianChan"
	slice := str[6:]
	fmt.Println(slice) //Chan
}
