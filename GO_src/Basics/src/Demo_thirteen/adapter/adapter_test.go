package test

import (
	"reflect"
	"testing"
)

// 函数测试
func TestReflectFunc(t *testing.T) {
	// 定义第一个函数
	call1 := func(v1, v2 int) {
		// 测试输出参数值
		t.Log("参数", v1, v2)
	}
	// 定义第二个函数
	call2 := func(v1, v2 int, s string) {
		// 测试输出参数值
		t.Log("参数", v1, v2, s)
	}

	var (
		// 用于调用不同函数
		function reflect.Value
		// 接收函数参数
		inValue []reflect.Value
		// 统计桥梁函数接收参数的长度
		n int
	)
	// 适配函数一和二，第一个参数是传递函数，第二个参数是传递的函数的参数，接收的是参数并放入到args切片中
	bridge := func(call interface{}, args ...interface{}) {
		// 统计传入函数的参数长度
		n = len(args)
		// 定义一个reflect.Value切片用于接收参数
		inValue = make([]reflect.Value, n)
		// 遍历接收参数并复制给reflect.Value切片
		for i := 0; i < n; i++ {
			inValue[i] = reflect.ValueOf(args[i])
		}
		// 把函数转成reflect.Value类型后赋值给function
		function = reflect.ValueOf(call)
		// 同一调用函数，inValue是参数
		function.Call(inValue)
	}
	// 调用桥梁函数
	bridge(call1, 1, 2)
	bridge(call2, 1, 2, "test2")
}

/*

=== RUN   TestReflectFunc
    adapter_test.go:13: 参数 1 2
    adapter_test.go:18: 参数 1 2 test2
--- PASS: TestReflectFunc (0.00s)
PASS
ok      vivianchan.cn/Basics/src/Demo_thirteen/adapter  0.321s

*/
