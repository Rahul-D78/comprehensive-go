package dsa

import "fmt"

// Queue represents a queue that holds a slice
type queue struct {
	items []int
}

// Enqueue will add a value at the end
func (q *queue) enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue will remove a value from the front
func (q *queue) dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func Queue() {
	myQueue := queue{}
	fmt.Println(myQueue)
	myQueue.enqueue(100)
	myQueue.enqueue(200)
	myQueue.enqueue(300)
	myQueue.enqueue(400)
	fmt.Println(myQueue)
	myQueue.dequeue()
	fmt.Println(myQueue)
}
