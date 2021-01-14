package main

import "fmt"

func main() {
	doublyLinkedList := &DoublyLinkedList{}
	doublyLinkedList.Insert2(3, "333")
	doublyLinkedList.Insert2(1, "111")
	doublyLinkedList.Insert2(2, "222")
	doublyLinkedList.Delete(2)
	doublyLinkedList.Add(4, "444")
	doublyLinkedList.Add(5, "555")

	doublyLinkedList.Print()
}

type DoublyLinkedList struct {
	HeadNode *DoublyNode
}

type DoublyNode struct {
	No   int
	Name string
	Pre  *DoublyNode
	Next *DoublyNode
}

func (list *DoublyLinkedList) Add(no int, name string) {
	node := &DoublyNode{
		No:   no,
		Name: name,
	}
	temp := list.HeadNode
	if temp == nil {
		list.HeadNode = node
		return
	}

	for temp.Next != nil {
		temp = temp.Next
	}

	temp.Next = node
	node.Pre = temp
}

func (list *DoublyLinkedList) Insert(no int, name string) {
	node := &DoublyNode{
		No:   no,
		Name: name,
	}

	if list.HeadNode == nil {
		list.HeadNode = node
		return
	}

	temp := list.HeadNode

	for temp != nil {
		if temp.No > node.No {
			if temp.Pre == nil {
				// 此时temp为头节点
				node.Next = temp
				temp.Pre = node
				list.HeadNode = node
			} else {
				node.Next = temp
				node.Pre = temp.Pre
				temp.Pre.Next = node
				temp.Pre = node
			}
			break
		} else {
			if temp.Next == nil {
				// 到了链表末尾，此时插入的节点 No 数值最大
				node.Pre = temp
				temp.Next = node
				// 末尾新增了一个节点，必须break，否则将无限循环
				break
			}
		}
		temp = temp.Next
	}
}

func (list *DoublyLinkedList) Insert2(no int, name string) {
	node := &DoublyNode{
		No:   no,
		Name: name,
	}

	if list.HeadNode == nil {
		list.HeadNode = node
		return
	}

	temp := list.HeadNode
	if temp.No > node.No {
		node.Next = temp
		temp.Pre = node
		list.HeadNode = node
	} else {
		for temp.Next != nil {
			if temp.Next.No > no {
				break
			}
			temp = temp.Next
		}

		node.Next = temp.Next
		node.Pre = temp
		if temp.Next != nil {
			temp.Next.Pre = node
		}
		temp.Next = node
	}
}

func (list *DoublyLinkedList) Delete(no int) {
	temp := list.HeadNode
	for temp != nil {
		if temp.No == no {
			temp.Pre.Next = temp.Next
			temp.Next.Pre = temp.Pre
			temp.Next = nil
			temp.Pre = nil
			break
		}
		temp = temp.Next
	}
}

func (list *DoublyLinkedList) Print() {
	temp := list.HeadNode
	for temp != nil {
		fmt.Println("No#", temp.No, "Name#", temp.Name)
		temp = temp.Next
	}
}
