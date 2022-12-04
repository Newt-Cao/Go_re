// 数据类型的转换

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var Vian int64 = 100
	var vian float64 = float64(Vian)

	fmt.Printf("%v,%T,%v,%T\n", Vian, Vian, vian, vian)
	// 输出 100,int64,100,float64

	// string转基本类型使用 strconv 包

	var str string = "true"
	var str1 string = "3.14"
	var str2 string = "100"

	// 注意：通常会返回两个值：value和error

	// string -> bool,strconv.ParseBool(string)
	boole, _ := strconv.ParseBool(str)

	// string -> float,strconv.ParseFloat(string,浮点类型)
	decimals, _ := strconv.ParseFloat(str1, 64)

	// string -> int,strconv.ParseInt(string,进制,整型类型)
	integer, _ := strconv.ParseInt(str2, 10, 64)
	//ParseInt可以简写成：strconv.Atoi(s string)(i int,err error)
	integer2, _ := strconv.Atoi(str2)

	fmt.Printf("boole = %v,type = %T\n", boole, boole)
	fmt.Printf("decimals = %v,type = %T\n", decimals, decimals)
	fmt.Printf("integer = %v,type = %T\n", integer, integer)
	fmt.Printf("integer2 = %v,type = %T\n", integer2, integer2)
	/*  boole = true,type = bool
	decimals = 3.14,type = float64
	integer = 100,type = int64 */
}
