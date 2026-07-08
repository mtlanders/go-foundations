package main

import (
	"fmt"
	"runtime"
	"time"
)

//*********************************************************

// Requirement 1
func leakyWorkers(n int, jobs <-chan int) <-chan int {
	retChan := make(chan int)
	for range n {
		go func() {
			for r := range jobs {
				retChan <- (r * 2)
			}
		}()
	}
	return retChan
}

//*********************************************************

func main() {

	jchan := make(chan int)

	fmt.Printf("NumGoroutine (leaky, before): %d\n", runtime.NumGoroutine())
	results := leakyWorkers(3, jchan)

	jchan <- 5  // Send a value for a leaky worker
	jchan <- 10 // Send another value for a leaky worker

	<-results // Read but don't store, we don't actually care about the values

	// No more reads (intentional)

	time.Sleep(1 * time.Second)
	fmt.Printf("NumGoroutine (leaky, after): %d\n", runtime.NumGoroutine())
}

//*********************************************************
