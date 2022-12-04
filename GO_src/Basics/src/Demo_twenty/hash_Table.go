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
 *	Date         : 2022-08-18 11:26:11
 *	LastEditors  : Mr.Chan
 *	LastEditTime : 2022-08-21 07:54:46
 *	FilePath     : /vivianchan.cn/Basics/src/Demo_twenty/hash_Table.go
 *******************************************************/

package main

import "fmt"

/*
案例：
	情景：当公司新人入职时，可以迅速将该员工信息（id、姓名等...）加入，当输入该员工id时，获得其所有信息。
	要求：
		1.不使用数据库和编译器内置函数和关键字；
		2.节省内存，速度快；
		3.加入员工信息时，id存在于链表的形式是从低到高。（解决链表前中后的问题）
	思路：
		1.使用链表实现哈希表，可带表头也可不带表头。
*/

/* ================================================================================================================================= */

// 声明员工结构体，指向单个员工
type staff struct {
	id       int    // 员工id
	name     string // 员工姓名
	nextNode *staff // 指向下一个员工
}

/* ================================================================================================================================= */

// 声明员工链表结构体，指向单个链表
type staffLink struct {
	head *staff // 链表，不含表头
}

// 定义插入员工函数
func (sl *staffLink) staffLink_Insert(sta *staff) {
	// 声明两个辅助节点
	temp := sl.head                // 默认为头节点
	var temp_previous *staff = nil // 默认为temp的前一个节点

	// 先判断是否为空链，若为空链直接用head接收
	if temp == nil {
		sl.head = sta
		return
	}

	for {
		/*
			若不是空链，则分三种情况插入链表（在循环体内执行）：
				1.当只有一个节点（即头结点指指向一个员工时）或插入位置是头结点时，需要插入到该节点前。
				2.当有多个节点需要插入到指定位置，且保证员工id是从小到大的。
				3.当插入的员工节点id是该链表最大的也就是插入链表末尾时。
		*/

		// 不为空链时，判断链表已存的id和将要插入链表的id，假设id唯一
		if temp != nil {

			// 第一种：第一次循环temp还是等于head，直接判断
			if temp == sl.head && temp.id >= sta.id {
				sta.nextNode = temp // 让插入的下一个节点指向头节点下一个员工
				sl.head = sta       // 头节点指向插入节点
				return
			}

			// 第二种：
			if temp.id > sta.id {
				break
			}

			// 推动循环
			temp_previous = temp // 不断指向temp
			temp = temp.nextNode // 不断指向temp下一个节点

			// 第三种：
		} else if temp == nil {
			break
		}
	}

	temp_previous.nextNode = sta // temp的nextNode指向插入员工
	sta.nextNode = temp          // 插入员工的nextNode指向temp，因为temp实际上就是指向下一个员工
	return
}

// 定义所有员工展示函数
func (sl *staffLink) staffLink_Show(i int) {
	// 先判断是否为空链
	if sl.head == nil {
		fmt.Printf("%d 是空链\n", i)
		return
	}

	// 若不是，声明一个辅助节点
	temp := sl.head

	// 遍历所有员工
	for {
		// 遍历到最后退出
		if temp == nil {
			break
		} else {
			// 否则按下面一直打印
			fmt.Printf("第 %d 条链 员工id:%d 名字:%s ——>", i, temp.id, temp.name)
			temp = temp.nextNode
		}
	}
	fmt.Println()
}

// 定义员工信息查找函数
func (sl *staffLink) staffLink_FindNO(id int) (sta *staff) {
	// 先判断是否为空链
	if sl.head == nil {
		fmt.Printf("该id对应的是空链\n")
		return
	}

	// 声明辅助节点
	temp := sl.head

	// 若不是，且因为是单向链表所以for循环遍历
	for {
		// 找到了
		if temp != nil && temp.id == id {
			return temp

			// 找不到
		} else if temp == nil {
			return nil
		}

		// 推动循环
		temp = temp.nextNode
	}
}

// 定义员工信息删除函数
func (sl *staffLink) staffLink_Del(id int) bool {
	// 先判断是否为空链
	if sl.head == nil {
		fmt.Printf("该id对应的是空链\n")
		return false
	}

	// 声明两个辅助节点
	temp := sl.head                // 默认为头节点
	var temp_previous *staff = nil // 默认为temp的前一个节点

	// 若不为空，查找id
	for {
		// 同样分三种情况
		if temp != nil {

			// 找到时
			if temp.id == id {

				// 第一种：删除第一个节点
				if temp == sl.head {
					sl.head = temp.nextNode
					return true
				}

				// 第二种：在链表员工之间，要删除的temp找到了
				if temp != sl.head {
					temp_previous.nextNode = temp.nextNode // 把要删除的temp前一个节点连接到temp下一个节点
					return true
				}

			}

			// 推动循环
			temp_previous = temp // 不断指向temp
			temp = temp.nextNode // 不断指向temp下一个节点

			// 第三种：找不到
		} else if temp == nil {
			return false
		}
	}
}

/* ================================================================================================================================= */

// 声明员工链表散列表，按下标接收指定员工，包含所有员工链表
type staff_hashTable struct {
	all_Link [7]staffLink // 所有链表用数组接收
}

// 定义员工实例创建函数
func (sh *staff_hashTable) instance(id int, name string) (sta *staff) {
	return &staff{
		id:   id,
		name: name,
	}
}

// 定义插入链表函数
func (sh *staff_hashTable) hashTable_Insert(sta *staff) {
	// 通过员工id取模后放入指定链表
	linkNo := sta.id % 7

	// 加入对应链表
	sh.all_Link[linkNo].staffLink_Insert(sta)
}

//定义所有链表展示函数
func (sh *staff_hashTable) hashTable_Show() {
	for i := 0; i < len(sh.all_Link); i++ {
		// 每一个i获取到对应的一个staffLink实例，所以可以调用其绑定的方法
		sh.all_Link[i].staffLink_Show(i)
	}
}

//定义查找指定员工信息函数
func (sh *staff_hashTable) hashTable_FindNO(id int) (linkNo int, sta *staff) {
	// 传入id取模后再选择查找的链表
	linkNo = id % 7

	// 找到指定链表调用函数
	sta = sh.all_Link[linkNo].staffLink_FindNO(id)
	if sta == nil {
		fmt.Println("找不到该员工信息！\n")
		return
	}

	return
}

// 定义删除指定员工函数
func (sh *staff_hashTable) hashTable_Del(id int) {
	// 传入id取模后再选择查找的链表
	linkNo := id % 7

	// 找到指定链表调用函数
	staDel := sh.all_Link[linkNo].staffLink_Del(id)
	if !staDel {
		fmt.Println("删除失败！\n")
	} else {
		fmt.Println("删除成功！\n")
	}
}

/* ================================================================================================================================= */

func main() {
	// 声明变量
	var (
		keyNum          string          // 用户输入选项
		id              int             // 录入员工id
		name            string          // 录入员工姓名
		staff_hashtable staff_hashTable // 实例化一个staff_hashTable
	)

	// 打印菜单
	for {
		fmt.Println(`<——员工信息路通——>
  i 录入
  d 删除
  f 查询
  s 展示`)
		fmt.Print("请输入：")
		fmt.Scanln(&keyNum)

		switch keyNum {
		case "i":
			// 信息录入
			fmt.Print("输入员工id：")
			fmt.Scanln(&id)
			fmt.Print("输入员工姓名：")
			fmt.Scanln(&name)

			// 插入到哈希表
			staff_hashtable.hashTable_Insert(staff_hashtable.instance(id, name))

			// id清零
			id = 0

		case "d":
			// 删除员工信息
			// 用户输入
			fmt.Print("输入员工id：")
			fmt.Scanln(&id)

			// 删除
			staff_hashtable.hashTable_Del(id)

			// id清零
			id = 0

		case "f":
			// 员工信息查询
			// 用户输入
			fmt.Print("输入员工id：")
			fmt.Scanln(&id)

			// 查询
			linkNo, sta := staff_hashtable.hashTable_FindNO(id)

			// 打印
			fmt.Printf("查询的员工在第 %d 条链表，员工id：%d，员工姓名：%s。\n", linkNo, sta.id, sta.name)

			// id清零
			id = 0

		case "s":
			// 展示所有链表
			staff_hashtable.hashTable_Show()

		default:
			fmt.Println("输入错误！\n")
			continue
		}
	}
}
