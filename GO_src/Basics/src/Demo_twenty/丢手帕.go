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
 *	Date         : 2022-08-22 13:57:49
 *	LastEditors  : Mr.Chan
 *	LastEditTime : 2022-08-26 20:27:38
 *	FilePath     : /vivianchan.cn/Basics/src/Demo_twenty/丢手帕.go
 *******************************************************/

package main

import "fmt"

/* ================================================================================================================================= */

/*
	约瑟夫问题的原理：
		n个人围成一圈，选择一个人开始报数，每过m次则当前报数的人则会被踢出圈，然后挨着的下
		一个人作为又一轮游戏的第一个报数的人，这样重复最后只会剩下一个人。

	例如:
		n = 6, m = 5,陆续被踢出圈的人的编号为：5,4,6,2,3,1.
*/

/* ================================================================================================================================= */

//声明小孩
type children struct {
	no        int
	nextChild *children // 连接下一个小孩
}

// 实例一圈小孩的函数，返回头节点
func addChild(num int) (headNode *children) {
	// 输入的孩子数需要在，num>=1
	if num < 1 {
		fmt.Println("输入错误！")
		return
	}

	//初始化一个头节点
	firstChild := &children{}
	temp := firstChild // 辅助节点

	// 实例化小孩
	for i := 1; i <= num; i++ {

		// 实例
		child := &children{
			no:        i,
			nextChild: nil,
		}

		// 第一种情况：当i=1时
		if i == 1 {
			firstChild = child          // 将这个小孩赋值给第一个小孩
			temp = child                // 辅助节点也是
			temp.nextChild = firstChild // 辅助节点再指向第一个小孩形成环状

		} else {
			// 第二种情况：当i>=1时
			temp.nextChild = child      // 现在辅助节点就是第一个小孩，所以需要辅助节点的nextChild指向第二个child
			temp = child                // 然后再让辅助节点向前移动一位指向第二个child
			temp.nextChild = firstChild // 最后temp.nextChild再指向头节点形成环状
		}
	}
	return firstChild
}

// 展示链表
func showChild(headNode *children) (countChild int) {
	// 先判断是否为空链
	if headNode.nextChild == nil {
		fmt.Println("链表为空！")
		return
	}

	temp := headNode // 辅助节点
	countChild = 0   // 统计人数

	// 不为空时
	for {
		fmt.Printf("第 %d 个小孩 -> ", temp.no)

		// 到链表末尾
		if temp.nextChild == headNode {
			break
		}

		// 计算几个小孩
		countChild++

		// 推动循环
		temp = temp.nextChild
	}

	return countChild
}

// 开始游戏
func playGame(headNode *children, startNo int, intervalChild int) {
	// 先判断是否为空链
	if headNode.nextChild == nil {
		fmt.Println("链表为空！")
		return
	}

	// 辅助节点
	temp := headNode       // 用于遍历和踢出
	tailHelper := headNode // 主要用于踢出第一个小孩时的情况

	// 规定startNo范围
	// countChild := showChild(headNode)
	// if startNo < 1 || startNo > countChild {
	// 	fmt.Println("开始人选择错误")
	// 	return
	// }

	// 当链表中小孩只有一个时，结束
	if temp.nextChild == headNode {
		fmt.Println("链表只有一个小孩，最后一个小孩是：", temp.no)
		return
	}

	// 让tailHelper取最后一个
	for {
		if tailHelper.nextChild == headNode {
			break
		}
		tailHelper = tailHelper.nextChild
	}

	// 定位到开始游戏的小孩startNo
	for {
		if temp.no == startNo {
			break
		}
		temp = temp.nextChild
		tailHelper = tailHelper.nextChild
	}

	// 开始游戏
	for {
		// 定位到结束时的小孩
		for i := 1; i <= intervalChild; i++ {
			temp = temp.nextChild
			tailHelper = tailHelper.nextChild
		}

		// 踢出小孩
		fmt.Println("退出的小孩是：", temp.no)

		// 最后的小孩
		if temp == tailHelper {
			fmt.Println("最后的小孩是：", temp.no)
			break
		}

		// 推动循环，重新形成环状
		temp = temp.nextChild
		tailHelper.nextChild = temp
	}
}

/* ================================================================================================================================= */

func main() {
	// 将小孩围成圈
	headNode := addChild(5)

	// 展示链表
	showChild(headNode)

	fmt.Println()

	// 开始游戏
	playGame(headNode, 3, 2)
}

/*
第 1 个小孩 -> 第 2 个小孩 -> 第 3 个小孩 -> 第 4 个小孩 -> 第 5 个小孩 ->
退出的小孩是： 5
退出的小孩是： 3
退出的小孩是： 2
退出的小孩是： 4
退出的小孩是： 1
最后的小孩是： 1
*/
