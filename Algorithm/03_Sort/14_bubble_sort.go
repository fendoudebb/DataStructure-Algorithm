package main

import "fmt"

func main() {
	//Evolution()

	arr := []int{-2, 3, 9, -1, 10}
	fmt.Println(arr)

	var temp = 0

	for i := 0; i < len(arr)-1; i++ {
		exchange := false
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				temp = arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
				exchange = true
			}
		}
		fmt.Println("第",i+1, "次排序结果#",arr)
		if !exchange {
			break
		}
	}
}

func Evolution() {
	arr := []int{3, 9, -1, 10, -2}
	fmt.Println(arr)

	var temp = 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			temp = arr[i+1]
			arr[i+1] = arr[i]
			arr[i] = temp
		}
	}

	fmt.Println(arr)

	for i := 0; i < len(arr)-1-1; i++ {
		if arr[i] > arr[i+1] {
			temp = arr[i+1]
			arr[i+1] = arr[i]
			arr[i] = temp
		}
	}

	fmt.Println(arr)

	for i := 0; i < len(arr)-1-2; i++ {
		if arr[i] > arr[i+1] {
			temp = arr[i+1]
			arr[i+1] = arr[i]
			arr[i] = temp
		}
	}

	fmt.Println(arr)

	for i := 0; i < len(arr)-1-3; i++ {
		if arr[i] > arr[i+1] {
			temp = arr[i+1]
			arr[i+1] = arr[i]
			arr[i] = temp
		}
	}

	fmt.Println(arr)
}
