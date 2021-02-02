package main

import "log"

func main() {
	arr := []int{10, 20, 5, 0}
	for i := 1; i < len(arr); i++ {
		j := i
		insertVal := arr[i] // 一定要使用临时变量接收，如果在循环中用 arr[i]获取，其实值已经变为移动后的值了
		// 当前元素与前一个元素相比
		for j-1 >= 0 && insertVal < arr[j-1] {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = insertVal
		log.Println(arr)
	}


}

func InsertSortEvolution()  {
	arr := []int{10, 20, 5, 0}


	// 将 [10] 看作有序数组，[20, 5, 0] 看作无序数组
	insertVal := arr[1]
	j := 1 - 1
	for j >= 0 && insertVal < arr[j] {
		arr[j+1] = arr[j]
		j--
	}
	arr[j+1] = insertVal
	log.Println(arr)


	// 将 [10, 20] 看作有序数组，[5, 0] 看作无序数组
	insertVal = arr[2]
	j = 2 - 1
	for j >= 0 && insertVal < arr[j] {
		arr[j+1] = arr[j]
		j--
	}
	arr[j+1] = insertVal
	log.Println(arr)


	// 将 [5, 10, 20] 看作有序数组，[0] 看作无序数组
	insertVal = arr[3]
	j = 3 - 1
	for j >= 0 && insertVal < arr[j] {
		arr[j+1] = arr[j]
		j--
	}
	arr[j+1] = insertVal
	log.Println(arr)
}

func Demo() {
	arr := []int{10, 20, 5, 0}
	for i := 1; i < len(arr); i++ {
		j := i
		for j-1 >= 0 && arr[j] < arr[j-1] {
			temp := arr[j-1]
			arr[j-1] = arr[j]
			arr[j] = temp
			j--
		}
		log.Println(arr)
	}
}
