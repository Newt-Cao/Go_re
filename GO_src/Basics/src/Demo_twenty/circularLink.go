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
 *	Date         : 2022-08-21 14:03:46
 *	LastEditors  : Mr.Chan
 *	LastEditTime : 2022-08-22 13:16:34
 *	FilePath     : /vivianchan.cn/Basics/src/Demo_twenty/circularLink.go
 *******************************************************/

package main

import "fmt"

/* ================================================================================================================================= */

type catNode struct {
	no       int
	name     string
	color    string
	nextNode *catNode
}

// 添加到链表最后
func cat_Insert(headNode *catNode, newCat *catNode) {
	// 声明辅助节点
	temp := headNode

	// 第一种情况：链表为空，直接加入，需要注意的是头节点还是一个单独的节点，只是头节点的信息为第一个猫的信息
	if temp.nextNode == nil {
		//把第一只猫的信息传给head
		headNode.no = newCat.no
		headNode.name = newCat.name
		headNode.color = newCat.color

		// 让头节点的下一个节点指向自己（头节点），用nextNode *catNode存储了自己
		headNode.nextNode = headNode
		return
	}

	// 第二种情况：链表不为空，获取链表最后一个节点，添加到链表最后
	for {
		if temp.nextNode == headNode {
			break
		}
		temp = temp.nextNode
	}

	// 执行添加
	temp.nextNode = newCat
	newCat.nextNode = headNode
}

// 删除指定节点，返回值是为了更改main函数的headNode
func cat_Del(headNode *catNode, no int) *catNode {
	//声明两个辅助节点
	temp := headNode
	helper := headNode // 主要处理删除的是多个节点时的第一个节点

	// 判断是否为空链，因为temp.nextNode如果不是空链必定会指向headNode
	if temp.nextNode == nil {
		fmt.Println("链表是空的！")
		return headNode
	}

	// 第一种情况：链表只有一只猫，即表头
	if headNode.nextNode == temp && temp.no == no {
		temp.nextNode = nil
		return headNode
	}

	// 若不是空链，有两个或以上节点
	for {
		//让helper取到末尾
		if helper.nextNode == headNode {
			break
		}
		helper = helper.nextNode
	}

	// 用temp找，helper删除
	flag := true
	for {
		if temp.nextNode == headNode {
			fmt.Printf("找不到编号 %d 猫的信息", no)
			break
		}
		if temp.no == no {

			// 第二种情况：多个节点时删除的是头节点
			if temp == headNode {
				headNode = headNode.nextNode // 让头节点指向它的下一个节点
			}
			helper.nextNode = temp.nextNode // 再让原本指向末尾的节点连接到headNode.nextNode
			flag = false
			break
		}

		// 移动，用来比较 id 和节点的 no
		temp = temp.nextNode
		helper = helper.nextNode
	}

	// flag是true时说名删除的节点不是第一个
	if flag {

		// 第三种情况：多个节点时删除头节点以外的节点
		if temp.nextNode.no == no {
			temp.nextNode = temp.nextNode.nextNode
		} else {
			fmt.Printf("找不到编号 %d 猫的信息", no)
			return headNode
		}
	}

	return headNode
}

// 展示环形单向链表
func cat_Show(headNode *catNode) {
	// 声明辅助节点
	temp := headNode

	// 判断是否为空链
	if temp.nextNode == nil {
		fmt.Println("空链")
		return
	}

	// 遍历链表
	for {
		fmt.Printf(" %d 编号的猫叫 %s ，颜色是 %s ——> ", temp.no, temp.name, temp.color)
		if temp.nextNode == headNode {
			break
		}
		temp = temp.nextNode
	}

	fmt.Println()
}

/* ================================================================================================================================= */

func main() {
	// 头节点
	headNode := &catNode{}

	// 其他节点
	cat1 := &catNode{
		no:    1,
		name:  "tom",
		color: "black",
	}
	cat2 := &catNode{
		no:    2,
		name:  "jack",
		color: "yellow",
	}
	cat3 := &catNode{
		no:    3,
		name:  "tiga",
		color: "blue",
	}
	cat4 := &catNode{
		no:    4,
		name:  "devil",
		color: "red",
	}

	// 加入
	cat_Insert(headNode, cat1)
	cat_Insert(headNode, cat2)
	cat_Insert(headNode, cat3)
	cat_Insert(headNode, cat4)

	// 展示
	cat_Show(headNode)

	// 删除2
	headNode = cat_Del(headNode, 2)

	// 展示
	cat_Show(headNode)

	// 删除1
	headNode = cat_Del(headNode, 1)

	// 展示
	cat_Show(headNode)
}
