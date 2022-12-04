package main

import (
	"fmt"
	"reflect"
)

// 结构体
type Human struct {
	Name string
	Age  int
}

// 绑定三个方法
func (H Human) OutPut() {
	fmt.Println(H)
}

func (H Human) FieldTraverse(n int) {
	for i := 0; i < n; i++ {
		fmt.Print(i)
	}
	fmt.Println()
}

func (H Human) Sum(n1, n2 float64) float64 {
	return n1 + n2
}

func main() {
	// 实例
	Human := Human{"Vian", 18}
	// 判断该实例是否为结构体
	if reflect.TypeOf(Human).Kind() == reflect.Struct {
		// 如果是则转换成reflect.Type类型
		HumanReflectType := reflect.TypeOf(Human)
		HumanReflectValue := reflect.ValueOf(Human)
		// 获取结构体方法数
		numMethod := HumanReflectType.NumMethod()
		numMethod2 := HumanReflectValue.NumMethod()
		fmt.Printf("numMethod: %v\n", numMethod)   // 3
		fmt.Printf("numMethod2: %v\n", numMethod2) // 3
		// 方法顺序： F ——> O ——> S
		// 调用第二个方法，无参数
		HumanReflectValue.Method(1).Call(nil) // {Vian 18}
		// 调用第一个方法，有参数
		param := []reflect.Value{}
		param = append(param, reflect.ValueOf(5))
		HumanReflectValue.Method(0).Call(param) // 01234
		// 调用第三个方法，有参数
		params := []reflect.Value{}
		num := []float64{1, 2}
		for _, v := range num {
			params = append(params, reflect.ValueOf(v))
		}
		// 返回的是reflect.Value切片，具体元素看返回值的多少，这里只有一个返回值，所以取索引为 0
		res := HumanReflectValue.Method(2).Call(params)
		fmt.Println(res[0]) // 3 但是类型还是 reflect.Value
		// 需要调用interface()后进行类型断言或者已知类型直接转换
		fmt.Println(res[0].Float())
	}
}
