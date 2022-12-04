/*******************************************************
 *	                       _oo0oo_
 *	                      o8888888o
 *	                      88" . "88
 *	                      (| -_- |)
 *	                      0\  =  /0
 *	                    ___/`---'\___
 *	                  .' \\|     |// '.
 *	                 / \\|||  :  |||// \
 *	                / _||||| -:- |||||- \
 *	               |   | \\\  - /// |   |
 *	               | \_|  ''\---/''  |_/ |
 *	               \  .-\__  '-'  ___/-. /
 *	             ___'. .'  /--.--\  `. .'___
 *	          ."" '<  `.___\_<|>_/___.' >' "".
 *	         | | :  `- \`.;`\ _ /`;.`/ - ` : | |
 *	         \  \ `_.   \_ __\ /__ _/   .-` /  /
 *	     =====`-.____`.___ \_____/___.-`___.-'=====
 *	                       `=---='
 *
 *
 *	     ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *
 *	           佛祖保佑     永不宕机     永无BUG
 *
 *	Author       : Mr.Chan
 *	Date         : 2022-08-13 11:57:51
 *	LastEditors  : Mr.Chan
 *	LastEditTime : 2022-08-20 22:10:52
 *	FilePath     : /vivianchan.cn/Basics/src/Demo_twenty/circleQueue.go
 *******************************************************/

package main

import (
	"errors"
	"fmt"
	"os"
)

/*
案例：使用数组模拟环形队列

思路：
	1.(tail+1) % MaxSize=head时，表示队列已满
	2.tail==head时，表示队列为空
	3.tail、head初始化为0，head表示队首包含第一个元素，tail表示队尾不含最后一个元素
	4.MaxSize为数组长度
	5.尾索引的下一个索引为头索引时表示队列已满，所以在判断是需要将队列容量空出一个作为判断
	6.统计队列总元素 =（tail+MaxSize-head）% MaxSize
	7.环形列表实际容量是MaxSize-1

优化：有效利用数组空间
*/

// 声明结构体操作数组
type CircleQueue struct {
	head, tail, MaxSize int    // 队首、队尾、数组长度
	array               [5]int // 数组
}

// 添加元素
func (cq *CircleQueue) pushElement(value int) (err error) {
	// 判断队列是否已满
	if cq.IsFull() {
		return errors.New("Queue full")
	}

	// 未满时将值赋给数组尾索引tail，tail从0开始
	cq.array[cq.tail] = value

	// 因为tail从0开始不包含最后一个元素，并且需要拿tail的后一个位置来判断队列是空还是满，所以(cq.tail + 1) % cq.MaxSize取模，tail就从1开始
	// 如果MaxSize固定为5时，0-3取模还是被除数本身，当tail=5时，取模为0则又返回一开始，因为是模拟队列，所以不是数组有无限长可以增加，而是一直在循环
	// 因此要进行第一步的if判断，不然就是一直在对数组的元素重新赋值，有了第一个判断才可以实现队列已满不可加入的原则
	cq.tail = (cq.tail + 1) % cq.MaxSize
	return
}

// 取出元素
func (cq *CircleQueue) popElement() (value int, err error) {
	// 判断队列是否为空
	if cq.IsEmpty() {
		return 0, errors.New("Queue empty")
	}

	// 不空时可以取出队列第一个元素赋值给value，head从0开始
	value = cq.array[cq.head]

	// 当取出第一个value时，head+1后队首就是原来队列的第二个，道理同上，数组MaxSize固定为5时，0-3可以取出，当到第5个时取模，head又变回0
	// 即满足cq.tail == cq.head时，表示队列为空，return 0, errors.New("Queue empty")
	cq.head = (cq.head + 1) % cq.MaxSize
	return
}

// 展示队列
func (cq *CircleQueue) showQueue() {
	// 判断队列是否为空
	size := cq.Queue()
	if size == 0 {
		fmt.Println("队列空")
	}

	// 取出一个临时值不影响原来head操作队列
	temphead := cq.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", temphead, cq.array[temphead])
		temphead = (temphead + 1) % cq.MaxSize
	}
	fmt.Println()
	return
}

// 判断队列是否已满
func (cq *CircleQueue) IsFull() bool {
	// 这里是tail尾索引的后一个位置作为判断时，与头索引相等，表示队列已满，个人理解取模是为了可以循环利用数组空间，使tail不断从0开始
	return (cq.tail+1)%cq.MaxSize == cq.head
}

// 判断队列是否为空
func (cq *CircleQueue) IsEmpty() bool {
	return cq.tail == cq.head
}

// 获取队列
func (cq *CircleQueue) Queue() int {
	// 这里的tail和head是相互关联的，当队列的容量MaxSize固定时，知道tail就知道0<=head<=tail，MaxSize也是为了可以循环展示队列范围固定为5个数
	return (cq.tail + cq.MaxSize - cq.head) % cq.MaxSize
}

func main() {
	queue := &CircleQueue{
		MaxSize: 5,
		head:    0,
		tail:    0,
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
			err := queue.pushElement(value)
			if err != nil {
				fmt.Println(err.Error())
				return
			} else {
				fmt.Println("加入队列成功！")
			}

		case "get":
			val, err := queue.popElement()
			if err != nil {
				fmt.Println(err.Error())
				return
			} else {
				fmt.Println("取出的元素是：", val)
			}

		case "show":
			queue.showQueue()
			fmt.Println("队列展示成功！")

		case "exit":
			os.Exit(0)

		default:
			fmt.Println("输入有误!")
			continue
		}
	}
}
