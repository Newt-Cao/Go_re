package main

import "fmt"

/*
案例：在单链表实现水浒传英雄顺序的增删改查
*/

// 声明结构体操作链表
type HumanNode struct {
	no        int        // 排序号
	name      string     // 姓名
	nextHuman *HumanNode // 指向下一个节点
}

// 对链表进行增加，需要放入头节点，和要加入的节点作为参数
func pushElement(headNode *HumanNode, newNode *HumanNode) {
	// 不可以随便改变头结点headNode，一旦改变整个链表就乱了，所以需要一个temp来间接调用链表的*HumanNode来存储
	tempNode := headNode

	// 由于不确定链表的长度所以用死循环遍历
	for {
		/*
			思路：当nextHuman为空时，说明是链表的尾部，退出循环，tempNode最后是上一次循环的
				 tempNode.nextHuman，这样就不会改变headNode，同时也可以不断遍历

			理解：假设不等于nil，就会把新的tempNode.nextHuman重新赋值给tempNode，便可以达到不断遍历知道链表尾部

			误区：break后就不会执行tempNode = tempNode.nextHuman，所以最后的tempNode就是上一次执行的
				 tempNode = tempNode.nextHuman，并且包含一个空的tempNode.nextHuman
		*/
		if tempNode.nextHuman == nil {
			break
		}

		tempNode = tempNode.nextHuman
	}

	// 跳出循环后，添加新元素，这时的tempNode.nextHuman指向链表的尾部
	tempNode.nextHuman = newNode
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
		newNode.nextHuman = tempNode.nextHuman // 用插入元素的nextHuman连接大于它的链表部分
		tempNode.nextHuman = newNode           // 再用原来连接大于它链表部分的节点连接插入元素
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
	} else {
		fmt.Println("找不到该编号对应的人物...")
		return
	}
}

// 更改节点数据
func changeNodeData(headNode *HumanNode, oldno int) {
	var (
		// 不可以随便改变头结点headNode，一旦改变整个链表就乱了，所以需要一个temp来间接调用链表的*HumanNode来存储
		tempNode = headNode
		// 标识符
		flag = true
		// 需要更改元素的编号
		newno int
		// 更改后的姓名
		newname string
	)

	// 用户输入
	fmt.Println("输入新排名：")
	fmt.Scanln(&newno)
	fmt.Println("输入新姓名：")
	fmt.Scanln(&newname)

	// 实例化一个
	for {
		// 同删除原理相似，查找对应编号的元素
		if tempNode.nextHuman == nil {
			break
		} else if tempNode.nextHuman.no == oldno {
			flag = false
			break
		}
		tempNode = tempNode.nextHuman
	}

	// flag=false时把更改的值填入
	if !flag {
		tempNode.nextHuman.no = newno
		tempNode.nextHuman.name = newname
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

	// 使用函数添加到链表
	pushElement(headNode, human_01)
	pushElement(headNode, human_02)
	pushElement(headNode, human_03)

	// 展示链表
	showLink(headNode) // [1 宋江] => [2 吴用] => [3 卢俊义]

	// 使用函数插入到链表
	insertElement(headNode, human_01)
	insertElement(headNode, human_03)
	insertElement(headNode, human_02)

	// 展示链表
	showLink(headNode) // [1 宋江] => [2 吴用] => [3 卢俊义]

	// 使用函数删除链表元素
	delNode(headNode, human_01.no)

	// 展示链表
	showLink(headNode) // [2 吴用] => [3 卢俊义]

	// 使用函数更改链表元素
	changeNodeData(headNode, human_03.no)

	// 展示链表
	showLink(headNode) // [2 吴用] => [4 林冲]
}
