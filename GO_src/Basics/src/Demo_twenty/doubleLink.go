package main

import "fmt"

/*
案例：在单链表实现水浒传英雄顺序的增删改查
*/

// 声明结构体操作链表
type HumanNode struct {
	no            int        // 排序号
	name          string     // 姓名
	previousHuman *HumanNode //指向上一个节点
	nextHuman     *HumanNode // 指向下一个节点
}

// 对双向链表进行增加
func pushElement(headNode *HumanNode, newNode *HumanNode) {
	// 不可以随便改变头结点headNode，一旦改变整个链表就乱了，所以需要一个temp来间接调用链表的*HumanNode来存储
	tempNode := headNode

	// 由于不确定链表的长度所以用死循环遍历
	for {

		if tempNode.nextHuman == nil {
			break
		}

		tempNode = tempNode.nextHuman
	}

	// 跳出循环后，添加新元素，这时的tempNode.nextHuman指向链表的尾部
	tempNode.nextHuman = newNode

	// 与单链表不同的是，新节点的尾部的newNode.previousHuman又指向前一个节点
	newNode.previousHuman = tempNode
}

// 对链表进行插入，需要放入头节点，和要加入的节点作为参数
func insertElement(headNode *HumanNode, newNode *HumanNode) {
	// 不可以随便改变头结点headNode，一旦改变整个链表就乱了，所以需要一个temp来间接调用链表的*HumanNode来存储
	tempNode := headNode

	//标识符
	flag := true

	// 遍历链表来确定编号排序
	for {
		// 结合下面循环判断一直到末尾，如一开始就是末尾则直接跳出
		if tempNode.nextHuman == nil {
			break

			// 若不为nil，再去判断编号，如果链表的编号大于要插入链表元素的编号，结束循环，插入的元素在该节点前
		} else if tempNode.nextHuman.no > newNode.no {
			break

			// 若不为nil，且编号相等时，因为水浒传人物不重复，所以返回错误信息，如果需要重复，在上一个条件判断时加上等于号
			// 但插入元素的位置受大于小于号影响
		} else if tempNode.nextHuman.no == newNode.no {
			flag = false
			break
		}

		// 推动循环
		tempNode = tempNode.nextHuman
	}

	// 判断标识符
	if !flag {
		fmt.Println("人物已存在！")
		return
	} else {
		// 在编号有序时，这里赋值顺序不可反
		newNode.nextHuman = tempNode.nextHuman // 插入元素nextHuman连接tempNode原来的元素
		newNode.previousHuman = tempNode       // 插入元素previousHuman连接tempNode

		// 再判断tempNode.nextHuman是否为最后一个元素，因为最后一个元素的nextHuman为nil，无法存储previousHuman
		if tempNode.nextHuman != nil {
			// 如果不是，再让tempNode的下一个元素前节点连接插入元素
			tempNode.nextHuman.previousHuman = newNode
		}

		tempNode.nextHuman = newNode // 这里tempNode.nextHuman再更改连接到插入元素
	}
}

// 删除链表节点，传入编号
func delNode(headNode *HumanNode, no int) {
	// 不可以随便改变头结点headNode，一旦改变整个链表就乱了，所以需要一个temp来间接调用链表的*HumanNode来存储
	tempNode := headNode

	// 标识符
	flag := true

	// 遍历查询
	for {
		// 判断是否为末尾
		if tempNode.nextHuman == nil {
			break

			// tempNode.nextHuman下一个人物不等于nil，比较notempNode.nextHuman.no是否等于no
		} else if tempNode.nextHuman.no == no {
			// 标识，并跳出循环
			flag = false
			break
		}

		// 推动循环
		tempNode = tempNode.nextHuman
	}

	// 判断标识
	if !flag {
		// 直接把tempNode.nextHuman.no == no条件成立的下一个nextHuman赋值给当前tempNode.nextHuman
		tempNode.nextHuman = tempNode.nextHuman.nextHuman

		// 要判断tempNode.nextHuman，即tempNode.nextHuman.nextHuman是否为nil，如果为nil，则tempNode.nextHuman.previousHuman也不存在
		if tempNode.nextHuman != nil {
			tempNode.nextHuman.previousHuman = tempNode
		}

	} else {
		fmt.Println("找不到该编号对应的人物...")
		return
	}
}

// 展示链表，同样需要头结点然后遍历
func showLink(headNode *HumanNode) {
	// 不可以随便改变头结点headNode，一旦改变整个链表就乱了，所以需要一个temp来间接调用链表的*HumanNode来存储
	tempNode := headNode

	// 先判断链表是否为空
	if tempNode.nextHuman == nil {
		fmt.Println("链表是空的。。。")
		return
	}

	// 若不为空，则链表肯定至少有一个
	for {
		// 所以先打印第一个
		fmt.Printf("[%d %s] => ", tempNode.nextHuman.no, tempNode.nextHuman.name)

		// 如果还有继续重新赋值判断后从循环开始进行打印，有点类似于添加时的操作
		tempNode = tempNode.nextHuman

		// 再遍历是否还有
		if tempNode.nextHuman == nil {
			break
		}
	}

	fmt.Println()
}

// 反向展示链表
func reverseShowLink(headNode *HumanNode) {
	// 不可以随便改变头结点headNode，一旦改变整个链表就乱了，所以需要一个temp来间接调用链表的*HumanNode来存储
	tempNode := headNode

	// 先判断链表是否为空
	if tempNode.nextHuman == nil {
		fmt.Println("链表是空的。。。")
		return
	}

	// 先让tempNode指向最后一个节点
	for {

		if tempNode.nextHuman == nil {
			break
		}

		tempNode = tempNode.nextHuman
	}

	// 若不为空，则链表肯定至少有一个
	for {
		// 所以先打印第一个
		fmt.Printf("[%d %s] => ", tempNode.no, tempNode.name)

		// 如果还有继续重新赋值判断后从循环开始进行打印，有点类似于添加时的操作
		tempNode = tempNode.previousHuman

		// 再遍历是否还有
		if tempNode.previousHuman == nil {
			break
		}
	}

	fmt.Println()
}

func main() {
	// 定义头节点，头节点可以直接使用默认值
	headNode := &HumanNode{}

	// 定义人物加入到链表
	human_01 := &HumanNode{
		no:   1,
		name: "宋江",
	}

	human_02 := &HumanNode{
		no:   2,
		name: "吴用",
	}

	human_03 := &HumanNode{
		no:   3,
		name: "卢俊义",
	}

	// 添加
	pushElement(headNode, human_01)
	pushElement(headNode, human_02)
	pushElement(headNode, human_03)

	// 正向展示
	showLink(headNode)

	// 反向展示
	reverseShowLink(headNode)

	/*
		[1 宋江] => [2 吴用] => [3 卢俊义] =>
		[3 卢俊义] => [2 吴用] => [1 宋江] =>
	*/
}
