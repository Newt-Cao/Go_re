package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	stackSize int    // 栈容量
	top       int    // 栈顶
	array     [5]int // 使用数组模拟
}

// 入栈
func (s *Stack) pushData(value int) (err error) {
	// 先判断是否栈满
	if s.top == s.stackSize {
		return errors.New("Stack full！")
	}

	// 若栈没满，可以添加数据
	s.top++
	s.array[s.top] = value

	return
}

// 出栈
func (s *Stack) popData() (value int, err error) {
	// 先判断是否栈空
	if s.top == -1 {
		return 0, errors.New("Stack empty!")
	}

	// 若不空弹出最后进入的一个元素
	value = s.array[s.top]
	s.top--

	return value, nil
}

// 展示
func (s *Stack) showStack() (err error) {
	// 先判断栈是否为空
	if s.top == -1 {
		return errors.New("Stack empty!")
	}

	// 若不为空，可以遍历，反向输出
	for i := s.top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, s.array[i])
	}

	return
}

func main() {
	// 实例化一个栈
	stack := &Stack{
		stackSize: 5,
		top:       -1,
	}

	stack.pushData(1)
	stack.pushData(2)
	stack.pushData(3)
	stack.pushData(4)
	stack.pushData(5)

	stack.showStack()

	val, err := stack.popData()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(val)
	}

	stack.showStack()
}

/*
arr[4]=5
arr[3]=4
arr[2]=3
arr[1]=2
arr[0]=1
5
arr[3]=4
arr[2]=3
arr[1]=2
arr[0]=1
*/
