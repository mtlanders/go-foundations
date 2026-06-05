// Exercise 3:
// Write a function that takes a pointer to a slice of integers and appends a variadic list of
// values to it directly, modifying the caller's slice without returning anything.

package main

import "fmt"

// Function with arg *[]int - Criteria 1
func appendSlice(before *[]int, list ...int) {
	for _, n := range list {
		(*before) = append((*before), n)
	}
}

func main() {
	var slice []int // Nil slice - Criteria 4
	ptr := &slice
	fmt.Printf("before: %d\n", (*ptr)) // Print before - Criteria 5
	appendSlice(ptr, 1, 2, 3, 4)
	fmt.Printf("after : %d\n", (*ptr)) // Print after - Criteria 5
}
