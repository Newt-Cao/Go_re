/*
	案例四：
		1.创建一个结构体Person，包含字段Name、Age、Address。
		2.使用rand方法配合随机创建10个Person实例，并放入到channel中。
		3.遍历channel，将各个Person实例的信息显示在终端。
*/

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// 结构体
type Person struct {
	// 统计几个Person实例
	No int

	Name    string
	Age     int
	Address string
}

// 随机生成姓名：Xxxx，传入参数分别是：姓名长度、组成姓名的字符
func randomChars(n int, allowedChars string) string {
	// 开辟切片空间，给x开辟n个储存空间去存放单个字符
	x := make([]byte, n)
	// 遍历x，把每个x里的空间被赋值上随机的a-z的Unicode编码
	for i := range x {
		// 随机传入allowedChars索引对应值的Unicode编码赋值给x
		x[i] = allowedChars[rand.Intn(len(allowedChars))]
	}
	// 因为人姓名开头是大写，所以需要转换，先把Unicode的[]rune强转换成字符串
	runeTostring := string(x)
	// 再利用字符串拼接，把第一个字母先大写后再和后面拼接，最后赋值给一个新的变量
	runeTostringres := strings.ToUpper(runeTostring[:1]) + runeTostring[1:]
	return runeTostringres
}

// 随机生成地址：传入参数是指定随机生成地址的范围
func randomStrings(allowedStrings []string) string {
	// 指定生成一个地址相当于 x := "" ，写成切片方便遍历，这里默认生成一个地址
	x := [...]string{""}
	// 遍历，实际上只有一个值可以遍历，并把随机地址复制给x
	for i := range x {
		// 把随机索引对应的值传给x，利用随机索引达到随机效果
		x[i] = allowedStrings[rand.Intn(len(allowedStrings))]
	}
	return x[0]
}

func main() {
	// 生成不固定的种子
	rand.Seed(time.Now().UnixNano())
	// 规定人名的字符，名字如果是中文，需要用[]rune()进行转换，[]rune()是按单个字符处理，[]byte()是按单个字节处理，所以是因为时可以用，又因为字符串就是byte拼接所以不需要转换
	randomCharsCreate := "abcdefghijklmnopqrstuvwxyz"
	// 规定随机地址范围
	randomStringCreate := []string{"上海", "北京", "广东", "湖南", "浙江", "广西", "云南", "江西", "湖北", "山东"}
	// 声明管道负责接收和发送
	personChan := make(chan []Person, 10)
	// 生成10个实例
	for i := 1; i <= 10; i++ {
		// 用于接收所有实例
		personSlice := []Person{}
		person := Person{
			No:      i,
			Name:    randomChars(4, randomCharsCreate),
			Age:     rand.Intn(100) + 1,
			Address: randomStrings(randomStringCreate),
		}
		// 把生成好的实例添加到切片中
		personSlice = append(personSlice, person)
		// 最后把切片传入管道
		personChan <- personSlice
	}
	// 关闭管道
	close(personChan)
	for {
		// 接收切片
		Data, ok := <-personChan
		if !ok {
			break
		}
		// 遍历切片和打印
		for _, v := range Data {
			fmt.Printf("第 %v 个Person：\n\n\t\t姓名：%v\n\t\t年龄：%v\n\t\t地址：%v\n\n", v.No, v.Name, v.Age, v.Address)
		}
	}
}
