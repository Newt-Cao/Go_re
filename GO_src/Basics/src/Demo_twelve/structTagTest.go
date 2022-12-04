package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title   string   `json:"电影名"`
	Release string   `json:"上映时间"`
	Price   int      `json:"价格"`
	Actor   []string `json:"演员"`
}

func main() {
	M := Movie{
		Title:   "大时代",
		Release: "1992",
		Price:   88,
		Actor:   []string{"刘青云", "郑少秋", "周慧敏"},
	}

	// 编码：结构体 ——> json
	// Marshal()方法,把结构体转换成json格式，并返回两个值。JsonStr,err(用于判断)
	jsonS, err := json.Marshal(M)
	if err != nil {
		fmt.Println("错误", err)
		return
	}
	fmt.Printf("JsonStr = %s\n", jsonS)
	// JsonStr = {"电影名":"大时代","上映时间":"1992","价格":88,"演员":["刘青云","郑少秋","周慧敏"]}

	// 解码：json ——>结构体
	m := Movie{}
	// 传入一个指针
	err = json.Unmarshal(jsonS, &m)

	if err != nil {
		fmt.Println("json Unmarshal error", err)
		return
	}
	fmt.Println(m)
	// {大时代 1992 88 [刘青云 郑少秋 周慧敏]}
}
