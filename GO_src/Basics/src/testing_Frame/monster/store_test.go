package monster

import (
	"testing"
)

// 测试函数一
func TestStore(t *testing.T) {
	// 实例化一个 Monster 对象
	monster := &Monster{
		Name:  "Vian",
		Age:   18,
		Skill: "飞天遁地",
	}
	// 测试
	res := monster.Store()
	if !res {
		t.Fatal("函数运行错误！")
	}
	t.Log("函数运行正常！")
}

// 测试函数二
func TestReStore(t *testing.T) {
	var monster Monster
	res := monster.ReStore()
	if !res {
		t.Fatal("函数运行错误！")
	}
	if monster.Name != "Vian" {
		t.Fatal("函数数据错误！")
	}
	t.Log("函数运行正常！")
}
