package main

import "fmt"

// 稀疏数组
func main() {
	// Two-dimensional array
	var tda [11][11]int

	tda[0][1] = 10
	tda[2][3] = 20
	tda[3][4] = 30
	tda[4][5] = 40
	tda[5][6] = 50

	// 保存非0个数
	count := 0
	for _, arr := range tda {
		//fmt.Printf("%d\t%d\n", index , value)
		for _, value := range arr {
			fmt.Printf("%d\t", value)
			if value > 0 {
				count++
			}
		}
		fmt.Println()
	}

	fmt.Println("--------------------------")

	// 初始化稀疏数组
	var sparseArray = make([][3]int, count+1)

	sparseArray[0][0] = len(tda)
	sparseArray[0][1] = len(tda[0])
	sparseArray[0][2] = count

	// 记录稀疏数组行数
	index := 0
	for row, arr := range tda {
		for column, value := range arr {
			if value > 0 {
				index++ // 从第二行开始记录具体的非0值的坐标
				sparseArray[index][0] = row
				sparseArray[index][1] = column
				sparseArray[index][2] = value
			}
		}
	}

	// 打印稀疏数组
	for _, array := range sparseArray {
		for _, value := range array {
			fmt.Printf("%d\t", value)
		}
		fmt.Println()
	}

	fmt.Println("--------------------------")

	// 根据稀疏数组第一行第一个元素初始化二维数组的行
	// Golang无法使用 make([][sparseArray[0][1]]int, ...)
	tda2 := make([][]int, sparseArray[0][0])

	// 初始化二维数组
	for i := range tda2 {
		tda2[i] = make([]int, sparseArray[0][1])
	}

	for row, array := range sparseArray {
		// 第一行保存的是二维数组的行和列以及非0的个数，跳过
		if row > 0 {
			tda2[array[0]][array[1]] = array[2]
		}
	}

	// 打印输出
	for _, array := range tda2 {
		for _, value := range array {
			fmt.Printf("%d\t", value)
		}
		fmt.Println()
	}

}
