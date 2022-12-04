//基本数据类型和string间的互换

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 方法一 fmt.Sprintf("数据原类型",要转换的数据)：

	var number int = 64
	var decimals float64 = 3.14
	var boole bool = true
	var char byte = 'h'

	str := fmt.Sprintf("%d", number)
	str1 := fmt.Sprintf("%.1f", decimals)
	str2 := fmt.Sprintf("%t", boole)
	str3 := fmt.Sprintf("%c", char)

	fmt.Printf("%T,%v,%v,%v,%v\n", str, str, str1, str2, str3)
	// string,64,3.1,true,hs

	// 方式二 strconv 包：

	var number1 int = 64
	var number2 int = 65
	var decimals1 float64 = 3.14
	var boole1 bool = false

	//  FormatInt(原数据变量名,进制)
	str4 := strconv.FormatInt(int64(number1), 10)

	// FormatFloat(原数据变量名,'转换的格式',保留小数点几位,浮点类型)
	str5 := strconv.FormatFloat(decimals1, 'f', 10, 64)

	// FormatBool(原数据类型)
	str6 := strconv.FormatBool(boole1)

	// Itoa(int类型)=>sting
	str7 := strconv.Itoa(number2)
	fmt.Printf("%T,%v\n", str7, str7) //string,65

	fmt.Printf("%T,%T,%T,%v,%v,%v", str4, str5, str6, str4, str5, str6)

	//输出 string,string,string,64,3.1400000000,false
}
