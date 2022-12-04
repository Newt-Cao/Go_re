package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 声明一个map,赋值
	hero := map[string]interface{}{
		"name":  "SuperMan",
		"age":   1000,
		"skill": []string{"飞天", "遁地"},
	}
	// 调用函数，map是引用类型，不需要传地址
	data, err := json.Marshal(hero)
	if err != nil {
		fmt.Println("Json err = ", err)
		return
	}
	fmt.Printf("%s", data) // {"age":1000,"name":"SuperMan","skill":["飞天","遁地"]}
}
