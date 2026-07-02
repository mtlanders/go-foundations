package main

import "fmt"

//*********************************************************

// Requirement 1
func produce(out chan<- int, done <-chan struct{}) {
	for ii := 1; ii <= 20; ii++ {
		select {
		case out <- ii:
		case <-done:
			close(out) // Requirement 5 - Sender on "jobs" responsible for closing
			return
		}
	}
	close(out) // Requirement 5 - Sender on "jobs" responsible for closing
}

//*********************************************************

// Requirement 2
func work(in <-chan int, out chan<- int, done <-chan struct{}) {
	for {

		select {
		case v, ok := <-in:
			if !ok {
				close(out) // Requirement 5 - Sender on "results" responsible for closing
				return
			}
			if v%2 == 0 { // Requirement 2
				out <- v
			}
		case <-done:
			close(out) // Requirement 5 - Sender on "results" responsible for closing
			return
		}
	}
}

//*********************************************************

func main() {

	jobs := make(chan int)        // Requirement 3a
	results := make(chan int, 10) // Requirement 3a
	done := make(chan struct{})   // Requirement 3b

	go produce(jobs, done)       // Requirement 3c
	go work(jobs, results, done) // Requirement 3c

	for r := range results { // Requirement 3d
		fmt.Printf("Received value from results: %d\n", r)
	}

	close(done) // Requirement 5 - Main responsible for closing "done" broadcast
}

//*********************************************************
