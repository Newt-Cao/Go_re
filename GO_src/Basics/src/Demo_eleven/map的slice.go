package main

import "fmt"

func main() {
	// map切片的简单案例(切片的每个元素是map类型)

	// 声明
	// monster := []map[string]string{}
	// 切片开辟空间
	// monster = make([]map[string]string, 2)
	// 简写
	monster := make([]map[string]string, 2)

	// 增加切片元素，map也需要开辟内存
	if monster[0] == nil {
		monster[0] = make(map[string]string)
		monster[0]["name"] = "蜘蛛精"
		monster[0]["age"] = "1000"
	}

	if monster[1] == nil {
		monster[1] = make(map[string]string)
		monster[1]["name"] = "白骨精"
		monster[1]["age"] = "100"
	}

	fmt.Println(monster) // [map[age:1000 name:蜘蛛精] map[age:100 name:白骨精]]

	// 当超过切片的make长度时，使用append动态增加
	newMonster := make(map[string]string)
	newMonster["name"] = "北京烤鸭"
	newMonster["age"] = "20"
	monster = append(monster, newMonster)

	fmt.Println(monster) // [map[age:1000 name:蜘蛛精] map[age:100 name:白骨精] map[age:20 name:北京烤鸭]]
}
