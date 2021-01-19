package bst

// NewPopulatedBTree Will return a populated a btree
func NewPopulatedBTree() *BTree {
	btree := NewBTree("Nóva")
	btree.Insert(&Node{
		Key:  "seventy seven",
		Data: 77,
	})
	btree.Insert(&Node{
		Key:  "seventy eight",
		Data: 78,
	})
	btree.Insert(&Node{
		Key:  "seventeen",
		Data: 17,
	})
	btree.Insert(&Node{
		Key:  "fourty two",
		Data: 42,
	})
	btree.Insert(&Node{
		Key:  "eighty four",
		Data: 84,
	})
	btree.Insert(&Node{
		Key:  "sixety nine",
		Data: 69,
	})
	btree.Insert(&Node{
		Key:  "four twenty",
		Data: 420,
	})
	btree.Insert(&Node{
		Key:  "eighty eight",
		Data: 88,
	})
	btree.Insert(&Node{
		Key:  "three seventeen",
		Data: 317,
	})
	return btree
}

// NewPopulatedBTree Will return a populated a btree
// This implementation used to work, but is now broken due to how we have 3 fields
// Key, Value, and Data
// This is in fact a BROKEN tree
func NewInvalidPopulatedBTree() *BTree {
	btree := NewBTree("Nóva")
	btree.Insert(&Node{
		Key:  "one hundred",
		Data: 100,
		Left: &Node{
			Key:  "seventy five",
			Data: 75,
			Left: &Node{
				Key:  "fifty",
				Data: 50,
				Left: &Node{
					Key:  "twenty five",
					Data: 25,
					Left: &Node{
						Key:  "fifteen",
						Data: 15,
						Left: &Node{
							Key:  "seven",
							Data: 7,
						},
						Right: &Node{
							Key:  "twenty",
							Data: 20,
						},
					},
				},
			},
			Right: &Node{
				Key:  "eighty five",
				Data: 85,
				Left: &Node{
					Key:  "eighty",
					Data: 80,
				},
				Right: &Node{
					Key:  "ninety",
					Data: 90,
				},
			},
		},
		Right: &Node{
			Key:  "one hundred twenty five",
			Data: 125,
			Left: &Node{
				Key:  "one hundred fifteen",
				Data: 115,
			},
			Right: &Node{
				Key:  "one hundred fifty",
				Data: 150,
			},
		},
	})
	return btree
}
