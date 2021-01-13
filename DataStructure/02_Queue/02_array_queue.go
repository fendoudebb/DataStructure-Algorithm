package main

import "fmt"

func main() {
	queue := &ArrayQueue{arr: make([]int, 3)}
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

func (queue *ArrayQueue) IsFull() bool {
	return queue.tail == len(queue.arr)
}

func (queue *ArrayQueue) IsEmpty() bool {
	return queue.head == queue.tail
}

func (queue *ArrayQueue) Add(element int) {
	if queue.IsFull() {
		panic("队列已满，无法加入")
	}
	queue.arr[queue.tail] = element
	queue.tail++
}

func (queue *ArrayQueue) Get() int {
	if queue.IsEmpty() {
		panic("队列为空，无法获取")
	}
	defer func() {
		queue.head++
	}()
	return queue.arr[queue.head]
}

func (queue *ArrayQueue) Size() int {
	return queue.tail - queue.head
}

func (queue *ArrayQueue) ShowQueue() {
	if queue.IsEmpty() {
		return
	}
	for i := queue.head; i < queue.tail; i++ {
		fmt.Println(queue.arr[i])
	}
}
