package main

import "fmt"

func main() {
	arrayStack := &ArrayStack{-1, make([]string, 10)}

	arrayStack.Put("aaa")
	arrayStack.Put("bbb")
	arrayStack.Put("ccc")

	arrayStack.Pop()
	arrayStack.Pop()
	arrayStack.Pop()
}

type ArrayStack struct {
	Top   int
	Array []string
}

func (stack *ArrayStack) Put(value string) {
	if stack.Top < len(stack.Array)-1 {
		stack.Top++
		stack.Array[stack.Top] = value
	} else {
		fmt.Println("stack is full")
	}
}

func (stack *ArrayStack) Pop() {
	if stack.Top < 0 {
		fmt.Println("stack is empty")
	} else {
		fmt.Println("pop#", stack.Array[stack.Top])
		stack.Top--
	}
}
