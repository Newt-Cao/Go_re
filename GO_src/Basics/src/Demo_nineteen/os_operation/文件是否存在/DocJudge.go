package main

import (
	"fmt"
	"os"
)

// 判断文件是否存在的函数，参数为路径或者文件名
func FileJudge(file string) (bool, error) {
	// 使用os.Stat()函数，返回两个值，第一个是该文件参数，第二个是错误判断
	data, err := os.Stat(file)
	// err==nil时，说明文件存在，返回true，nil
	if err == nil {
		fmt.Println(data)
		//&{abc.txt 32 {2002871072 30970178} {2999585239 30970281} {2999005197 30970281} 0 70 0 0 {0 0} D:\Study\Programming\GoCode\src\vivianchan.cn\Basics\src\Demo_nineteen\os\abc.txt 0 0 0 false}
		return true, nil
		// os.IsNotExist(err) 判断err为true说明文件不存在返回false，nil
	} else if os.IsNotExist(err) {
		return false, nil
	}
	// 其他错误情况，返回false，和错误类型err
	return false, err
}

func main() {
	judge, err := FileJudge("D:/Study/Programming/GoCode/src/vivianchan.cn/Basics/src/Demo_nineteen/os/文件的读写/abc.txt")
	fmt.Println(judge, err) // true <nil>
}
