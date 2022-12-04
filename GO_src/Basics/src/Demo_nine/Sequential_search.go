package main

import "fmt"

func Sequential_Search(name string, array [4]string) {
	/* 给出一个数组， 从键盘输入一个元素，判断该数列是否包含此元素。*/
	var flag = true

	i := 0
	for i < len(array) {
		if name == array[i] {
			flag = false
			break
		}
		i++
	}

	if flag {
		fmt.Println("找不到指定人物！")
	} else {
		fmt.Printf("找到名字 %q 序号为：%v", name, i)
	}
}

func main() {
	array := [4]string{"LeetC", "Vian", "Sand", "Chan"}
	Sequential_Search("LeetC", array)
	/* 找到名字 "LeetC" 序号为：0 */
}
