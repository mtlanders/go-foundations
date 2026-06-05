//Exercise 4:
//Write a function that takes a pointer to a Pair struct and swaps its two integer fields in place using auto-deref, with no return value.

package main

import "fmt"

type Pair struct {
	X int
	Y int
}

// Function with arg *Pair - Requirement
// No return value - Criteria 2
func swapStructValues(ptr *Pair) {

	// Swap values - Requirement
	temp := ptr.X // Auto-deref to swap - Criteria 1
	ptr.X = ptr.Y // Auto-deref to swap - Criteria 1
	ptr.Y = temp  // Auto-deref to swap - Criteria 1
}

func main() {
	pair := &Pair{5, 10}

	fmt.Printf("Before - X: %d, Y: %d\n", pair.X, pair.Y) // Print before - Criteria 3
	swapStructValues(pair)
	fmt.Printf("After  - X: %d, Y: %d\n", pair.X, pair.Y) // Print after  - Criteria 3
}
