package main

import (
	"container/list"
	"fmt"
)

// 二分查找，前提是数组为有序数组
func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 9}

	fmt.Println(BinarySearchByRecursion(arr, 0, len(arr)-1, 9))
	//l := BinarySearchByLoop(arr, 9)
	//fmt.Printf("%+v", l)

}

func BinarySearchByRecursion(arr []int, left, right, no int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2

	if no < arr[mid] {
		return BinarySearchByRecursion(arr, left, right-1, no)
	}

	if no > arr[mid] {
		return BinarySearchByRecursion(arr, left+1, right, no)
	}

	// 往左找
	temp := mid - 1
	for temp >= left && arr[temp] == no {
		fmt.Println("往左找找到一个#", temp)
		temp--
	}
	// 往右找
	temp = mid + 1
	for temp <= right && arr[temp] == no {
		fmt.Println("往右找找到一个#", temp)
		temp++
	}
	return mid
}

func BinarySearchByLoop(arr []int, no int) *list.List {
	l := list.New()
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		if no == arr[mid] {
			fmt.Println("找到了，索引#", mid)
			l.PushBack(mid)
			// 往左找
			temp := mid - 1
			for temp >= left && arr[temp] == no {
				l.PushBack(temp)
				temp--
			}
			// 往右找
			temp = mid + 1
			for temp <= right && arr[temp] == no {
				l.PushBack(temp)
				temp++
			}
			return l
		}
		if no < arr[mid] {
			right = mid - 1
		}
		if no > arr[mid] {
			left = mid + 1
		}
	}
	return nil
}
