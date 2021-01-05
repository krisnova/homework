package main

import "fmt"

// This program can be used to create a LinkedList of arbitrary length
// as well as insert a node anywhere in the list.
// TODO we still need to check appending to the end of the list.

func main() {
	list := NewEmptyList(10)
	fmt.Println("Len: ", list.len)
	newNode := &Node{}
	InsertAt(list, newNode, 0)
	fmt.Println("Len: ", list.len)
}

// NewEmptyList creates a new empty list
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

// InstertAt is used to insert a node in the LinkedList
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

// Node represents a Node in the list
// Linked Lists do not have array like IDs you can look up (but do have names, content, etc)
// Linked Lists only have next* and not previous*
type Node struct {
	next *Node
}

// LinkedList is basically just a Node but with a len and is the
// structure you pass around in the program.
type LinkedList struct {
	next *Node
	len  int
}
