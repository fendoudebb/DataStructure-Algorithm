package main

import (
	"log"
	"math/rand"
)

func main() {
	//ShellSortExchangeEvolution()
	// [8,3] [9,5] [1,4] [7,6] [2,0]

	// [3 5 1 6 0 8 9 4 7 2]
	// [3,1,0,9,7] [5,6,8,4,2]

	// [0,2,1,4,3,5,7,6,9,8]

	//Exchange()

	//arr := []int{8, 9, 1, 7, 2, 3, 5, 4, 6, 0}

	arr := make([]int, 80000)
	for i := 0; i < 80000; i++ {
		arr[i] = rand.Intn(80000)
	}
	log.Println("start")

	for step := len(arr) / 2; step > 0; step /= 2 {
		for i := step; i < len(arr); i++ {
			j := i
			insertVal := arr[i]
			if arr[j] < arr[j-step] {
				for j-step >= 0 && insertVal < arr[j-step] {
					arr[j] = arr[j-step]
					j -= step
				}
				arr[j] = insertVal
			}
		}
		//log.Println(arr)
	}
	log.Println("end")
}

func Exchange() {
	arr := []int{8, 9, 1, 7, 2, 3, 5, 4, 6, 0}

	for step := len(arr) / 2; step > 0; step /= 2 {
		for i := step; i < len(arr); i++ {
			for j := i - step; j >= 0; j -= step {
				if arr[j] > arr[j+step] {
					temp := arr[j]
					arr[j] = arr[j+step]
					arr[j+step] = temp
				}
			}
		}
		log.Println(arr)
	}
}

// 交换法
func ShellSortExchangeEvolution() {
	arr := []int{8, 9, 1, 7, 2, 3, 5, 4, 6, 0}
	log.Println(arr)
	//step := 5 // len(arr)/2
	for i := 5; i < len(arr); i++ {
		log.Println("step=5, iiiiiiiiii=", i, arr)
		for j := i - 5; j >= 0; j -= 5 {
			if arr[j] > arr[j+5] {
				temp := arr[j]
				arr[j] = arr[j+5]
				arr[j+5] = temp
			}
			log.Println("step=5, j=", j, arr)
		}
	}

	log.Println(arr)

	for i := 2; i < len(arr); i++ {
		/*if arr[i-2] > arr[i] {
			temp := arr[i-2]
			arr[i-2] = arr[i]
			arr[i] = temp
		}*/
		log.Println("step=2, iiiiiiiiii=", i, arr)
		for j := i - 2; j >= 0; j -= 2 {
			if arr[j] > arr[j+2] {
				temp := arr[j]
				arr[j] = arr[j+2]
				arr[j+2] = temp
			}
			log.Println("step=2, j=", j, arr)
		}
	}

	log.Println(arr)

	for i := 1; i < len(arr); i++ {

		log.Println("step=1, iiiiiiiiii=", i, arr)
		for j := i - 1; j >= 0; j -= 1 {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
			log.Println("step=1, j=", j, arr)
		}
	}

	log.Println(arr)
}
