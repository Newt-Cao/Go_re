package main //程序的包名

/* 导包
import "fmt"
*/
import (
	"fmt"
	"time"
)

// 创建main函数
func main() { //左大括号必须与函数名在同一行
	//golang中的分号 ";" 可加可不加
	fmt.Println("你好,世界!")
	time.Sleep(1 * time.Second)
}
