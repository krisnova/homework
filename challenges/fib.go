package main

import "fmt"

// The Fibonacci sequence begins with  and  as its first and second terms.
// After these first two elements, each subsequent element is equal to the
// sum of the previous two elements.
//
// (0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 141, ...)

func main() {
	fmt.Println(fibI(0))
	fmt.Println(fibI(1))
	fmt.Println(fibI(2))
	fmt.Println(fibI(3))
	fmt.Println(fibI(4))
	fmt.Println(fibI(5))
	fmt.Println(fibI(6))
	fmt.Println("--")
	fmt.Println(fibR(0))
	fmt.Println(fibR(1))
	fmt.Println(fibR(2))
	fmt.Println(fibR(3))
	fmt.Println(fibR(4))
	fmt.Println(fibR(5))
	fmt.Println(fibR(6))

}

// This is the classic recursive example
// Basically you are just doing additon and changing the
// indexes recursively until everything is less than or equal to 1
//
// O(2^n) (Exponential and SLOW)
func fibR(n int) int {
	if n <= 1 {
		return n
	}
	return fibR(n-1) + fibR(n-2)
}

// This is the much more efficient iterative solution
// O(n) (Fast)
func fibI(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	// Start at 2 because we know 0,1 have been handled
	var f = make(map[int]int)
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}
