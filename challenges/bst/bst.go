package main

import (
	"fmt"

	"github.com/kris-nova/job2021/cs/bst"
)

func main() {
	btree := bst.NewPopulatedBTree()
	btree.Insert(&bst.Node{
		Key:  "seventy seven",
		Data: 77,
	})
	btree.Insert(&bst.Node{
		Key:  "seventy eight",
		Data: 78,
	})
	btree.Insert(&bst.Node{
		Key:  "seventeen",
		Data: 17,
	})
	btree.Insert(&bst.Node{
		Key:  "fourty two",
		Data: 42,
	})
	btree.Insert(&bst.Node{
		Key:  "eighty four",
		Data: 84,
	})
	btree.Insert(&bst.Node{
		Key:  "sixety nine",
		Data: 69,
	})
	btree.Insert(&bst.Node{
		Key:  "four twenty",
		Data: 420,
	})
	btree.Insert(&bst.Node{
		Key:  "eighty eight",
		Data: 88,
	})
	btree.Insert(&bst.Node{
		Key:  "three seventeen",
		Data: 317,
	})

	// -----

	node := btree.Search("four twenty")
	fmt.Println(*node)
}
