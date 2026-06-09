// COVERS:
// := not used at package level
// Variadic ...int placement
// make for a slice
// copy return value captured and used
// copy partial-copy behavior (detected and handled)
// Error string conventions (lowercase, no punctuation)
// Named return values without shadowing
// Slice expansion at call site (dst...)

// SCENARIO
// A small data pipeline utility processes batches of integer readings from a sensor.
// It needs to normalize the readings into a fixed-size output buffer, compute a running
// sum using a helper, and report any issues clearly.

// REQUIREMENTS

// Declare a package-level constant bufferSize = 8. Do not use := for any package-level declarations.
// Write a variadic function sum(nums ...int) int that returns the sum of all provided integers.
// Write a function normalize(src []int, dst []int) (int, error) using named return values (copied int, err error):

// Use copy to copy from src into dst
// Assign the result of copy to copied
// If copied < len(src), set err to an appropriate error — follow error string conventions
// Inside the function body, do not shadow copied with := after it has been assigned

// In main:

// Initialize dst using make with capacity bufferSize
// Call normalize and handle the error
// Pass dst to sum using the correct expansion syntax

// ACCEPTANCE CRITERIA

// No package-level :=
// Variadic parameter uses ...int not int ...nums
// make is called correctly for a slice
// copy return value is captured and used
// Error string is lowercase with no trailing punctuation
// copied is not shadowed inside normalize
// dst is expanded correctly when passed to sum

package main

import (
	"fmt"
)

// Requirement 1
const bufferSize = 8

// Requirement 2
func sum(nums ...int) int {
	total := 0
	if len(nums) > 0 {
		for _, n := range nums {
			total += n
		}
	}
	return total
}

// Requirement 3
func normalize(dst []int, src []int) (copied int, err error) {
	copied = copy(dst, src) // Requirement 4
	if copied < len(src) {  // Requirement 5
		// Requirement 6
		return copied, fmt.Errorf("could not fully copy slice src: %d values copied", copied)
	}
	return copied, nil
}

func main() {

	dest := make([]int, bufferSize)
	source := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	numCopied, err := normalize(dest, source)
	if err != nil {
		fmt.Println(err)
	}

	sumTotal := sum(dest...)
	fmt.Printf("Sum total of %d copied values from normalized slice(s): %d\n", numCopied, sumTotal)
}
