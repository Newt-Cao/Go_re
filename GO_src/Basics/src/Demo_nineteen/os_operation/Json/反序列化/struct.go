package main

import (
	"encoding/json"
	"fmt"
)

/*
其他类型类似，只需要把结构体更改为其他数据类型，但是要匹配序列化时的数据类型，否则反序列化失去意义
*/

// 结构体声明，名称可以和序列化时的不一样，但是字段必须一样，最好名称也一样
type Hero struct {
	Name  string   `json:"name"`
	Skill []string `json:"skill"`
}

func main() {
	// 需要反序列化的字符串
	str := `{"name":"Vian","skill":["飞天","遁地"]}`
	// 实例化一个默认值结构体，用于接收反序列化数据
	var hero Hero
	// 反序列化，需要传入地址，才可改变hero
	err := json.Unmarshal([]byte(str), &hero)
	if err != nil {
		fmt.Println("Antitone err = ", err)
	}
	// 打印，反序列化后，hero一样可以调用字段
	fmt.Println(hero, hero.Name) // {Vian [飞天 遁地]} Vian
}
