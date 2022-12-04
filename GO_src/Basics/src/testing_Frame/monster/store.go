package monster

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// 声明结构体 Monster
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

// 绑定 Store 方法
func (mon *Monster) Store() bool {
	// 序列化 Monster 对象
	date, err := json.Marshal(mon)
	if err != nil {
		fmt.Println("Json serialize err = ", err)
		return false
	}
	// 文件路径
	fileWritePath := "./test_Case.txt"
	//打开一个文件，如果不存在则创建
	file, err := os.OpenFile(fileWritePath, os.O_CREATE|os.O_WRONLY, 0660)
	if err != nil {
		fmt.Println("Open file err = ", err)
		return false
	}
	// 及时关闭文件
	defer file.Close()
	// 创建带缓存的文件写对象
	writer := bufio.NewWriter(file)
	// 把内容写入到缓存，date为[]byte切片，需要转化
	writer.WriteString(string(date))
	// 最后写入到文件
	writer.Flush()

	/*
		也可以使用io包的方法
		err = ioutil.WriteFile(fileWritePath,date,0660)
		if err != nil {
		fmt.Println("Write err = ",err)
		return false
		}
	*/

	return true
}

// 绑定 ReStore 方法
func (mon *Monster) ReStore() bool {
	// 文件路径
	fileReadPath := "./test_Case.txt"
	// 读取文件，刚好返回的是一个 []byte 切片，不需要关闭
	date, err := ioutil.ReadFile(fileReadPath)
	if err != nil {
		fmt.Println("Open file err =", err)
		return false
	}
	// 传入file，[]byte切片和结构体对象
	err = json.Unmarshal(date, &mon)
	if err != nil {
		fmt.Println("Json Unmarshal err = ", err)
		return false
	}

	return true
}
