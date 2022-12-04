package calc

import (
	"testing"
)

// 测试函数必须是 TestXxx(t *testing.T){}
func TestAddUpper(t *testing.T) {
	// 测试第一个累加
	res := addUpper(10)
	// 错误时执行
	if res != 55 {
		t.Fatalf("错误!期望值：%v", 55)
	}
	// 正确时执行
	t.Logf("正确！")
}

func TestGetSubtract(t *testing.T) {
	// 测试第二个相减
	res := getSubtract(10, 7)
	if res != 3 {
		// t.Errorf() 输出一般错误并退出
		// t.Fatalf() 输出致命错误并退出
		t.Fatalf("错误！期望值：%v", 3)
	}
	t.Logf("正确!")
}

// 基准测试
func BenchmarkAddUpper(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addUpper(10)
	}
}
