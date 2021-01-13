package main

import (
	"fmt"
	"math"
)

// Given an integer N and an array of N integers. Calculate the sum of products of all pairs of integers of the array.
// Note: Since the answer can be large, return the answer modulo 109+7.
//
//Input:
// N=3
// A=[1,2,3]
//Output:
// 11
//Explanation:
// 1x2+2x3+1x3=11
//So, the answer is 11.
func main() {
	tests := []struct {
		n int
		a []int
	}{
		{3, []int{1, 2, 3}},
		{3, []int{2, 2, 3}},
	}
	for _, test := range tests {
		fmt.Println(sum_pairs(test.n, test.a))
	}
}

func sum_pairs(n int, a []int) int {
	n = n - 1
	total := 0
	// Decrement n and assert against every value except self
	for n > 0 {
		for i := n; i >= 0; i-- {
			if i == n {
				continue
			}
			x := a[n] * a[i]
			total = total + x
		}
		n--
	}
	// Default case
	return total % int(math.Pow(10, 7)+7)
}
