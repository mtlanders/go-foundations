package main

import "fmt"

//*********************************************************

// Requirement 1
func generate(out chan<- string, done <-chan struct{}) {
	strs := []string{"go", "rust", "c", "typescript", "python", "java", "c++"}
	for _, s := range strs {
		select {
		case out <- s:
		case <-done:
			close(out)
			return
		}
	}
	close(out)
}

//*********************************************************

// Requirement 2
func filter(in <-chan string, out chan<- string, done <-chan struct{}) {
	for {
		select {
		case str, ok := <-in:
			if !ok {
				close(out)
				return
			}
			if len(str) > 3 {
				out <- str
			}
		case <-done:
			close(out)
			return
		}
	}
}

//*********************************************************

func main() {
	words := make(chan string)      // Requirement 3a
	results := make(chan string, 7) // Requirement 3a
	done := make(chan struct{})     // Requirement 3b

	go generate(words, done)        // Requirement 3c
	go filter(words, results, done) // Requirement 3c

	for s := range results { // Requirement 3d
		fmt.Printf("Received string from 'results': %s\n", s)
	}

	close(done) // Requirement 3e
}

//*********************************************************
