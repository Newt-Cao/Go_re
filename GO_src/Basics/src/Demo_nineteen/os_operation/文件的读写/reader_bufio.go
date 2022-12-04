package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开文件
	filename, err := os.Open("D:/Study/Programming/GoCode/src/vivianchan.cn/Basics/src/Demo_nineteen/os/文件的读取/test.txt")
	if err == io.EOF {
		fmt.Println("Open err = ", err)
	}

	// 程序结束时，需要及时关闭文件，防止内存泄漏
	defer filename.Close()

	// 创建一个带缓冲的读取对象，*Reader，默认缓存4096个字节，这样可以读完指定内容，然后再往下读，一点一点读取
	reader := bufio.NewReader(filename)

	// 因为是按照规定一点一点读取，所以需要不断遍历
	for {
		// 这里是每读取到第一次 \n 返回一次后，再往下读到下一个 \n 返回
		str, err := reader.ReadString('\n')
		// io.EOF指读取到文件末尾，当err等于文件末尾，说明文件读取没有出错，break结束循环，当遇到错误时，回返回错误之前读取的数据以及该错误
		if err == io.EOF {
			break
		}
		// 读取无错误，则打印读取内容
		fmt.Print(str)
	}
}

/*
Study golang!
go to school!
mather fuck!
aaaaaaaaa!
*/
