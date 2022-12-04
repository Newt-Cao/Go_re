package main

import (
	"fmt"
	"reflect"
)

// 用于判断接收任意类型并转换成reflect类型
func reflectChange(x interface{}) {
	// 将x转成reflect.Type，虽然输出时仍然是x对应的类型，但是底层是x对应的reflect类型
	changeLateType := reflect.TypeOf(x)
	fmt.Printf("%T\t%v\n", changeLateType, changeLateType) // *reflect.rtype	int

	changeLateValue := reflect.ValueOf(x)
	fmt.Printf("%v\n", changeLateValue) // 10

	// num3 := changeLateValue + num2，不可以直接操作，使用Int()方法可以取出reflect.Value对应类型的值后使用
	num2 := changeLateValue.Int() + 10
	fmt.Println(num2) // 20

	// 将其还原，先还原为接口，再断言还原为原来类型
	restore := changeLateValue.Interface()
	res := restore.(int)
	fmt.Printf("%T\t%v\n", res, res) // int	10
}

func main() {
	var num int = 10
	reflectChange(num)
}
