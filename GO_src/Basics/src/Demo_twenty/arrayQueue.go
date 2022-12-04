package main

import (
	"errors"
	"fmt"
	"os"
)

/*
案例：使用数组模拟队列

思路：
	1.当rear=front时，说明队列为空
	2.当rear=MaxSize时，说明队列已满
	3.front表示队首，但不包含第一个元素
	4.rear表示队尾，但包含最后一个元素
	5.初始化
		MaxSize=len(Array)
		front,rear=-1

弊端：
	无法有效利用数组空间
*/

// 声明一个结构体操作数组
type ArrayQueue struct {
	MaxSize, front, rear int    // 数组最大索引，队首，对尾
	Array                [5]int // 数组
}

// 向队列添加元素
func (aq *ArrayQueue) AddElement(value int) (err error) {
	// 当rear=MaxSize时，说明队列已满
	if aq.rear == aq.MaxSize {
		return errors.New("Queue full!")
	}

	// 没满时，加入元素
	aq.rear++
	aq.Array[aq.rear] = value
	fmt.Println("Element add successfully!")

	return
}

// 展示队列
func (aq *ArrayQueue) ShowList() {
	// 从队首遍历到队尾
	for i := aq.front + 1; i <= aq.rear; i++ {
		fmt.Printf("array[%d]=%d\n", i, aq.Array[i])
	}
	fmt.Println("Queue Show successfully!")
}

// 向队列取出元素
func (aq *ArrayQueue) GetElement() (value int, err error) {
	// 当rear=front时，说明队列为空
	if aq.front == aq.rear {
		return -1, errors.New("Queue empty!")
	}

	// 当不是空时
	aq.front++
	value = aq.Array[aq.front]
	return value, err
}

func main() {
	// 实例化一个模拟队列
	arrayqueue := &ArrayQueue{
		MaxSize: 5,
		front:   -1,
		rear:    -1,
	}

	// 用户输入
	var keyControl string
	var value int

	for {
		fmt.Print("请输入操作：")
		fmt.Scanln(&keyControl)

		switch keyControl {

		case "add":
			fmt.Print("请输入添加的元素：")
			fmt.Scanln(&value)
			err := arrayqueue.AddElement(value)
			if err != nil {
				fmt.Println(err.Error())
				return
			} else {
				fmt.Println("加入队列成功！")
			}

		case "get":
			val, err := arrayqueue.GetElement()
			if err != nil {
				fmt.Println(err.Error())
				return
			} else {
				fmt.Println("取出的元素是：", val)
			}

		case "show":
			arrayqueue.ShowList()
			fmt.Println("队列展示成功！")

		case "exit":
			os.Exit(0)

		default:
			fmt.Println("输入有误!")
			continue
		}
	}
}
