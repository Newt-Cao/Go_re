package main

import "fmt"

/*
案例：二维数组模拟通路寻找
思路：
	1.用二维数组绘制二维地图
	2.用 1 代表墙，0 代表未探索，2 代表通路，3 代表死路
	3.先使用下右上左的方向进行探索
	4.不断用递归调用直到地图终点等于2
问题：最优路径（未解决）
*/

// 坐标探索，传入地图和x、y坐标，初始值也就是出发点是 1,1
func coordinate(labyMap *[8][7]int, x, y int) bool {
	// 先判断是否已经到达终点，同时也是递归结束条件
	if labyMap[6][5] == 2 {
		return true

	} else {

		// 如果未到达终点，先判断是否可探索，若可探索则开始按照规则下右上左探索该坐标是否可为通路
		if labyMap[x][y] == 0 {

			// 第一次执行先把出生点设为2通路，不然无法开始，但往后的每次等于2代表的是通路
			labyMap[x][y] = 2

			// 向下探索
			if coordinate(labyMap, x+1, y) {
				return true

				// 向右
			} else if coordinate(labyMap, x, y+1) {
				return true

				// 向上
			} else if coordinate(labyMap, x-1, y) {
				return true

				// 向左
			} else if coordinate(labyMap, x, y-1) {
				return true

				// 都不通则判定为死路
			} else {
				labyMap[x][y] = 3
				return false
			}

		} else {

			// 如果不可探索直接返回false
			return false

		}
	}
}

func main() {
	// 绘制地图
	var labyMap [8][7]int

	for i := 0; i < 7; i++ {
		for j := 0; j < 8; j++ {
			labyMap[0][i] = 1
			labyMap[7][i] = 1
			labyMap[j][0] = 1
			labyMap[j][6] = 1
		}
	}

	labyMap[3][1] = 1
	labyMap[3][2] = 1

	// 展示地图
	for i := 0; i < len(labyMap); i++ {
		for j := 0; j < len(labyMap[i]); j++ {
			fmt.Print(labyMap[i][j], " ")
		}
		fmt.Println()
	}

	coordinate(&labyMap, 1, 1)
	fmt.Println()

	// 展示地图
	for i := 0; i < len(labyMap); i++ {
		for j := 0; j < len(labyMap[i]); j++ {
			fmt.Print(labyMap[i][j], " ")
		}
		fmt.Println()
	}

}

/*
1 1 1 1 1 1 1
1 0 0 0 0 0 1
1 0 0 0 0 0 1
1 1 1 0 0 0 1
1 0 0 0 0 0 1
1 0 0 0 0 0 1
1 0 0 0 0 0 1
1 1 1 1 1 1 1

1 1 1 1 1 1 1
1 2 0 0 0 0 1
1 2 2 2 0 0 1
1 1 1 2 0 0 1
1 0 0 2 0 0 1
1 0 0 2 0 0 1
1 0 0 2 2 2 1
1 1 1 1 1 1 1
*/
