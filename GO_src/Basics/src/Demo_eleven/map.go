/* map声明 */

package main

import "fmt"

func main() {

	//========>第一种

	// 声明一个 map 类型, map[key:type]value:type ，基本属性与切片 slice 相似。
	var mymap map[string]string         //仅声明，没有开辟空间，需要使用 make（）函数
	mymap = make(map[string]string, 10) //10可省略
	mymap["one"] = "java"
	mymap["two"] = "golang"
	mymap["three"] = "python"

	fmt.Println(mymap)

	fmt.Println("---------------------------")

	// ========>第二种
	mymap2 := make(map[int]string)
	mymap2[1] = "java"
	mymap2[2] = "golang"
	mymap2[3] = "python"

	fmt.Println(mymap2)

	fmt.Println("---------------------------")

	// ========>第三种（不需要make，用于知道初始化值时）
	mymap3 := map[int]int{
		1: 10,
		2: 20,
		3: 30,
	}

	fmt.Println(mymap3)

	// 复杂的map声明 + 遍历
	students := make(map[string]map[string]string) // 开辟的是第一层map内存，而它的每个value都是一个独立的map，需要单独开辟内存
	students["1001"] = make(map[string]string)     // 需要开辟第二层的第一个map内存才可以使用，1001相当于第二次第一个map的变量名
	students["1001"]["name"] = "Vian"
	students["1001"]["gender"] = "girl"

	students["1002"] = make(map[string]string) // 开辟第二层的第二个map内存
	students["1002"]["name"] = "Jack"
	students["1002"]["gender"] = "man"

	for k, v := range students {
		fmt.Println("k=", k)
		for k1, v1 := range v {
			fmt.Printf("\tk1=%v,v1=%v", k1, v1)
		}
		fmt.Println()
	}
	/*
			k= 1001
		        k1=name,v1=Vian k1=gender,v1=girl
			k= 1002
		        k1=name,v1=Jack k1=gender,v1=man
	*/
}

// 输出对象是哈希序列
