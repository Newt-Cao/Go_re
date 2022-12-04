/*
一、tag说明
先介绍一下``(反引号)：反引号用来创建原生的字符串字面量 ，这些字符串可能由多行组成(不支持任何转义序列)，
原生的字符串字面量多用于书写多行消息、HTML以及正则表达式。
*/

package main

import (
	"fmt"
	"reflect"
)

type People struct {
	Ability string `json:"技能" doc:"Skill"` //Tag 使用 key：value 引用。
	Sex     string `json:"性别"`
	Age     int    `json:"年龄"`
}

func GetTag(Pp interface{}) {
	// Elem() 返回全部元素。(Array, Chan, Map, Ptr,Slice)
	P := reflect.TypeOf(Pp).Elem()

	for i := 0; i < P.NumField(); i++ {
		// 已知标签名，获取标签值
		tagjson := P.Field(i).Tag.Get("json")
		tagdoc := P.Field(i).Tag.Get("doc")

		fmt.Println("json = ", tagjson, "doc = ", tagdoc)
	}
}

func main() {
	var people People
	GetTag(&people)
	/*
		json =  技能 doc =  Skill
		json =  性别 doc =
		json =  年龄 doc =
	*/
}
