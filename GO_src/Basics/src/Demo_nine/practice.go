package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ----------------------------------------------------------------------------------------------------------

func randomArray() {
	/*
		练习一：
			随机生成范围在 1-100 内的 10 个整数，然后保存到数组，并倒序打印以及求出平均值、最大值、最大值下标、并查找是否包含 55 元素。

		分析：
			1.random生成随机数。
			2.声明一个数组，接收随机数。
			3.固定数组的一个数和数组的其他数比较，当大于时保持不变，小于时将其他数重新赋值给它(最大值和最小值只有一个的情况，多个最值时情况参考练习八)。
			4.使用冒泡排序，外层循环是数组总长度 -1 ，内层循环是数组长度 -1 后递减。
			5.遍历查询是否包含 55 。
	*/

	// 生成随机数种子
	rand.Seed(time.Now().UnixNano())

	// 给数组赋值随机数
	var array [10]int
	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(100) + 1
	}

	// 排序前
	fmt.Printf("排序前数组：%v\n", array)

	// 最大值、最大值下标，重点：需要固定一个数不动再和其他数进行比较。
	maxNum := 0
	for i := 1; i < len(array); i++ {
		if array[maxNum] < array[i] {
			maxNum = i
		}
	}
	fmt.Printf("最大值：%v\n最大值下标：%v\n", array[maxNum], maxNum)

	// 倒序：第二个循环的 if 判断需要视情况
	for i := 0; i < len(array)-1; i++ {
		flag := true
		for j := 0; j < len(array)-1-i; j++ {
			// 引用传递
			if array[j] < array[j+1] {
				temp := (&array)[j]
				array[j] = (&array)[j+1]
				array[j+1] = temp
				flag = false
			}
		}
		if flag {
			// 排序后
			fmt.Printf("排序后数组：%v\n", array)
			break
		}
	}

	// 求平均值、判断是否包含 55
	sum := 0.0
	flag := true
	for i := 0; i < len(array); i++ {
		// 如果没有flag为false
		if array[i] != 55 {
			flag = false
		}
		sum += float64(array[i])
	}
	fmt.Printf("该数组平均值为：%v\n", sum/float64(len(array)))
	if flag {
		fmt.Println("该数组包含元素 55 ！")
	} else {
		fmt.Println("该数组不包含元素 55 ！")
	}
}

// ----------------------------------------------------------------------------------------------------------

func insertElement(array *[6]int, userInsert int) {
	/*
		练习二：
			已知一个排序好（升序）的数组，要求插入一个元素，最后打印该数组，顺序依然是升序。

		分析（使用数组完成）：
			1.声明一个升序数组，用户输入插入元素。
			2.数组不可动态增长，所以要拷贝数组，使用 for-range 遍历拷贝到新数组。
			3.将用户输入元素赋值到新数组，因为遍历是从 0 开始，从第一个数组长度 -1 下标结束，所以只需要将用户输入赋值到新数组的最后一个下标。
			4.将新数组冒泡排序后打印。
			5.如果使用切片的方法，使用append函数即可。
	*/

	// 原数组
	fmt.Printf("原数组：%v\n", *array)

	// 声明另外一个数组为输入数组的长度 +1
	var array2 [len(array) + 1]int

	// 将输入数组拷贝到新数组
	for i, v := range array {
		array2[i] = v
	}

	// 然后把输入元素插入到新数组
	array2[len(array)] = userInsert
	fmt.Printf("插入元素后数组：%v\n", array2)

	// 最后进行排序
	for i := 0; i < len(array2)-1; i++ {
		flag := true
		for j := 0; j < len(array2)-i-1; j++ {
			if array2[j] > array2[j+1] {
				temp := array2[j]
				array2[j] = array2[j+1]
				array2[j+1] = temp
				flag = false
			}
		}
		if flag {
			break
		}
	}
	fmt.Printf("插入元素排序后数组：%v\n", array2)
}

// ----------------------------------------------------------------------------------------------------------

func twodimensionAroundClear() {
	/*
		练习三：
			声明一个 3 行 4 列的二维数组，逐个从键盘输入值，编写程序将四周的数据清 0 。

		分析：
			1.声明二维数组。
			2.用户遍历输入。
			3.遍历二维数组，当 i == 0 || i == 2 || j == 0 || j == 3 输出 0 ，其余照常输出 。
	*/

	// 声明二维数组
	var array [3][4]int

	// 输入
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			fmt.Printf("请输入第 %v 行，第 %v 个元素的值：", i+1, j+1)
			fmt.Scanln(&array[i][j])
		}
	}

	// 修改四周清零
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			if i == 0 || i == 2 || j == 0 || j == 3 {
				fmt.Print(0, " ")
			} else {
				fmt.Print(array[i][j], " ")
			}
		}
		// 当外围数组的第一个元素循环完毕后回车
		fmt.Println()
	}
}

// ----------------------------------------------------------------------------------------------------------

func twoDimensionSwap() {
	/*
		练习四：
			声明一个 4 行 4 列的二维数组，逐个从键盘输入值，然后将第 1 行和第 4 行的数据进行交换，将第 2 行和第 3 行的数据进行交换。

		分析：
			1.声明二维数组。
			2.用户遍历输入。
			3.当 i == 0 和 i == 3 交换，i == 1 和 i == 2 交换。
	*/

	// 声明
	var array [4][4]int

	// 输入
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			fmt.Printf("输入第 %v 行，第 %v 个数：", i+1, j+1)
			fmt.Scanln(&array[i][j])
		}
	}

	fmt.Println("交换前：")
	for _, v := range array {
		fmt.Println(v)
	}

	// 交换
	/*
		for i := 0; i < len(array); i++ {
			temp := (&array)[0][i]
			(&array)[0][i] = (&array)[3][i]
			(&array)[3][i] = temp
			temp2 := (&array)[1][i]
			(&array)[1][i] = (&array)[2][i]
			(&array)[2][i] = temp
		}
	*/
	temp, temp2 := (&array)[0], (&array)[1]
	(&array)[0], (&array)[1] = (&array)[3], (&array)[2]
	(&array)[3], (&array)[2] = temp, temp2

	fmt.Println("交换后：")
	for _, v := range array {
		fmt.Println(v)
	}
}

// ----------------------------------------------------------------------------------------------------------

func invertedOrder() {
	/*
		练习五：
			保存 1,3,5,7,9 这五个奇数到数组，并倒序打印。

		分析：
			1.声明一个数组。
			2.遍历赋值，奇数是 2n+1。
			3.两种排序方法。
	*/

	// 声明
	var array [5]int
	fmt.Printf("赋值前：%v\n", array)

	// 遍历赋值
	for i := 0; i < 5; i++ {
		(&array)[i] = (2 * i) + 1
	}
	fmt.Printf("赋值后：%v\n", array)

	// 直接倒序打印
	fmt.Print("直接倒序打印：")
	for j := 4; j >= 0; j-- {
		fmt.Print(array[j], " ")
	}
	fmt.Println()

	// 冒泡排序打印
	for i := 0; i < len(array)-1; i++ {
		flag := true
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] < array[j+1] {
				temp := (&array)[j]
				array[j] = (&array)[j+1]
				array[j+1] = temp
				flag = false
			}
		}
		if flag {
			break
		}
	}
	fmt.Printf("冒泡排序打印：%v", array)
}

// ----------------------------------------------------------------------------------------------------------

func arrSearch(array *[10]string, find string) {
	/*
		练习六：
			已知数组 array [10]string 保存了十个元素，要求查找是否存在 "AA" 并打印提示，如果有多个 "AA" ，找到对应下标。

		分析：
			1.声明数组。
			2.遍历查询,找到了记录下标，继续查找 continue，找不到直接 break。
	*/

	// 查询
	for i := 0; i < len(array); i++ {
		if array[i] == find {
			fmt.Printf("找到了！下标为：%v\n", i)
			continue
		} else {
			fmt.Println("找不到该元素！")
			break
		}
	}
}

// ----------------------------------------------------------------------------------------------------------

func binarySearch() {
	/*
		练习七：
			在 1-100 间随机生成 10 个整数，使用冒泡排序法进行排序，然后使用二分查找法，查找是否存在 90 这个数，并显示下标，如果没有则提示 “找不到该数！”。

		分析：
			1.生成随机数。
			2.冒泡排序法。
			3.二分查找法(不能查找重复元素下标)。
	*/

	// 生成随机种子
	rand.Seed(time.Now().UnixNano())

	// 数组接收赋值
	var array [10]int
	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(100) + 1
	}
	fmt.Printf("排序前：%v\n", array)

	// 冒泡排序法
	for i := 0; i < len(array)-1; i++ {
		flag := true
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				temp := (&array)[j]
				array[j] = (&array)[j+1]
				array[j+1] = temp
				flag = false
			}
		}
		if flag {
			break
		}
	}
	fmt.Printf("排序后：%v\n", array)

	// 二分查找法
	var (
		// 首下标和尾下标、标识、结果
		leftIndex  = 0
		rightIndex = len(array) - 1
		flag       = true
		result     = 0
	)
	// 左边元素不能大于右边元素时
	for leftIndex <= rightIndex {
		// 获取中间值下标
		middleIndex := (leftIndex + rightIndex) / 2
		// 当中间值大于查找值时说明查找值可能在下标 0 ~ （middleIndex-1）之间，即中间值左边部分。同理反之是中间值右边部分。等于时表示找到了。
		if array[middleIndex] > 90 {
			rightIndex = middleIndex - 1
		} else if array[middleIndex] < 90 {
			leftIndex = middleIndex + 1
		} else {
			result = middleIndex
			flag = false
			continue
		}
	}
	if flag {
		fmt.Println("找不到该数！")
	} else {
		fmt.Printf("存在 90 元素，下标是：%v\n", result)
	}
}

// ----------------------------------------------------------------------------------------------------------

func max_min(array [5]int) {
	/*
		练习八：
			编写一个函数可以接收一个数组，该数组有 5 个数，请找出最大的数和最小的数以及其对应的下标是多少。

		分析：
			1.声明数组。
			2.函数形参接收。
			3.是一个混乱排序，不然打印下标没意义（复制一个进行排序）。
			4.可能有多个最大值和最小值。
			6.直接遍历比较，如果和新数组的下标第 0 或 4 个元素相等输出。
	*/

	// 拷贝，排序
	var array2 = array
	for i := 0; i < len(array2)-1; i++ {
		flag := true
		for j := 0; j < len(array2)-1-i; j++ {
			if array2[j] > array2[j+1] {
				temp := array2[j]
				array2[j] = array2[j+1]
				array2[j+1] = temp
				flag = false
			}
		}
		if flag {
			break
		}
	}

	// 接收最大值和最小值以及下标
	var maxIndex, minIndex []int
	var max, min []int
	for i, _ := range array {
		if array[i] == array2[0] {
			minIndex = append(minIndex, i)
			min = append(min, array[i])
		}
	}
	for i, _ := range array {
		if array[i] == array2[4] {
			maxIndex = append(maxIndex, i)
			max = append(max, array[i])
		}
	}
	for i, _ := range min {
		fmt.Printf("最小值下标及值：%v：%v\n", minIndex[i], min[i])
	}
	for i, _ := range max {
		fmt.Printf("最大值下标及值：%v：%v\n", maxIndex[i], max[i])
	}
}

// ----------------------------------------------------------------------------------------------------------

func averageNum() {
	/*
		练习九：
			定义一个数组，并给出 8 个整数，求该数组中大于平均值的数的个数，和小于平均值数的个数。

		分析：
			1.声明数组。
			2.随机数赋值、累加。
			3.求出平均值 8 个数相加 / 8。
			4.遍历比较、计算个数。
	*/

	// 声明
	var (
		array        [8]int
		sum          int
		average      float64
		greaterCount int
		lessCount    int
	)

	// 随机赋值、累加
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(1000)
		sum += array[i]
	}
	average = float64(sum) / float64(len(array))
	fmt.Printf("数组：%v\n平均值：%.3f\n", array, average)

	// 遍历比较、统计个数
	for j := 0; j < len(array); j++ {
		if float64(array[j]) > average {
			greaterCount++
		} else if float64(array[j]) < average {
			lessCount++
		}
	}
	fmt.Printf("大于平均值的数的个数：%v\n小于平均值的数的个数：%v\n", greaterCount, lessCount)
}

// ----------------------------------------------------------------------------------------------------------

func divingScores() {
	/*
		练习十:
			跳水比赛，有 8 个评委打分。运动员的成绩是 8 个成绩去掉一个最高分，去掉一个最低分，剩下的 6 个分数的平均分就是最后得分。使用一维数组实现以下功能：
				1.请把打最高分的评委和最低分的评委找出来；（找出最高分和最低分的下标）
				2.找出最佳评委和最差评委。（最佳评委就是打分和最后得分最接近的评委，最差评委就是打分和最后得分相差最大的评委。）

		分析：
			1.长度为 8 的数组。
			2.用户输入 8 个整数，范围在 0 - 10 之间，小数精确到 0.1 。
			3.拷贝一个数组排序，不影响原来数组，可以找到原来对应下标的评委
			4.找出最大值和最小值下标（解决多个最大值和最小值）。
			5.去掉下标为 0 和 7 的分数。
			6.遍历剩余的元素累加求出平均分即得分。
			7.把 8 个分数分别减去得分后排序（会出现负数,把负数变正）：
				最大值 --> 最差评委
				最小值 --> 最佳评委
				因为也可能出现多个，所以还要把分差排序
	*/

	// 1.声明
	var (
		array [8]float64
		// 拷贝一份原数组
		array2 [8]float64

		// 解决 Scanf BUG
		/* stdin = bufio.NewReader(os.Stdin)
		stdin.ReadString('\n') */

		// 打分最高和最低的评委
		minIndex, maxIndex []int
		// 总分、平均分
		totalPoints   float64
		averageScores float64
		//分差,拷贝一份作为比较
		scoresGap  [8]float64
		scoresGap2 [8]float64
		// 最佳和最差评委
		scoresGapmax, scoresGapmin []int
	)

	// 评委打分
	for i := 0; i < len(array); i++ {
	Label:
		fmt.Printf("请第 %v 个评委打分(0 ~ 10)：", i+1)
		fmt.Scanln(&array[i])
		if array[i] > 10 || array[i] < 0 {
			fmt.Println("评分错误！请重新输入！！！")
			goto Label
		}
	}
	fmt.Printf("分数总览：%2.1f\n", array)

	// copy
	array2 = array

	// 排序
	for i := 0; i < len(array)-1; i++ {
		flag := true
		for j := 0; j < len(array)-1-i; j++ {
			if array2[j] > array2[j+1] {
				temp := array2[j]
				array2[j] = array2[j+1]
				array2[j+1] = temp
				flag = false
			}
		}
		if flag {
			break
		}
	}
	fmt.Printf("排序总览：%2.1f\n", array2)

	// 最大值和最小值下标(可能有多个)
	for i, _ := range array {
		if array[i] == array2[0] {
			minIndex = append(minIndex, i+1)
		} else if array[i] == array2[7] {
			maxIndex = append(maxIndex, i+1)
		}
	}
	fmt.Print("评分最低的评委有：")
	for _, v := range minIndex {
		fmt.Print(v, " ")
	}
	fmt.Println()
	fmt.Print("评分最高的评委有：")
	for _, v := range maxIndex {
		fmt.Print(v, " ")
	}
	fmt.Println()

	// 去掉下标 0 和下标 7 的元素，求最后得分，使用拷贝数组计算
	for i, _ := range array2 {
		array2[0] = 0
		array2[7] = 0
		totalPoints += array2[i]
	}
	averageScores = totalPoints / 6.0
	fmt.Printf("最后成绩：%2.1f\n", averageScores)

	// 因为分差会出现负值，所以必须把所有值都变正
	for i := 0; i < len(array); i++ {
		if array[i]-averageScores < 0 {
			temp := i
			scoresGap[temp] = -(array[temp] - averageScores)
		} else if array[i]-averageScores >= 0 {
			temp2 := i
			scoresGap[temp2] = array[temp2] - averageScores
		}
	}

	// 统计、比较分差的最大值、最小值，和上面差不多(可能有多个)
	// 先排序分差
	scoresGap2 = scoresGap // copy

	for i := 0; i < len(scoresGap2)-1; i++ {
		flag := true
		for j := 0; j < len(scoresGap2)-1-i; j++ {
			if scoresGap2[j] > scoresGap2[j+1] {
				temp := scoresGap2[j]
				scoresGap2[j] = scoresGap2[j+1]
				scoresGap2[j+1] = temp
				flag = false
			}
		}
		if flag {
			break
		}
	}

	for i, _ := range scoresGap {
		if scoresGap[i] == scoresGap2[0] {
			scoresGapmin = append(scoresGapmin, i+1)
		} else if scoresGap[i] == scoresGap2[7] {
			scoresGapmax = append(scoresGapmax, i+1)
		}
	}
	fmt.Print("最佳评委有：")
	for _, v := range scoresGapmin {
		fmt.Print(v, " ")
	}
	fmt.Println()
	fmt.Print("最差评委有：")
	for _, v := range scoresGapmax {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {

	// randomArray()

	/* array := [...]int{1, 20, 68, 100, 480, 1000}
	insertElement(&array, 3) */

	// twodimensionAroundClear()

	// twoDimensionSwap()

	// invertedOrder()

	/* var array2 [10]string = [10]string{"BB", "CC", "TT", "SDA", "FS", "BB", "SDW", "BB", "FG", "BB"}
	arrSearch(&array2, "AA") */

	// binarySearch()

	/* array3 := [...]int{7, 15, 101, 55, 7}
	max_min(array3) */

	// averageNum()

	// divingScores()

}
