/*
	案例三：
		1.开一个协程writeDataToFile，随机生成1000个数据，写入到文件中。
		2.当writeDataToFile完成操作后，让readDataFromFile协程从文件中读取1000个文件，并完成排序，重新写入到另一个文件中。
		3.使用协程、管道和文件完成。
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

// 开一个协程writeDataToFile，随机生成1000个数据，写入到文件中。
func writeDataToFile(filePath string, exitChan chan int) {
	// 生成随机1-1000随机数
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 1000; i++ {
		Data := rand.Intn(1000) + 1
		// 将随机数转换成字符串才可以写入文件
		str := strconv.Itoa(Data)
		// 打开文件，没有就创建
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND, 0660)
		if err != nil {
			fmt.Println("File Open err = ", err)
		}
		// 函数执行结束及时关闭文件
		defer file.Close()
		// 创建一个写入缓冲
		Writer := bufio.NewWriter(file)
		// 写入数据，并回车，更加清晰
		Writer.WriteString(str + "\n")
		Writer.Flush()
		// 让readDataFromFile等待写入数据完成
		exitChan <- 1
	}
}

// 让readDataFromFile协程从文件中读取1000个文件
func readDataFromFile(filePath string, readChan chan []byte, exitChan chan int) {
	// 等待写入数据完成后才可以进行读取操作
	for i := 1; i <= 1000; i++ {
		<-exitChan
	}
	// 打开写好的文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("File open err = ", err)
	}
	// 函数执行结束及时关闭文件
	defer file.Close()
	// 创建一个读取缓冲
	Reader := bufio.NewReader(file)
	// 死循环，读取文件内容，当读取到结尾时结束
	for {
		// Reader.ReadLine()读取时不读转义符，如：\n \t \r，并返回一个[]byte
		Data, _, err := Reader.ReadLine()
		if err == io.EOF {
			break
		}
		// 把读取内容发送到管道
		readChan <- Data
	}
	// 管道使用结束后关闭，方便判断
	close(readChan)
}

// 对读取的管道数据进行排序
func sortDataToFile(filePath2 string, intSlice []int, readChan chan []byte, exitChan2 chan int) {
	// 接收管道数据
	for {
		// Data的类型是[]byte
		Data, err := <-readChan
		if !err {
			break
		}
		// 将管道数据[]bye类型转换成字符串
		str := fmt.Sprintf("%s", Data)
		// 再把字符串转换成整型，方便排序，因为上面已声明err，所以用errr区分
		num, errr := strconv.Atoi(str)
		if errr != nil {
			fmt.Println("Swap err = ", errr)
		}
		// 把转换好的整型添加到整型切片中
		intSlice = append(intSlice, num)
	}
	// 使用sort包对intSlice进行排序，因为是引用传递，所以会改变intSlice的值
	sort.Ints(intSlice)
	// 创建一个新文件用于接收排序后的数据
	file, err := os.OpenFile(filePath2, os.O_CREATE|os.O_APPEND, 0660)
	if err != nil {
		fmt.Println("Open file err = ", err)
	}
	// 及时关闭文件
	defer file.Close()
	// 遍历切片，取出它的值
	for _, v := range intSlice {
		// 因为是int类型，所以需要转换成string后才可以写入文件
		str := strconv.Itoa(v)
		// 创建一个写入缓冲
		Writer := bufio.NewWriter(file)
		// 写入数据
		Writer.WriteString(str + "\n")
		err = Writer.Flush()
		if err != nil {
			fmt.Println("Write err = ", err)
		}
		// 用于让主程序等待子程序执行
		exitChan2 <- 1
	}
	// 函数运行结束，则关闭通道，让主程序知道子程序已经结束
	close(exitChan2)
}

func main() {
	// 创建两个通信管道，分别是 writeDataToFile 和 readDataFromFile 通信，sortDataToFile 和 主程序 main 通信
	exitChan := make(chan int, 1000)
	exitChan2 := make(chan int, 1000)
	// 创建一个读写数据通道，是 readDataFromFile 和 sortDataToFile 进行通信
	readChan := make(chan []byte, 1000)
	// 创建两个路径，分别是第一次写入随机数据文件，和最后排序后写入数据文件
	filePath := "./Data.txt"
	filePath2 := "./SortData.txt"
	// 声明一个整型切片用于对数据进行排序
	intSlice := []int{}
	// 启动三个协程
	go writeDataToFile(filePath, exitChan)
	go readDataFromFile(filePath, readChan, exitChan)
	go sortDataToFile(filePath2, intSlice, readChan, exitChan2)
	// 判断sortDataToFile是否结束，主程序才结束
	for {
		_, err := <-exitChan2
		if !err {
			break
		}
	}
}
