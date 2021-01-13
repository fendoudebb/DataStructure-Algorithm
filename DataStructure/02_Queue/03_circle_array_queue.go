package main

import "fmt"

// 空一个位置
func main() {
	queue := &CircleArrayQueue{arr: make([]int, 3)}
	queue.Add(1)
	queue.Add(2)
	queue.Get()
	queue.Get()
	queue.Add(3)
	queue.Add(4)
	queue.ShowQueue()
	fmt.Println("circle queue size#", queue.Size())
}

type CircleArrayQueue struct {
	head int
	tail int
	arr  []int
}

func (queue *CircleArrayQueue) IsFull() bool {
	return (queue.tail+1)%len(queue.arr) == queue.head
}

func (queue *CircleArrayQueue) IsEmpty() bool {
	return queue.head == queue.tail
}

func (queue *CircleArrayQueue) Add(element int) {
	if queue.IsFull() {
		panic("队列已满，无法加入")
	}
	queue.arr[queue.tail] = element
	queue.tail = (queue.tail + 1) % len(queue.arr)
}

func (queue *CircleArrayQueue) Get() int {
	if queue.IsEmpty() {
		panic("队列为空，无法获取")
	}
	defer func() {
		queue.head = (queue.head + 1) % len(queue.arr)
	}()
	return queue.arr[queue.head]
}

func (queue *CircleArrayQueue) Size() int {
	return (queue.tail + len(queue.arr) - queue.head) % len(queue.arr)
}

func (queue *CircleArrayQueue) ShowQueue() {
	if queue.IsEmpty() {
		return
	}
	for i := queue.head; i < queue.head+queue.Size(); i++ {
		fmt.Println(queue.arr[i%len(queue.arr)])
	}
}
