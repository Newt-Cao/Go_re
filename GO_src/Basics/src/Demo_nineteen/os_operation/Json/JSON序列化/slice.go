package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 切片声明
	sliceJson := make([]map[string]interface{}, 2)
	// 赋值
	sliceJson[0] = make(map[string]interface{})
	sliceJson[0]["name"] = "蜘蛛侠"
	sliceJson[0]["skill"] = []string{"吐丝", "爬墙"}

	sliceJson[1] = make(map[string]interface{})
	sliceJson[1]["name"] = "蝙蝠侠"
	sliceJson[1]["skill"] = []string{"科技", "飞天"}
	// 序列化
	data, err := json.Marshal(sliceJson)
	if err != nil {
		fmt.Println("Json err = ", err)
	}
	// 打印，转化
	fmt.Printf("%s", data) // [{"name":"蜘蛛侠","skill":["吐丝","爬墙"]},{"name":"蝙蝠侠","skill":["科技","飞天"]}]
}
