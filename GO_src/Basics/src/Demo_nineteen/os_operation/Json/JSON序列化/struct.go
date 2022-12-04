package main

import (
	"encoding/json"
	"fmt"
)

// 结构体
type Hero struct {
	Name  string   `json:"name"`
	Skill []string `json:"skill"`
}

func main() {
	// 实例化结构体对象
	var hero = Hero{
		Name:  "Vian",
		Skill: []string{"飞天", "遁地"},
	}
	// 调用方法
	data, err := json.Marshal(&hero)
	if err != nil {
		fmt.Println("Json err = ", err)
		return
	}
	// 因为返回的是[]byte，所以需要转化，如果没有使用标签，则默认输出字段大写字母
	fmt.Printf("%s", data) // {"name":"Vian","skill":["飞天","遁地"]}
}
