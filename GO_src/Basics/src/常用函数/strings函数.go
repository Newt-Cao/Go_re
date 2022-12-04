package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 1.统计字符串长度
	str := "hello,world中国"
	fmt.Println("str len=", len(str)) // 11
	// 搭配遍历使用，可以遍历字符串，使用len()是按一个字节遍历，一个汉字占3个字节所以出现乱码
	// 转换成切片 []rune(str) 进行遍历。
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c\n", str2[i])
	}
	// 使用range遍历
	for j, i := range str {
		fmt.Printf("%v:%c\n", j, i)
	}

	// 2.string转int：i,err := strconv.Atoi(string),是ParseInt(s,10,0)的简写
	i, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误:", err)
	} else {
		fmt.Printf("%T,%v\n", i, i) // int,123
	}

	// 3.int转string：str := strconv.Itoa(int)，是FormatInt(i, 10) 的简写。
	str3 := strconv.Itoa(123456)
	fmt.Printf("%T,%#v\n", str3, str3) // string,"123456"

	// 4.string转[]byte(): byte := []byte(string) ,常用在写入二进制文件。
	var bytee = []byte("hello")
	fmt.Printf("%T,%c\n", bytee, bytee) // []uint8,[h e l l o]

	// 5.[]byte()转string: str := string([]byte{97,98,99})。
	str4 := string([]byte{97, 98, 99})
	fmt.Printf("%T,%v\n", str4, str4)

	// 6.10进制转2、8、16进制。
	num2 := strconv.FormatInt(10, 2)
	num8 := strconv.FormatInt(20, 8)
	num16 := strconv.FormatInt(200, 16)
	fmt.Printf("num2 = %T,%v\nnum8 = %T,%v\nnum16 = %T,%v\n", num2, num2, num8, num8, num16, num16)
	/*
		num2 = string,1010
		num8 = string,24
		num16 = string,c8
	*/

	// 7.查询字符串中是否存在指定子字符串,通常与if使用,存在返回 true ，否则返回 false。
	if strings.Contains("Vivianchan", "vi") {
		fmt.Printf("存在该子字符串\n")
	} else {
		fmt.Printf("不存在该子字符串\n")
	}
	// 查询是否以指定字符串开头：
	fmt.Println(strings.HasPrefix("Vivianchan", "V")) //true
	// 查询是否以指定字符串结尾：
	fmt.Println(strings.HasSuffix("Vivianchan", "V")) //false

	// 8.统计字符串中指定的子字符串出现的次数。
	str5 := strings.Count("VivianChan", "v")
	fmt.Println("v出现的次数为：", str5) // v出现的次数为： 1

	// 9.字符串间的比较。
	// 区分大小写的比较：
	fmt.Println("Abc" == "abc") //false
	fmt.Println("abc" == "abc") //true
	// 不区分大小写：
	fmt.Println(strings.EqualFold("Abc", "abc")) //true
	fmt.Println(strings.EqualFold("abc", "abc")) //true

	// 10.子字符串索引的查询，如果没有返回 -1 。
	// 查询第一次出现：
	if err := strings.Index("vivianchan", "b"); err == -1 {
		fmt.Println("不存在该字符\n", err)
	} else {
		fmt.Println("该字符出现的索引为：", err)
	} //不存在该字符 -1

	// 查询最后一次出现：
	if err := strings.LastIndex("vivianchan", "an"); err == -1 {
		fmt.Println("不存在该字符", err)
	} else {
		fmt.Println("该字符出现的索引为：", err)
	} //该字符出现的索引为： 8

	// 11.子字符串间的替换，值拷贝，把替换完的string赋给一个新的变量，可替换一部分，也可替换全部: strings.Replace()。
	str6 := "python python python"
	str7 := strings.Replace(str6, "python", "golang", 1)  // 替换第一个
	str8 := strings.Replace(str6, "python", "golang", -1) // 替换全部
	fmt.Printf("%v\n%v\n%v\n", str6, str7, str8)
	/*
		python python python
		golang python python
		golang golang golang
	*/

	// 12.用指定字符作为分割标识进行分割，将字符串拆分成字符串切片。
	str9 := strings.Split("v,i,v,i,a,n", ",")
	fmt.Println(str9) //[v i v i a n]

	// 13.字符串大小写相互转换。
	str10 := strings.ToLower("VIVIANCHAN")
	fmt.Println(str10) // vivianchan
	str11 := strings.ToUpper("vivianchan")
	fmt.Println(str11) // VIVIANCHAN

	// 14.去掉字符串两边的空格。
	str12 := strings.TrimSpace("  Vian  ")
	fmt.Printf("%q\n", str12) // "Vian"
	// 去掉指定字符。
	str13 := strings.Trim("! Vian !", " ! ")
	fmt.Printf("%q\n", str13) // "Vian"
	// 去掉左边指定字符。
	fmt.Println(strings.TrimLeft("!Vian", "!")) //vian
	// 去掉右边指定字符。
	fmt.Println(strings.TrimRight("Vian!", "!")) //vian
}
