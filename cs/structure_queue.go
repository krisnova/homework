package main

import (
	"fmt"
	"sync"
)

// A queue is a data structure that enforces FIFO
// Items, Elements, Entities

func main() {
	queue := NewQueue()
	e1 := &Element{
		value: "hi",
	}
	e2 := &Element{
		value: "nova",
	}
	queue.Append(e1)
	queue.Append(e2)
	fmt.Println(queue.Length())
	next := queue.Next()
	fmt.Println(next)
	fmt.Println(queue.Length())
}

// NewQueue will create a new FIFO queue
func NewQueue() *Queue {
	return &Queue{
		mutex: sync.Mutex{},
	}
}

// Queue represents the queue and holds the elements
type Queue struct {
	mutex    sync.Mutex
	elements []*Element
}

// Append will add an item to the queue
func (q *Queue) Append(element *Element) error {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.elements = append(q.elements, element)
	return nil
}

// Next will pull the next element off the end of the queue
func (q *Queue) Next() *Element {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	var element *Element
	if len(q.elements) == 0 {
		return nil
	}
	// TODO Nova commit this to memory
	element, q.elements = q.elements[len(q.elements)-1], q.elements[:len(q.elements)-1]
	return element
}

// Length will calculate the length of the queue
func (q *Queue) Length() int {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	return len(q.elements)
}

// Element is an item in the queue
type Element struct {
	value string
}
