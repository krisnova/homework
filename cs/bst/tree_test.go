package bst

import "testing"

// TestBSTHappy will check to make sure
// we have a happy binary search tree
func TestBSTHappy(t *testing.T) {
	tree := NewPopulatedBTree()
	if !tree.IsBST() {
		t.Errorf("Failure: isnotbst")
	}
}

func TestBSTInvalidHappy(t *testing.T) {
	tree := NewInvalidPopulatedBTree()
	if tree.IsBST() {
		t.Errorf("Failure: isbst")
	}
}

func TestBSTInsertHappy(t *testing.T) {
	btree := NewPopulatedBTree()
	btree.Insert(&Node{
		Key:   "seventy seven",
		Value: 77,
	})
	btree.Insert(&Node{
		Key:   "seventy eight",
		Value: 78,
	})
	btree.Insert(&Node{
		Key:   "seventeen",
		Value: 17,
	})
	btree.Insert(&Node{
		Key:   "fourty two",
		Value: 42,
	})
	btree.Insert(&Node{
		Key:   "eighty four",
		Value: 84,
	})
	btree.Insert(&Node{
		Key:   "sixety nine",
		Value: 69,
	})
	btree.Insert(&Node{
		Key:   "four twenty",
		Value: 420,
	})
	btree.Insert(&Node{
		Key:   "eighty eight",
		Value: 88,
	})
	btree.Insert(&Node{
		Key:   "three seventeen",
		Value: 317,
	})
	if !btree.IsBST() {
		t.Errorf("Insert created invalid binary search tree")
	}
}
