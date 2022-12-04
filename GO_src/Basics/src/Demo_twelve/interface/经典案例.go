package main

import (
	"fmt"
	"sort"
)

/*
使用sort的Sort函数，需要实现接口Interface，包含方法：
	len()
	Less()
	Swap()
*/

// 声明一个结构体
type Hero struct {
	Name string
	Age  int
}

// 结构体切片
type HeroSlice []Hero

// 实现接口方法
func (hero HeroSlice) Len() int {
	return len(hero)
}

func (hero HeroSlice) Less(i, j int) bool {
	return hero[i].Age < hero[j].Age
}

func (hero HeroSlice) Swap(i, j int) {
	hero[i], hero[j] = hero[j], hero[i]
}

func main() {
	// 定义一个HeroSilce实例化对象,对其进行排序
	var heros HeroSlice = []Hero{{"Vian-10", 99}, {"Jack-9", 98}, {"Fanc-33", 88}, {"Candy-46", 78}, {"Joke-28", 90}}

	// 排序前
	for _, v := range heros {
		fmt.Println(v)
		/*
			{Vian-10 99}
			{Jack-9 98}
			{Fanc-33 88}
			{Candy-46 78}
			{Joke-28 90}
		*/
	}

	fmt.Println()

	sort.Sort(heros)
	for _, v := range heros {
		fmt.Println(v)
		/*
			{Candy-46 78}
			{Fanc-33 88}
			{Joke-28 90}
			{Jack-9 98}
			{Vian-10 99}
		*/
	}

}
