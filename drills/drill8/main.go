package main

import "fmt"

//*********************************************************

// Requirement 1
func makeLimiter(max int) func() bool {
	// Making unsigned to prevent falling below 0
	var counter uint8 = uint8(max)
	return func() bool {
		if counter > 0 {
			counter--
			return true
		}
		return false
	}
}

//*********************************************************

func main() {

	// Requirement 2
	lim1 := makeLimiter(3)
	lim2 := makeLimiter(5)

	// Requirement 3
	// Limiter 1 count: 2, expected: true
	// Can't tell if independent so far
	valid1 := lim1()

	// Limiter 2 count: 4, expected: true
	// Can't tell if independent so far
	valid2 := lim2()

	// Limiter 1 count: 1, expected: true
	// Can't tell if independent so far
	valid1 = lim1()

	// Limiter 2 count: 3, expected: true
	// Both passed, so let's continue checking
	valid2 = lim2()

	// If the counter was shared, they would both fail here (lim1 has 3 tokens)
	fmt.Printf("Checking for independence for lim1: %t (expected true)\n", valid1)
	fmt.Printf("Checking for independence for lim2: %t (expected true)\n", valid2)

	// Limiter 1 count: 0, expected: true
	// Still can't be certain, but relatively confident
	valid1 = lim1()

	// This is where we think the divergence might really show
	fmt.Printf("lim1 at    3 (max)     calls: %t  (expected true)\n", valid1) // Actual: true

	// Limiter 1 count: 0, expected: false
	valid1 = lim1()

	// Lim1 fails here, but lim2 is still valid. So lim1 and lim2 are truly independent
	fmt.Printf("lim1 after 4 (max + 1) calls: %t (expected false)\n", valid1) // Actual: false
	fmt.Printf("lim2 after 2           calls: %t  (expected true)\n", valid2) // Actual: true

	// Limiter 2 count: 2, expected: true
	// Lim2 is truly independent, so let's now just exhaust
	valid2 = lim2()

	// Limiter 2 count: 1, expected: true
	valid2 = lim2()

	// Limiter 2 count: 0, expected: true
	valid2 = lim2()

	fmt.Printf("lim2 at    5 (max)     calls: %t  (expected true)\n", valid2) // Actual: true

	// Limiter 2 count: 0, expected: false
	valid2 = lim2()

	// Limiter 2 count: 0, expected: false (exhausted)
	fmt.Printf("lim2 after 6 (max + 1) calls: %t (expected false)\n", valid2) // Actual: false
}
