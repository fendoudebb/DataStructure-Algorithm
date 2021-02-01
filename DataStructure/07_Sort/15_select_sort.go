package main

import "fmt"

func main() {
	arr := []int{100, 20, 300, 0}

	fmt.Println(arr)

	for i := 0; i < len(arr)-1; i++ {
		for j := 1 + i; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
		fmt.Println("第", i+1, "次排序#", arr)
	}

}

func SelectSortEvolution() {
	arr := []int{100, 20, 300, 0}

	// 假设arr[0]为最小值，从第二个元素开始遍历（计数都从1开始）
	for i := 1; i < len(arr); i++ {
		// 如果元素比第一个元素（第0位索引）大，则交换
		if arr[0] > arr[i] {
			temp := arr[0]
			arr[0] = arr[i]
			arr[i] = temp
		}
	}
	fmt.Println(arr)

	for i := 2; i < len(arr); i++ {
		// 如果元素比第二个元素（第1位索引）大，则交换
		if arr[1] > arr[i] {
			temp := arr[1]
			arr[1] = arr[i]
			arr[i] = temp
		}
	}
	fmt.Println(arr)

	for i := 3; i < len(arr); i++ {
		// 如果元素比第三个元素（第2位索引）大，则交换
		if arr[2] > arr[i] {
			temp := arr[2]
			arr[2] = arr[i]
			arr[i] = temp
		}
	}
	fmt.Println(arr)
}
