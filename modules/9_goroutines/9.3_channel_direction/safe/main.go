package main

import (
	"fmt"
	"runtime"
	"time"
)

//*********************************************************

// Requirement 2
func safeWorkers(n int, jobs <-chan int, done <-chan struct{}) <-chan int {
	retChan := make(chan int)
	for range n {
		go func() {
			for {
				var v int

				// First select:
				// Need to gate the receive with a done, since receive
				// can block indefinitely. Don't need to nest the send
				// within the case body because that can also block
				// indefinitely but will be without an exit path
				select {
				case val, ok := <-jobs:
					if !ok {
						return
					}
					v = val * 2
				case _, ok := <-done:
					if !ok {
						return
					}
				}

				// Second select:
				// Gating the send with done to prevent blocking,
				// comes after the receive due to data dependency
				// (can't send modified data if you don't have it)
				select {
				case retChan <- v:
				case <-done:
					return
				}
			}
		}()
	}
	return retChan
}

//*********************************************************

func main() {

	jchan := make(chan int)
	done := make(chan struct{})

	fmt.Printf("NumGoroutine (safe, before): %d\n", runtime.NumGoroutine())
	results := safeWorkers(3, jchan, done)

	jchan <- 5  // Send a value for a safe worker
	jchan <- 10 // Send another value for a safe worker

	<-results // Read but don't store, we don't actually care about the values

	// No more reads (intentional)

	close(done)

	time.Sleep(1 * time.Second)
	fmt.Printf("NumGoroutine (safe, after): %d\n", runtime.NumGoroutine())
}

//*********************************************************
