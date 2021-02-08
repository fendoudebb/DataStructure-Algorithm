package main

import (
	"log"
)

func main() {
	//QuickSortEvolutionDemo()

	//arr := []int{-9, 78, 0, 23, -567, 70}
	arr := []int{23,46,0,8,11,18}
	//arr := make([]int, 800000)
	//for i := 0; i < 800000; i++ {
	//	arr[i] = rand.Intn(800000)
	//}
	log.Println("start")
	QuickSort(arr, 0, len(arr)-1)

	log.Println("end")
	//log.Println(arr)

}


func QuickSort(arr []int, left, right int) {
	// 中轴值
	var pivot = arr[(left+right)/2]
	var l = left  // 数组第0位
	var r = right // 数组最后一位

	for l < r {
		for arr[l] < pivot { // 找到中轴左边比中轴值大的，暂停循环
			l++ // 往中轴靠近，右移一位
		}

		for arr[r] > pivot { // 找到中轴右边比中轴值小的，暂停循环
			r-- // 往中轴靠近，左移一位
		}

		if l >= r {
			break
		}

		// 左边的值和右边的值交换
		temp := arr[l]
		arr[l] = arr[r]
		arr[r] = temp

		if arr[l] == pivot {
			r--
		}

		if arr[r] == pivot {
			l++
		}
	}

	if l == r {
		l++
		r--
	}

	// 左边
	if left < r {
		QuickSort(arr, left, r)
	}

	// 右边
	if right > l {
		QuickSort(arr, l, right)
	}

}

func QuickSortEvolutionDemo() {
	arr := []int{-9, 78, 0, 23, -567, 70}

	QuickSortEvolution(arr, 0, len(arr)-1)

	// 左边排序
	QuickSortEvolution(arr, 0, (0+len(arr)-1)/2-1)

	// 右边排序
	QuickSortEvolution(arr, (0+len(arr)-1)/2+1, len(arr)-1)
}

func QuickSortEvolution(arr []int, left, right int) {
	// 中轴值
	var pivot = arr[(left+right)/2]
	var l = left  // 数组第0位
	var r = right // 数组最后一位

	for l < r {
		for arr[l] < pivot { // 找到中轴左边比中轴值大的，暂停循环
			l++ // 往中轴靠近，右移一位
		}

		for arr[r] > pivot { // 找到中轴右边比中轴值小的，暂停循环
			r-- // 往中轴靠近，左移一位
		}

		if l >= r {
			break
		}

		// 左边的值和右边的值交换
		temp := arr[l]
		arr[l] = arr[r]
		arr[r] = temp
	}

}
