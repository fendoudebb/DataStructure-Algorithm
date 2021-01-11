package main

import "fmt"

func main() {
	queue := NewArrayQueue(3)
	queue.Add(1)
	queue.Add(2)
	queue.Add(3)
	queue.Get()
	queue.Get()
	queue.ShowQueue()
	fmt.Println("queue size#", queue.Size())

}

type ArrayQueue struct {
	head int
	tail int
	arr  []int
}

func NewArrayQueue(size int) *ArrayQueue {
	instance := new(ArrayQueue)
	instance.head = 0
	instance.tail = 0
	instance.arr = make([]int, size)
	return instance
}

func (receiver *ArrayQueue) IsFull() bool {
	return receiver.tail == len(receiver.arr)
}

func (receiver *ArrayQueue) IsEmpty() bool {
	return receiver.head == receiver.tail
}

func (receiver *ArrayQueue) Add(element int) {
	if receiver.IsFull() {
		panic("队列已满，无法加入")
	}
	receiver.arr[receiver.tail] = element
	receiver.tail++
}

func (receiver *ArrayQueue) Get() int {
	if receiver.IsEmpty() {
		panic("队列为空，无法获取")
	}
	defer func() {
		receiver.head++
	}()
	return receiver.arr[receiver.head]
}

func (receiver *ArrayQueue) Size() int {
	return receiver.tail - receiver.head
}

func (receiver *ArrayQueue) ShowQueue() {
	if receiver.IsEmpty() {
		return
	}
	for i := receiver.head; i < receiver.tail; i++ {
		fmt.Println(receiver.arr[i])
	}
}
