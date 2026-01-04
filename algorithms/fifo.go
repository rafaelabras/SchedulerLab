package algorithms

import (
	_ "github.com/rafaelabras/SchedulerLab/process"
)

type Queue struct {
	items []int
}

func NewQueue() *Queue {
	queue := new(Queue)
	return queue
}

func (queue *Queue) Enqueue(value int) {
	queue.items = append(queue.items, value)
}

func (queue *Queue) Pop() int {
	if len(queue.items) == 0 {
		panic("Queue is empty")
	}

	item := queue.items[0]
	queue.items = queue.items[1:]

	return item
}

func (queue *Queue) Len() int {
	return len(queue.items)
}
