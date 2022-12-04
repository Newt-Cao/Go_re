/* reflect 包，识别底层数据类型和具体值 */

package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Id    int
	Name  string
	Class string
}

func (this Student) Subject() {
	fmt.Printf("%v Go to class\n", this.Name)
}

func JudgType(arg interface{}) {
	//reflect.Type 和 reflect.Value,分别返回 reflect.Type 和 reflect.Value
	V := reflect.ValueOf(arg)
	A := reflect.TypeOf(arg)

	fmt.Printf("Value = %v\n", V) // Value = {10086 VivianChan 20软件}
	// Name() 方法，从该结构体的 reflect.Type 获取了结构体的名字
	fmt.Println("Name = ", A.Name()) //获取类型的名称 Name =  Student

	// reflect.Kind (type返回一个数据类型,kind返回该数据类型的种类)
	K := A.Kind()

	fmt.Println("type = ", A)
	fmt.Println("Kind = ", K) //  type =  main.Student  Kind =  struct

	// NumField() 方法返回结构体中字段的数量 | Field(i int) 方法返回字段 i 的 reflect.Value。
	if reflect.TypeOf(arg).Kind() == reflect.Struct {
		TypeNum := reflect.ValueOf(arg)

		fmt.Printf("Field num = %d\n", TypeNum.NumField())

		for i := 0; i < TypeNum.NumField(); i++ {
			// i为第1 、2、3 字段,返回(Name,PkgPath,Type,Tag,Offset,Index,Anonymous(嵌入式字段))
			fmt.Printf("type:%T		Value:%v\n", TypeNum.Field(i), TypeNum.Field(i))
		}

	}

	// Int() 和 String() 方法,帮助我们分别把 reflect.Value 还原为 int类型 和 string
	testInt := 137

	fmt.Printf("type = %T	value = %v\n", reflect.ValueOf(testInt).Int(), reflect.ValueOf(testInt).Int())

	testString := "Maven"

	fmt.Printf("type = %T	value = %v\n", reflect.ValueOf(testString).String(), reflect.ValueOf(testString).String())

	// Interface(),获取对应的value
	for i := 0; i < A.NumField(); i++ {
		field := A.Field(i)                                              //A: Type
		fmt.Println(field.Name, field.Type, "=", V.Field(i).Interface()) //V: Value
	}

	// NumMethod() 获取对象arg绑定的方法个数，Method(i) 返回字段 i 的方法(与NumField、Field类似，获取值时必须输入的时 reflect.ValueOf , Name、Type 输入 reflect.TypeOf)
	if reflect.TypeOf(arg).Kind() == reflect.Struct {
		numMethod := reflect.TypeOf(arg)

		// 注意NumMethod只对接口类型统计未导出的方法
		for i := 0; i < numMethod.NumMethod(); i++ {
			meThod := numMethod.Method(i)

			fmt.Printf("%s:%v", meThod.Name, meThod.Type)

		}
	}

}

func main() {
	Vivian := Student{
		Id:    10086,
		Name:  "VivianChan",
		Class: "20软件",
	}

	Vivian.Subject()

	JudgType(Vivian)
}

/*
type Type interface {
    // Kind返回该接口的具体分类
    Kind() Kind

    // Name返回该类型在自身包内的类型名，如果是未命名类型会返回""
    Name() string

    // PkgPath返回类型的包路径，即明确指定包的import路径，如"encoding/base64"
    // 如果类型为内建类型(string, error)或未命名类型(*T, struct{}, []int)，会返回""
    PkgPath() string

    // 返回类型的字符串表示。该字符串可能会使用短包名（如用base64代替"encoding/base64"）
    // 也不保证每个类型的字符串表示不同。如果要比较两个类型是否相等，请直接用Type类型比较。
    String() string

    // 返回要保存一个该类型的值需要多少字节；类似unsafe.Sizeof
    Size() uintptr

    // 返回当从内存中申请一个该类型值时，会对齐的字节数
    Align() int

    // 返回当该类型作为结构体的字段时，会对齐的字节数
    FieldAlign() int

    // 如果该类型实现了u代表的接口，会返回真
    Implements(u Type) bool

    // 如果该类型的值可以直接赋值给u代表的类型，返回真
    AssignableTo(u Type) bool

    // 如该类型的值可以转换为u代表的类型，返回真
    ConvertibleTo(u Type) bool

    // 返回该类型的字位数。如果该类型的Kind不是Int、Uint、Float或Complex，会panic
    Bits() int

    // 返回array类型的长度，如非数组类型将panic
    Len() int

    // 返回该类型的元素类型，如果该类型的Kind不是Array、Chan、Map、Ptr或Slice，会panic
    Elem() Type

    // 返回map类型的键的类型。如非映射类型将panic
    Key() Type

    // 返回一个channel类型的方向，如非通道类型将会panic
    ChanDir() ChanDir

    // 返回struct类型的字段数（匿名字段算作一个字段），如非结构体类型将panic
    NumField() int

    // 返回struct类型的第i个字段的类型，如非结构体或者i不在[0, NumField())内将会panic
    Field(i int) StructField

    // 返回索引序列指定的嵌套字段的类型，
    // 等价于用索引中每个值链式调用本方法，如非结构体将会panic
    FieldByIndex(index []int) StructField

    // 返回该类型名为name的字段（会查找匿名字段及其子字段），
    // 布尔值说明是否找到，如非结构体将panic
    FieldByName(name string) (StructField, bool)

    // 返回该类型第一个字段名满足函数match的字段，布尔值说明是否找到，如非结构体将会panic
    FieldByNameFunc(match func(string) bool) (StructField, bool)
    // 如果函数类型的最后一个输入参数是"..."形式的参数，IsVariadic返回真
    // 如果这样，t.In(t.NumIn() - 1)返回参数的隐式的实际类型（声明类型的切片）
    // 如非函数类型将panic
    IsVariadic() bool

    // 返回func类型的参数个数，如果不是函数，将会panic
    NumIn() int

    // 返回func类型的第i个参数的类型，如非函数或者i不在[0, NumIn())内将会panic
    In(i int) Type

    // 返回func类型的返回值个数，如果不是函数，将会panic
    NumOut() int

    // 返回func类型的第i个返回值的类型，如非函数或者i不在[0, NumOut())内将会panic
    Out(i int) Type

    // 返回该类型的方法集中方法的数目
    // 匿名字段的方法会被计算；主体类型的方法会屏蔽匿名字段的同名方法；
    // 匿名字段导致的歧义方法会滤除
    NumMethod() int

    // 返回该类型方法集中的第i个方法，i不在[0, NumMethod())范围内时，将导致panic
    // 对非接口类型T或*T，返回值的Type字段和Func字段描述方法的未绑定函数状态
    // 对接口类型，返回值的Type字段描述方法的签名，Func字段为nil
    Method(int) Method

    // 根据方法名返回该类型方法集中的方法，使用一个布尔值说明是否发现该方法
    // 对非接口类型T或*T，返回值的Type字段和Func字段描述方法的未绑定函数状态
    // 对接口类型，返回值的Type字段描述方法的签名，Func字段为nil
    MethodByName(string) (Method, bool)
    // 内含隐藏或非导出方法
}
*/
