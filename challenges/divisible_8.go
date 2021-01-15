package main

import (
	"fmt"
	"strconv"
)

// Given a number, you need to check whether any permutation of the number is divisible by 8 or not. Print Yes or No.

//Input:
// The first line of input contains a single integer T denoting the number of test cases.
// Then T test cases follow. Each test case consist of one line.
// The first line of each test case consists of an integer N.

// Corresponding to each test case, in a new line, print Yes if divisible by 8, else No.

func main() {

	fmt.Println(perm_div_8(46))
}

func permutations(n int) [][]int {
	str := strconv.Itoa(n)
	var arr []int
	for _, r := range str {
		s := string(r)
		i, _ := strconv.Atoi(s)
		arr = append(arr, i)
	}
	fmt.Println(arr)
	var result [][]int
	recperm := func(arr []int, n int) {
		if len(arr) == 1 {
			result = append(result, arr)
		}
	}

	//fmt.Println(len(str))
	return result
}

func perm_div_8(n int) string {
	var tests []int
	// Calculate permutations and append to tests
	tests = permutations(n)
	for _, test := range tests {
		if (test % 8) == 0 {
			return "yes"
		}
	}
	return "no"
}
