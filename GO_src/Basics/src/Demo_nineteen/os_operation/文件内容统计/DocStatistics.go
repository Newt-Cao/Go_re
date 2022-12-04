package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

// 声明统计结构体
type charCount struct {
	// 数字统计
	numCount int
	// 字母统计
	letterCount int
	// 空格统计
	spaceCount int
	// 字符串统计
	stringCount int
	// 其他字符统计
	otherCount int
}

func main() {
	// 实例化一个统计对象
	var statistics charCount
	// 文件路径
	filePath := "./charcount.txt"
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Open err = ", err)
	}
	// 读取文件内容到缓存
	Reader := bufio.NewReader(file)
	// 循环读取
	for {
		content, err := Reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		// 转化为[]rune()类型，读取汉字
		data := []rune(content)
		// 循环读取每个字符
		for _, v := range data {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				statistics.letterCount++
			case v >= '0' && v <= '9':
				statistics.numCount++
			case v == ' ':
				statistics.spaceCount++
			// 使用unicode包的Is函数可以识别文件汉字
			case unicode.Is(unicode.Scripts["Han"], v):
				statistics.stringCount++
			default:
				statistics.otherCount++
			}
		}
	}
	fmt.Printf("整数有 %v 个，字母有 %v 个，空格有 %v 个，汉字有 %v 个，其他字符有 %v 个。", statistics.numCount, statistics.letterCount, statistics.spaceCount, statistics.stringCount, statistics.otherCount)
	// 整数有 4 个，字母有 19 个，空格有 6 个，汉字有 2 个，其他字符有 6 个。
}
