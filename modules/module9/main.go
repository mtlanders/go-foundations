package main

import "fmt"

// Requirement 2 - Worker goroutine
func worker(j chan int, r chan int) {
	val := 0
	for ii := 0; ii < 5; ii++ {
		val = <-j
		r <- (val * 2)
	}
}

func main() {

	// Requirement 1
	jobs := make(chan int)
	results := make(chan int, 3)

	// Requirement 2
	go worker(jobs, results)

	val := -1
	for ii := 0; ii < 5; ii++ {

		// Requirement 3
		jobs <- (ii + 1)

		// Requirement 4
		val = <-results
		fmt.Printf("Interation %d received '%d' from results\n", ii, val)

		// Dev note: literal spec reading implies "send all first, then receive all"
		// This is not feasible, so I modified the code to interweave the sends/receives
	}

	/*
	   Yes, the program would still run if results were unbuffered.
	   The send/receive order has every send proceeded by its
	   corresponding receive/send. Any blocked send is immediately
	   followed by the receive which wakes it up.

	   1. Main: jobs sends
	   2. worker: j (jobs) receives
	   2. worker: r (results) sends
	   3. Main: results receives

	   At any point if a send was blocked it is immediately unblocked
	   by the receive operation on that same channel. So yes, buffered
	   or unbuffered would work identically here.
	*/
}
