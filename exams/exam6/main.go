package main

import (
	"fmt"
	"runtime"
	"time"
)

//*********************************************************

// Requirement 1
func source(id int, count int, done <-chan struct{}) <-chan int {
	retChan := make(chan int)
	go func() {
		for ii := range count {
			select {
			case retChan <- (id*100 + (ii + 1)):
			case _, ok := <-done:
				if !ok {
					return
				}
			}
		}
		close(retChan)
	}()
	return retChan
}

//*********************************************************

// Requirement 2
func merge(done <-chan struct{}, sources ...<-chan int) <-chan int {
	retChan := make(chan int, 10)
	for _, s := range sources {
		go func() {
			for {
				var val int
				select {
				case v, ok := <-s:
					if !ok {
						return
					}
					val = v
				case _, ok := <-done:
					if !ok {
						return
					}
				}

				select {
				case retChan <- val:
				case _, ok := <-done:
					if !ok {
						return
					}
				}
			}
		}()
	}
	return retChan
}

//*********************************************************

// Requirement 3
func conditionalTap(active bool, in <-chan int, done <-chan struct{}) <-chan int {
	var retChan chan int
	if !active {
		retChan = nil
		return retChan
	}

	retChan = make(chan int)
	go func() {
		for {
			var val int
			select {
			case v, ok := <-in:
				if !ok {
					return
				}
				val = v
			case _, ok := <-done:
				if !ok {
					return
				}
			}

			select {
			case retChan <- val:
			case _, ok := <-done:
				if !ok {
					return
				}
			}
		}
	}()
	return retChan
}

//*********************************************************

func main() {
	done := make(chan struct{})

	fmt.Printf("starting source 1: NumGoroutines: %d\n", runtime.NumGoroutine())
	ch1 := source(1, 5, done)

	fmt.Printf("starting source 2: NumGoroutines: %d\n", runtime.NumGoroutine())
	ch2 := source(2, 5, done)

	fmt.Printf("starting source 3: NumGoroutines: %d\n", runtime.NumGoroutine())
	ch3 := source(3, 5, done)

	mergedChan := merge(done, ch1, ch2, ch3)
	tap := conditionalTap(false, mergedChan, done)

	count := 0
	for {
		select {
		case v, ok := <-mergedChan:
			if ok {
				count++
				fmt.Printf("(%d) received from mergedChan: %d\n", count, v)
			}
		case v, ok := <-tap: // Should never fire - bad news if it does
			if ok {
				// Dead code - but disregard
				count++
				fmt.Printf("(%d) received from tap: %d\n", count, v)
			}
		}

		if count == 15 {
			break
		}
	}

	close(done)
	time.Sleep(1 * time.Second)

	fmt.Printf("end main(): NumGoroutines: %d\n", runtime.NumGoroutine())
}

//*********************************************************
