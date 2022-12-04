package main

import "fmt"

func main() {
	// 一、map的增加。
	hero := make(map[int]string)
	hero[1] = "Jack"
	hero[2] = "Fuck"
	fmt.Println(hero) // map[1:Jack 2:Fuck]

	// 二、map的删除（内置函数delete(map,key)，不管key是否存在都不会报错。）。

	// 删除单个key-value
	delete(hero, 1)
	fmt.Println(hero) // map[2:Fuck]

	// 一次性删除所有key-value
	// 1.重新make一个新的空间，相当于格式化之前的数据。
	hero = make(map[int]string)
	fmt.Println(hero) // map[]
	// 2.遍历删除
	hero[1] = "Jack"
	hero[2] = "Fuck"
	fmt.Println(hero) // map[1:Jack 2:Fuck]
	for key, _ := range hero {
		delete(hero, key)
	}
	fmt.Println(hero) // map[]

	// 三、map的更改（重复使用key时的操作识别为更改。）。
	hero[1] = "Jack"
	hero[2] = "Fuck"
	fmt.Println(hero) // map[1:Jack 2:Fuck]
	hero[1] = "Vian"
	fmt.Println(hero) // map[1:Vian 2:Fuck]

	// 四、map的查找，返回一个状态判断，true为有，false为无。
	val, keyState := hero[3]
	if keyState {
		fmt.Println("找到对应的value：", val) // 找到对应的value： Vian
	} else {
		fmt.Println("找不到！") // 找不到！
	}
}
