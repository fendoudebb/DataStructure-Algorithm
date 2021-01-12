package main

import "fmt"

func main() {
	singlyLinkedList := &SinglyLinkedList{}
	singlyLinkedList.Insert(1, "111")
	//singlyLinkedList.Append(1, "sj")
	singlyLinkedList.Insert(3, "333")
	singlyLinkedList.Insert(2, "222")
	singlyLinkedList.Insert(5, "555")
	singlyLinkedList.Insert(4, "444")
	singlyLinkedList.Print()
	fmt.Println("链表长度#", singlyLinkedList.Length())
	fmt.Println("-------------")
	singlyLinkedList.ReversePrint()
	fmt.Println("-------------")

	fmt.Printf("LastIndex#%+v\n", singlyLinkedList.LastIndex(1))

	singlyLinkedList.Update(1, "111111111111")
	singlyLinkedList.Update(3, "333333333333")

	singlyLinkedList.Delete(5)
	singlyLinkedList.Delete(4)
	singlyLinkedList.Delete(1)

	singlyLinkedList.Print()
	fmt.Println("链表长度#", singlyLinkedList.Length())

	fmt.Println("-------------")

	singlyLinkedList.Add(6, "666")

	singlyLinkedList.Print()
	fmt.Println("链表长度#", singlyLinkedList.Length())

	fmt.Println("-------------")
	singlyLinkedList.Reverse().Print()


	fmt.Println("-------------")


	first := &SinglyLinkedList{}
	//first.Insert(1, "111")
	first.Insert(3, "333")
	first.Insert(5, "555")
	first.Insert(7, "777")
	first.Insert(9, "999")

	second := &SinglyLinkedList{}
	second.Insert(1, "111")
	second.Insert(2, "222")
	second.Insert(4, "444")
	second.Insert(6, "666")
	second.Insert(8, "888")

	MergeTwoOrderLinkedList(first, second).Print()


}

type Node struct {
	No   int
	Name string
	Next *Node
}

type SinglyLinkedList struct {
	HeadNode *Node
}

// Append 从尾部添加
func (list *SinglyLinkedList) Append(no int, name string) {
	node := &Node{No: no, Name: name}
	temp := list.HeadNode
	if temp == nil {
		list.HeadNode = node
		return
	}
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = node
}

// Add 从头部添加
func (list *SinglyLinkedList) Add(no int, name string) {
	node := &Node{No: no, Name: name}
	if list.HeadNode == nil {
		list.HeadNode = node
	} else {
		node.Next = list.HeadNode
		list.HeadNode = node
	}
}

func (list *SinglyLinkedList) Insert(no int, name string) {
	node := &Node{No: no, Name: name}
	temp := list.HeadNode
	if temp == nil {
		list.HeadNode = node
		return
	}
	if temp.No > no {
		node.Next = temp
		list.HeadNode = node
	} else {
		for temp.Next != nil {
			if temp.Next.No > no {
				break
			}
			temp = temp.Next
		}
		node.Next = temp.Next
		temp.Next = node

	}

}

func (list *SinglyLinkedList) Update(no int, name string) {
	if list.HeadNode == nil {
		return
	}

	temp := list.HeadNode
	for temp.No != no {
		temp = temp.Next
	}

	if temp != nil {
		temp.Name = name
	}

}

func (list *SinglyLinkedList) Delete(no int) {
	if list.HeadNode == nil {
		return
	}
	temp := list.HeadNode
	if temp.No == no {
		list.HeadNode = list.HeadNode.Next
		return
	}

	for temp.Next != nil {
		if temp.Next.No == no {
			break
		}
		temp = temp.Next
	}

	temp.Next = temp.Next.Next
}

// LastIndex 查找倒数第 N 个节点
func (list *SinglyLinkedList) LastIndex(index int) *Node {
	if index <= 0 || list.Length() < index {
		return nil
	}

	temp := list.HeadNode
	for i := 0; i < list.Length()-index; i++ {
		temp = temp.Next
	}
	return temp
}

func (list *SinglyLinkedList) Reverse() *SinglyLinkedList {
	temp := list.HeadNode

	if temp == nil || temp.Next == nil {
		return list
	}

	reverseList := new(SinglyLinkedList)

	var next *Node

	for temp != nil {
		next = temp.Next
		temp.Next = reverseList.HeadNode
		reverseList.HeadNode = temp
		temp = next
	}
	return reverseList
}

func (list *SinglyLinkedList) Length() int {
	temp := list.HeadNode
	count := 0
	for temp != nil {
		count++
		temp = temp.Next
	}
	return count
}

func (list *SinglyLinkedList) Print() {
	temp := list.HeadNode
	for temp != nil {
		fmt.Println("No#", temp.No, "Name#", temp.Name)
		temp = temp.Next
	}
}

func (list *SinglyLinkedList) ReversePrint() {
	temp := list.HeadNode
	for temp != nil {
		defer func(no int, name string) {
			fmt.Println("No#", no, "Name#", name)
		}(temp.No, temp.Name)
		temp = temp.Next
	}
}

func MergeTwoOrderLinkedList(first, second *SinglyLinkedList) *SinglyLinkedList {
	if first == nil {
		return second
	}
	if second == nil {
		return first
	}

	temp2 := second.HeadNode

	var next *Node

	for temp2 != nil {
		next = temp2.Next

		temp1 := first.HeadNode
		if temp1.No < temp2.No {
			for temp1.Next != nil {
				if temp1.Next.No > temp2.No {
					break
				}
				temp1 = temp1.Next
			}
			temp2.Next = temp1.Next
			temp1.Next = temp2
		} else {
			temp2.Next = temp1
			first.HeadNode = temp2
		}
		temp2 = next
	}

	return first

}
