package main

import "fmt"

// Given an array of integers (both positive and negative). Find the
// contiguous sequence with the largest sum. Return the sum.
//
// Input: 2, -8, 3, -2, 4, -10
// Output: 5 (3, -2, 4)
//
// Sequence On ...Xn
func main() {

	type test struct {
		input    []int
		solution int
	}

	tests := []test{
		{
			input:    []int{2, -8, 3, -2, 4, -10},
			solution: 5, // 5  (3, -2, 4)
		},
		{
			input:    []int{1, 2, 3},
			solution: 6, // 6  (1, 2, 3)
		},
		{
			input:    []int{0, -10, 1, 2},
			solution: 3, // (1, 2)
		},
		{
			input:    []int{0, -10, 4, 4, 0, 1, -1, 2, -9},
			solution: 10, // 10 (4, 4, 0, 1, -1, 2)
		},
	}

	for _, test := range tests {
		x := contiguousSeqSum(test.input)
		if test.solution == x {
			// Win
			fmt.Printf("Success [%d]\n", x)
		} else {
			// Lose
			fmt.Printf("Failure have(%d) want(%d)\n", x, test.solution)
		}
	}

}

// Worst case computational overhead θ(n)
// where n is the size of the input parameters (arr int)
func contiguousSeqSum(arr []int) int {
	l := len(arr) // 1 indexed
	x := 0
	max := 0
	for i := 0; i < l; i++ {
		x = x + arr[i]
		if x > max {
			max = x
		} else if x < 0 {
			x = 0
		}
	}
	return max
}
