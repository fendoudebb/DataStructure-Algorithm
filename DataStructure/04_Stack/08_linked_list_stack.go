package main

import "fmt"

func main() {
	linkedListStack := &LinkedListStack{}
	linkedListStack.Put("aaa")
	linkedListStack.Put("bbb")
	linkedListStack.Put("ccc")
	linkedListStack.Pop()
	linkedListStack.Pop()
	linkedListStack.Pop()
	linkedListStack.Put("ddd")
	linkedListStack.Put("eee")
	linkedListStack.Put("fff")
	linkedListStack.Pop()
	linkedListStack.Print()
}

type LinkedListStack struct {
	PutPointer *Node // 链表最后一个节点
	HeadNode   *Node
}

type Node struct {
	Value string
	Next  *Node
}

func (stack *LinkedListStack) Put(value string) {
	node := &Node{Value: value}

	// 第一个节点
	if stack.HeadNode == nil {
		stack.HeadNode = node
		stack.PutPointer = stack.HeadNode
		return
	}

	// 之后的节点
	stack.PutPointer.Next = node

	// Put指针往后移
	stack.PutPointer = stack.PutPointer.Next
}

func (stack *LinkedListStack) Pop()  {
	if stack.HeadNode.Next == nil {
		fmt.Println("Pop#", stack.HeadNode.Value)
		stack.HeadNode = nil
		stack.PutPointer = nil
	} else {
		temp := stack.HeadNode
		for temp.Next != nil{
			if temp.Next.Next == nil {
				break
			}
			temp = temp.Next
		}
		fmt.Println("Pop#", temp.Next.Value)
		temp.Next = nil
		stack.PutPointer = temp
	}

}

func (stack *LinkedListStack) Print() {
	temp := stack.HeadNode
	for temp != nil {
		fmt.Println("value#", temp.Value)
		temp = temp.Next
	}
}
