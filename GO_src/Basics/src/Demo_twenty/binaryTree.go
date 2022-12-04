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
 *	Date         : 2022-08-21 09:56:55
 *	LastEditors  : Mr.Chan
 *	LastEditTime : 2022-08-21 10:59:25
 *	FilePath     : /vivianchan.cn/Basics/src/Demo_twenty/binaryTree.go
 *******************************************************/

package main

import "fmt"

/* ================================================================================================================================= */

// 采用递归回溯
// 前序遍历：先遍历根节点，然后遍历根节点左边子树，最后遍历右节点子树（遍历子树也按照这个顺序）。
// 中序遍历：先遍历根节点左边节点，再遍历根节点，最后遍历根节点右边节点。（遍历子树也按照这个顺序）。
// 后序遍历：先遍历左边节点包括左节点的子树，再遍历右边节点包括右节点的子树，最后遍历根节点（遍历子树也按照这个顺序）。

/* ================================================================================================================================= */

type human struct {
	no        int
	name      string
	leftNode  *human // 左节点或右节点连接的树
	rightNode *human // 右节点或右节点连接的树
}

// 结构体实例方法
func NewHuman(no int, name string) *human {
	return &human{
		no:   no,
		name: name,
	}
}

// 前序遍历
func previousOrder(Node *human) {
	if Node != nil {
		fmt.Printf("no=%d\tname=%s\n", Node.no, Node.name)
		previousOrder(Node.leftNode)
		previousOrder(Node.rightNode)
	}
}

// 中序遍历
func infixOrder(Node *human) {
	if Node != nil {
		previousOrder(Node.leftNode)
		fmt.Printf("no=%d\tname=%s\n", Node.no, Node.name)
		previousOrder(Node.rightNode)
	}
}

// 后序遍历
func epilogueOrder(Node *human) {
	if Node != nil {
		previousOrder(Node.leftNode)
		previousOrder(Node.rightNode)
		fmt.Printf("no=%d\tname=%s\n", Node.no, Node.name)
	}
}

/* ================================================================================================================================= */

func main() {
	//创建实例
	root := NewHuman(1, "宋江") // 根节点

	rootLeft := NewHuman(2, "吴用")   // 根节点的左节点
	leftLeft1 := NewHuman(3, "李逵")  // 左节点的左子节点
	leftRight1 := NewHuman(4, "花荣") // 左节点的右子节点

	rootRight := NewHuman(5, "卢俊义")   // 根节点的右节点
	rightLeft1 := NewHuman(6, "史进")   // 右节点的左子节点
	rightRight1 := NewHuman(7, "鲁智深") // 右节点的右子节点

	// 构建树结构

	// 左结构
	root.leftNode = rootLeft
	rootLeft.leftNode = leftLeft1
	rootLeft.rightNode = leftRight1

	// 右结构
	root.rightNode = rootRight
	rootRight.leftNode = rightLeft1
	rootRight.rightNode = rightRight1

	/*
			1
		   /  \
		  2    5
		 /\    /\
		3  4  6  7
	*/

	//前序
	previousOrder(root) // 1-2-3-4-5-6-7
	fmt.Println()

	// 中序
	infixOrder(root) // 2-3-4-1-5-6-7
	fmt.Println()

	// 后序
	epilogueOrder(root) // 2-3-4-5-6-7-1
	fmt.Println()
}
