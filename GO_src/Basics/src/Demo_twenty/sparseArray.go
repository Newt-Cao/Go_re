package main

import (
	"fmt"
)

// 思想：把结构体元素存入切片中，每一个结构体的实例就是一个有效数据（切片的本质是可变数组）
type DataNode struct {
	row    int // 行
	column int // 列
	value  int // 值
}

func main() {
	// 使用二维数组存储起码，数据量大，效率低
	var dimensional_Array [11][11]int // 棋盘
	dimensional_Array[1][2] = 1       // 黑棋
	dimensional_Array[2][3] = 2       // 白棋
	for _, r := range dimensional_Array {
		for _, c := range r {
			fmt.Printf("%d  ", c)
		}
		fmt.Println()
	}
	/*
		0  0  0  0  0  0  0  0  0  0  0
		0  0  1  0  0  0  0  0  0  0  0
		0  0  0  2  0  0  0  0  0  0  0
		0  0  0  0  0  0  0  0  0  0  0
		0  0  0  0  0  0  0  0  0  0  0
		0  0  0  0  0  0  0  0  0  0  0
		0  0  0  0  0  0  0  0  0  0  0
		0  0  0  0  0  0  0  0  0  0  0
		0  0  0  0  0  0  0  0  0  0  0
		0  0  0  0  0  0  0  0  0  0  0
		0  0  0  0  0  0  0  0  0  0  0
	*/

	// 转成稀疏数组（标准的稀疏数组包含行列总数据规模），压缩该数组

	// 稀疏数组
	var SparseArray []DataNode

	// 行列总数据规模，记录棋盘
	dataNode := DataNode{
		row:    11,
		column: 11,
		value:  0,
	}

	// 存入切片
	SparseArray = append(SparseArray, dataNode)

	// 存储数据，先遍历原二维数组
	for i, r := range dimensional_Array {
		for j, c := range r {
			// 把大于0的数据取出存放
			if c != 0 {
				// 记录
				dataNode = DataNode{
					row:    i,
					column: j,
					value:  c,
				}

				// 存入切片
				SparseArray = append(SparseArray, dataNode)
			}
		}
	}
	/*
		{11 11 0}
		{1 2 1}
		{2 3 2}
	*/
	fmt.Println()

	var dimensional_Array2 [11][11]int
	// 再转回去原始数组
	for i, n := range SparseArray {
		if i == 0 {
			continue
		}
		dimensional_Array2[n.row][n.column] = n.value
	}

	for _, r := range dimensional_Array2 {
		for _, c := range r {
			fmt.Printf("%d  ", c)
		}
		fmt.Println()
	}
}
