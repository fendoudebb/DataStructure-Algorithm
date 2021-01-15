package main

import (
	"fmt"
)

func main() {

	circularLinkedList := &CircularLinkedList{}
	circularLinkedList.Init(5)
	circularLinkedList.Print()

	circularLinkedList.Pop(5, 1, 1)
	//circularLinkedList.Pop2(5, 1, 2)
}

type CircularLinkedList struct {
	first *CircularNode
}

type CircularNode struct {
	No   int
	Next *CircularNode
}

func (list *CircularLinkedList) Init(nums int) {
	if nums < 1 {
		return
	}

	var temp *CircularNode
	for i := 1; i <= nums; i++ {
		node := &CircularNode{No: i}
		if i == 1 {
			list.first = node
			list.first.Next = list.first
			temp = list.first
		} else {
			temp.Next = node
			node.Next = list.first
			temp = temp.Next
		}
	}
}

func (list *CircularLinkedList) Pop(listLen, startNo, interval int) {
	if startNo < 1 || startNo > listLen {
		fmt.Println("参数不合法")
		return
	}
	temp := list.first

	// 找开始节点的前一个节点
	for {
		if temp.Next.No == startNo {
			break
		}
		temp = temp.Next
	}

	count := 0

	for {
		if count+1 == interval {
			fmt.Println("出圈 No#", temp.Next.No)
			if temp == temp.Next { // 只剩下一个节点
				break
			}
			count = 0
			temp.Next = temp.Next.Next
			continue
		}
		count++
		temp = temp.Next
	}
}

func (list *CircularLinkedList) Pop2(listLen, startNo, interval int) {
	if startNo < 1 || startNo > listLen {
		fmt.Println("参数不合法")
		return
	}
	temp := list.first

	// 找开始节点的前一个节点
	for {
		if temp.Next.No == startNo {
			break
		}
		temp = temp.Next
	}

	countPointer := temp.Next // 指向开始节点
	for  {
		if temp == countPointer {
			fmt.Println("出圈 No#", countPointer.No)
			break
		}
		for i := 0; i < interval-1; i++ {
			countPointer = countPointer.Next
			temp = temp.Next
		}

		fmt.Println("出圈 No#", countPointer.No)
		countPointer = countPointer.Next
		temp.Next = countPointer
	}
}

func (list *CircularLinkedList) Print() {
	temp := list.first
	for temp != nil {
		fmt.Println("No#", temp.No)
		if temp.Next == list.first {
			return
		}
		temp = temp.Next
	}
}
