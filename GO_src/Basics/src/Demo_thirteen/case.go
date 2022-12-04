package main

/*
	案例一：
		1.写一个Cal结构体，写入字段Num1和Num2。
		2.绑定一个方法GetSub(name string)。
		3.使用反射遍历Cal结构体所有的字段信息。
		4.使用反射完成对GetSub的调用，输出的格式为：
			"xxx 完成了减法运算：x-x=x"
*/

import (
	"fmt"
	"reflect"
)

type Cal struct {
	Num1 int
	Num2 int
}

func (cal Cal) GetSub(name string) string {
	return fmt.Sprintf("%s 完成了减法运算题：%d - %d = %d", name, cal.Num1, cal.Num2, cal.Num1-cal.Num2)
}

func calReflect(x interface{}) {
	// 转成reflect类型
	xty := reflect.TypeOf(x).Elem()
	xva := reflect.ValueOf(x).Elem()
	// 获取字段数量
	fieldNum := xva.NumField()
	fmt.Println("字段数：", fieldNum)
	// 获取绑定方法数量
	methodNum := xva.NumMethod()
	fmt.Println("方法数：", methodNum)
	// 获取字段名、类型、种类
	for i := 0; i < fieldNum; i++ {
		// 获取字段名
		fieldName := xty.Field(i).Name
		// 获取字段类型
		fieldType := xty.Field(i).Type
		// 字段种类
		fieldKind := xva.Field(i).Kind()
		// 获取字段的值
		fieldValue := xva.Field(i)
		fmt.Printf("字段名：%v 字段类型：%v 字段种类：%v 字段值：%v\n", fieldName, fieldType, fieldKind, fieldValue)
	}
	// 给字段赋值
	xva.FieldByName("Num1").SetInt(4)
	xva.FieldByName("Num2").SetInt(2)
	// 传入参数调用函数
	params := []reflect.Value{reflect.ValueOf("Tom")}
	res := xva.Method(0).Call(params)
	fmt.Println(res)
}

func main() {
	cal := &Cal{}
	calReflect(cal)
}

/*
字段数： 2
方法数： 1
字段名：Num1 字段类型：int 字段种类：int 字段值：0
字段名：Num2 字段类型：int 字段种类：int 字段值：0
[Tom 完成了减法运算题：4 - 2 = 2]
*/
