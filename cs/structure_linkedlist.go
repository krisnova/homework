package main

import "fmt"

// Implement a LinkedList of 10 nodes in Go
func main() {
	list := NewEmptyList(10)
	fmt.Println("Len: ", list.len)
	newNode := &Node{}
	InsertAt(list, newNode, 0)
	fmt.Println("Len: ", list.len)
}

// Creates a new empty list
func NewEmptyList(size int) *LinkedList {
	root := &LinkedList{}
	var current *Node
	// 0 - 10 == 9 items + 1 root = 10 total
	j := 0
	for i := 0; i < size; i++ {
		new := &Node{}
		if i == 0 {
			root.next = new
		} else {
			current.next = new
		}
		current = new
		j++
	}
	// The reason I would do this in prod is because of course
	// someone is going to change the code in the loop
	root.len = j
	return root
}

// Q: Since this is a pointer do we have to return anything?
func InsertAt(list *LinkedList, node *Node, pos int) {
	if pos == 0 {
		newNode := &Node{
			next: list.next,
		}
		list.next = newNode
		list.len++
		return
	}
	var current *Node
	for i := 0; i < list.len; i++ {
		if i == 0 {
			current = list.next
		} else if i == pos {
			// Insert element and increase length
			node.next = current.next
			current = node
			list.len++
		} else {
			current = current.next
		}
	}
}

// Linked Lists do not have array like IDs you can look up (but do have names, content, etc)
// Linked Lists only have next* and not previous*
type Node struct {
	next *Node
}

// Basically just a Node but with a len
type LinkedList struct {
	next *Node
	len  int
}
