package main

import "fmt"

//*******************************************************************

/*
	Justification:

	This justification was made because startingPrice will be captured
	as a reference by the closure. This is the price we're checking against
	to begin with. This is the price provided when the watcher is created.

	Result grabs the boolean result of comparing the new price to startingPrice
	in the new bound return function. I used the example from the module as a
	reference but used the parameter variable instead of creating a new one.

	If result is "true" then update startingPrice to n

	Regardless, return boolean "result"
*/
func makeHighWatcher(startingPrice int) func(int) bool {
	return func(n int) bool {
		result := n > startingPrice // New price greater than starting price (ties don't trigger the if block)
		if result {
			startingPrice = n // Set starting price to new price
		}
		return result
	}
}

//*******************************************************************

func main() {

	// Creating unique watchers and setting startingPrice to 10
	f1 := makeHighWatcher(10)
	f2 := makeHighWatcher(10)

	res := f1(9)                           // Expected: false
	fmt.Printf("Actual f1(9): %v\n", res)  // Actual: false
	res = f1(10)                           // Expected: false (tie-breaker)
	fmt.Printf("Actual f1(10): %v\n", res) // Actual: false
	res = f1(15)                           // Expected: true
	fmt.Printf("Actual f1(15): %v\n", res) // Actual: true

	// Running new watcher to prove the two are divergent
	res = f2(6)                            // Expected: false
	fmt.Printf("Actual f1(6): %v\n", res)  // Actual: false
	res = f2(11)                           // Expected: true, would fail on f1, since high is now 15
	fmt.Printf("Actual f1(11): %v\n", res) // Actual: true
	res = f2(12)                           // Expected: true, would fail on f1, since high is now 15
	fmt.Printf("Actual f1(12): %v\n", res) // Actual: true
	res = f2(13)                           // Expected: true, would fail on f1, since high is now 15
	fmt.Printf("Actual f1(13): %v\n", res) // Actual: true

	/*
		Justification:

		The spec can be interpreted as distinct from use makeHighWatcher.
		I reimplemented the logic as part of the loop closure body. My approach
		was identical, but uses a temporary copy into x to prevent the shared
		i bug. I then tested with a data set explicitly designed to make the
		shared bug surface: 5, 4, 3

		- Shared i initially: 3 (at loop exit)
		- f[0](5): 5 > 3 (true, but propagates to all)
		- f[1](4): 4 > 5 (false)
		- f[2](3): 3 > 5 (false)

		With the temporary x in each closure, the closure is now bound
		to the value in X rather than i, which will negate the shared i
		bug.

	*/
	fs := []func(int) bool{}
	for i := 0; i < 3; i++ {
		x := i
		fs = append(fs, func(n int) bool {
			result := n > x
			if result {
				x = n
			}
			return result
		})
	}

	// Should be true if i was independent
	res = fs[0](5) // Actual: true
	fmt.Printf("Actual fs[0](5): %v\n", res)

	// Should be true if i was independent
	res = fs[1](4) // Actual: true
	fmt.Printf("Actual fs[1](4): %v\n", res)

	// Should be true if i was independent
	res = fs[2](3) // Actual: true
	fmt.Printf("Actual fs[2](3): %v\n", res)
}
