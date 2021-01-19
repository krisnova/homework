package bst

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"sync"
)

// BTree is a Binary Search Tree
// used for efficient lookups on large
// data sets
type BTree struct {
	Name      string
	Root      *Node
	NodeCount int
	Mutex     sync.Mutex
}

// Node represents a single object in
// the tree
type Node struct {

	// Key is the string of the key
	Key string

	// Value is the calculated integer value of the key
	Value int

	// Data is the content of the Node
	Data interface{}

	// Left represents another (recursive)
	// BST who's values never exceed the value
	// of this Node
	Left *Node // LT < Less Than

	// Right represents another (recursive)
	// BST who's values always exceed the value
	// of this Node
	Right *Node // GT > Greater Than
}

// CalculateValue will populate and return
// the calculated integer value of the key
func (node *Node) CalculateValue() int {
	// Calculate MD5 sum and form an int result
	h := md5.New()
	h.Write([]byte(node.Key))
	md5 := hex.EncodeToString(h.Sum(nil))
	// Now that we have a idempotent hash, we can calculate an integer value
	bigInt := big.NewInt(0)
	bigInt.SetString(md5, 16) // Probably get away with base 2 or base 16
	i64 := bigInt.Int64()
	n := int(i64)
	// n is the int result of our arithmetic
	if n < 0 {
		n = n * -1 // Take the absolute value of n
	}
	node.Value = n
	return n
}

// NewBTree will initialize a new binary search tree
func NewBTree(name string) *BTree {
	btree := &BTree{
		Name: name,
	}
	return btree
}

// Insert is used to insert a new Node
// in the binary search tree
func (b *BTree) Insert(node *Node) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	node.CalculateValue()
	if b.NodeCount < 1 {
		b.NodeCount++
		b.Root = node
		return
	}
	b.NodeCount++
	recursiveInsertNode(b.Root, node)
}

func recursiveInsertNode(root, node *Node) {
	// Replace
	if node.Value <= root.Value {
		if root.Left == nil {
			root.Left = node
		} else {
			recursiveInsertNode(root.Left, node)
		}
		// Right
	} else {
		if root.Right == nil {
			root.Right = node
		} else {
			recursiveInsertNode(root.Right, node)
		}
	}
}

func (b *BTree) Search(key string) *Node {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	search := &Node{
		Key: key,
	}
	search.CalculateValue()
	return recursiveSearch(search, b.Root)
}

func recursiveSearch(search, node *Node) *Node {
	if node == nil {
		return nil
	}
	if search.Value == node.Value && search.Key == node.Key {
		return node
	}
	if search.Value <= node.Value {
		// Left
		return recursiveSearch(search, node.Left)
	}
	if search.Value >= node.Value {
		// Right
		return recursiveSearch(search, node.Right)
	}
	return nil
}

func (b *BTree) String() (string, error) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	jbytes, err := json.Marshal(b)
	if err != nil {
		return "", fmt.Errorf("unable to marshal: %v", err)
	}
	return string(jbytes), nil
}

// IsBST can be used to validate that the BST
// is in fact a binary search tree
func (b *BTree) IsBST() bool {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	//return iterativeCheck(b.Root)
	return recursiveCheck(b.Root)
}

func iterativeCheck(node *Node) bool {
	// Faster
	checking := true
	//processing := node
	for checking {
		// logic to operate on processing
		// left and right are greater than =
		// update processing
		// break
	}

	return true
}

func recursiveCheck(node *Node) bool {
	result := true
	if node.Left != nil {
		// We have a left Node
		// Left is always less than or equal
		if node.Left.Value > node.Value {
			return false
		}
		result = recursiveCheck(node.Left)
	}
	if node.Right != nil {
		// We have a right Node
		// Right is always greater than or equal
		if node.Right.Value < node.Value {
			return false
		}
		result = recursiveCheck(node.Right)
	}
	// Add more rules here
	return result
}
