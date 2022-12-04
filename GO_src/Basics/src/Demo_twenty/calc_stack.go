package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
案例：模拟计算机底层简单计算
思路：
	1.实例一个符号栈operStack、数字栈numStack。
	2.用一个变量scanIndex := 0作为下标扫描表达式exp：3+2*60-2
	3.扫描到如果是数字则放入到数字栈
	4.扫描到符号栈时：
		（1）若栈为空时，直接放入
		（2）若栈不为空时：
				a.当栈顶的数据优先级高于将要压栈的数据时，需要先把栈顶的数据弹出，然后数字栈弹出两个数据，计算结果后再压入数字栈，压栈的符号待计算
				  完毕后也压入符号栈

				b.当栈顶的数据优先级小于将要压栈的数据时，直接放入
	5.表达式压栈结束后，依次从栈中弹出计算，默认是一个符号对应两个数字，知道符号栈空
*/

// 声明一个结构体操作模拟栈
type Stack struct {
	maxSize int     // 容量
	top     int     // 栈顶
	array   [20]int // 数组模拟栈
}

// 入栈
func (s *Stack) pushStack(value int) (err error) {
	// 判断栈是否已满
	if s.top == s.maxSize {
		return errors.New("Stack full！")
	}

	// 不满则可以加入
	s.top++
	s.array[s.top] = value

	return
}

// 出栈
func (s *Stack) popStack() (value int, err error) {
	// 判断是否栈空
	if s.top == -1 {
		err = errors.New("Stack empty!")
	}

	// 若不空则可以弹出
	value = s.array[s.top]
	s.top--

	return value, err
}

// 栈浏览
func (s *Stack) showStack() (err error) {
	// 判断是否栈空
	if s.top == -1 {
		return errors.New("Stack empty!")
	}

	fmt.Println("目前栈中数据：")

	// 不空则遍历，倒着遍历
	for i := s.top; i >= 0; i-- {
		fmt.Printf("arr[%d]=[%d]\n", i, s.array[i])
	}

	return
}

// 判断是符号或者数字，如果是符号还要返回优先级
func (s *Stack) judgeChar(char int) (int, bool) {
	// 符号都返回true，[*、/] 优先级：1，[+、-] 优先级：0
	if char == 42 || char == 47 {
		return 1, true
	} else if char == 43 || char == 45 {
		return 0, true
	} else {
		return -1, false
	}
}

// 用于计算弹出的数据
func (s *Stack) calc(num1, num2, oper int) int {
	res := 0
	switch oper {
	case 42:
		res = num2 * num1
	case 47:
		res = num2 / num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	}
	return res
}

// 扫描表达式，并分别压入数字栈或符号栈
func scanExp(exp string, numStack, operStack *Stack) (err error) {
	// scanIndex := 0作为下标扫描
	scanIndex := 0

	// 用于拼接多位数
	jointChar := ""

	// 循环扫描
	for {
		// 每次扫描一个字符
		char := exp[scanIndex : scanIndex+1]

		// 把扫描的字符转成ascll码，[]byte(char)把char转成ascll，[0]再取出第一个元素，类型为uint8，转成int
		ascllChar := int([]byte(char)[0])

		// 判断该字符是数组还是符号，返回优先级和布尔值
		if prioriy, judgeBool := operStack.judgeChar(ascllChar); judgeBool {

			// 如果是true，则分两种情况压入符号栈，栈为空直接入栈
			if operStack.top == -1 {

				operStack.pushStack(ascllChar)

				// 不为空时也分两种情况入栈
			} else {

				/* 取得栈顶符号优先级和将要压栈的符号优先级比较，如果栈顶优先级比较大，则要弹出同时也弹出
				数字栈的两个数字计算完后再压入数字栈，最后再把优先级低于栈顶的符号压栈 */
				if topPrioriy, _ := operStack.judgeChar(operStack.array[operStack.top]); topPrioriy > prioriy {

					// 弹出符号栈，栈顶符号
					oper, err := operStack.popStack()
					if err != nil {
						return err
					}

					// 弹出数字栈的两个数字
					num1, err := numStack.popStack()
					if err != nil {
						return err
					}
					num2, err := numStack.popStack()
					if err != nil {
						return err
					}

					// 计算出结果
					res := operStack.calc(num1, num2, oper)

					// 计算完压入数字栈
					numStack.pushStack(res)

					// 最后再把优先级低于栈顶的符号压栈
					operStack.pushStack(ascllChar)

				} else {

					// 如果topPrioriy < prioriy直接压入
					operStack.pushStack(ascllChar)
				}
			}
		} else {

			// 如果是false，则压入数字栈，需要把ascll转回来再压进去，如果数字后面还是数字，一直执行拼接
			jointChar += char

			// 解决多位数问题

			// 如果扫描完已经到表达式尾部则直接压栈
			if scanIndex == len(exp)-1 {
				val, _ := strconv.ParseInt(jointChar, 10, 64)
				numStack.pushStack(int(val))
			} else {

				// 数字后面不是数字就压栈，把jointChar清零
				if _, judgeNum := numStack.judgeChar(int([]byte(exp[scanIndex+1 : scanIndex+2])[0])); judgeNum {
					val, _ := strconv.ParseInt(jointChar, 10, 64)
					numStack.pushStack(int(val))
					jointChar = ""
				}
			}
		}

		// 判断是否到达表达式尾部
		if scanIndex+1 == len(exp) {
			break
		}

		// 没到达继续循环
		scanIndex++
	}

	// 把数和符号从栈里取出来运算
	for {
		// 先判断符号栈是否为空
		if operStack.top == -1 {

			break

		} else {
			// 弹出符号栈，栈顶符号
			oper, err := operStack.popStack()
			if err != nil {
				return err
			}

			// 弹出数字栈的两个数字
			num1, err := numStack.popStack()
			if err != nil {
				return err
			}
			num2, err := numStack.popStack()
			if err != nil {
				return err
			}

			// 计算出结果
			res := operStack.calc(num1, num2, oper)

			// 计算完压入数字栈
			numStack.pushStack(res)
		}
	}

	// 弹出数字栈最后一个字符
	answer, _ := numStack.popStack()
	fmt.Printf(" %s 的计算结果是：%d ", exp, answer)

	return
}

func main() {
	// 实例数字栈和符号栈
	numStack := &Stack{
		maxSize: 20,
		top:     -1,
	}
	operStack := &Stack{
		maxSize: 20,
		top:     -1,
	}

	// 表达式
	exp := "3+2*60-2"

	// 扫描表达式
	scanExp(exp, numStack, operStack) //  3+2*60-2 的计算结果是：121
}
