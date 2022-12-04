package main

import (
	"fmt"
)

func application() {
	/*
		定义二维数组，用于保存三个班，每个班五名同学的成绩，并求出每个班级平均分，和所有班级总平均分。
	*/

	// 声明一个二维数组
	var score [3][5]float64
	// 遍历输入成绩
	for i := 0; i < len(score); i++ {
		for j := 0; j < len(score[i]); j++ {
			fmt.Printf("请输入第 %v 个班，第 %v 名同学的成绩：", i+1, j+1)
			fmt.Scanln(&score[i][j])
		}
	}
	// 统计总分
	totalScore := 0.0
	// 统计人数
	peopleCount := 0
	// 遍历分数累加
	for i := 0; i < len(score); i++ {
		// 统计每个班总分
		classTotal := 0.0
		for j := 0; j < len(score[i]); j++ {
			classTotal += score[i][j]
			peopleCount++
		}
		totalScore += classTotal
		// 计算
		fmt.Printf("第 %v 班的平均分是：%v\n", i+1, classTotal/float64(len(score)))
	}
	// 计算
	fmt.Printf("总平均分是：%v", totalScore/float64(peopleCount))
}

func main() {
	// 声明二维数组，第一个是横，第二个是列。
	var array [4][6]int
	// var array [...][6] = [...][6]int{{初始化},{...},{...}}
	// var array [4][6]int = [4][6]int{{0,0,0,0,0,0},...} 声明+初始化
	// var array [4][6]int = [4][6]int{1:{0,0,0,0,0,0},2:...} 声明+外围数组指定索引元素初始化
	// var array [4][6]int = [4][6]int{1:{1:0,2:0,0,0,0,0},2:...} 声明+内层数组指定索引元素初始化

	fmt.Println(len(array)) // 4，len得出的是外围数组的长度
	fmt.Println(len(array[0]))

	// 给多维数组指定元素赋值。
	array[1][2] = 1
	array[2][1] = 2
	array[2][3] = 3

	// 遍历多维数组，以元素的方式打印
	/*
		0 0 0 0 0 0
		0 0 1 0 0 0
		0 2 0 3 0 0
		0 0 0 0 0 0
	*/
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			fmt.Print(array[i][j], " ")
		}
		fmt.Println()
	}

	array2 := [2][2]int{{1, 2}, {3, 4}}
	for i, v1 := range array2 {
		for j, v2 := range v1 {
			fmt.Printf("[%v][%v]=%v\n", i, j, v2)
		}
	}
	/*
		[0][0]=1
		[0][1]=2
		[1][0]=3
		[1][1]=4
	*/

	application()
}
