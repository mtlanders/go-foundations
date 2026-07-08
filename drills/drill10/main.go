package main

import (
	"fmt"
	"runtime"
	"time"
)

//*********************************************************

// Requirement 1
func generator(n int) <-chan int {
	retChan := make(chan int)
	go func() {
		for ii := 1; ii <= n; ii++ {
			retChan <- ii
		}
		close(retChan)
	}()
	return retChan
}

//*********************************************************

// Requirement 2
func worker(jobs <-chan int, done <-chan struct{}) <-chan int {
	retChan := make(chan int, 5)
	go func() {
		for {
			var val int
			select {
			case v, ok := <-jobs:
				if !ok {
					return
				}
				val = v * v
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

	fmt.Printf("Start main(): NumGoroutines: %d\n", runtime.NumGoroutine())

	done := make(chan struct{})

	ch := generator(20)

	fmt.Printf("Starting worker: NumGoroutines: %d\n", runtime.NumGoroutine())
	wch1 := worker(ch, done)

	fmt.Printf("Starting worker: NumGoroutines: %d\n", runtime.NumGoroutine())
	wch2 := worker(ch, done)

	fmt.Printf("Starting worker: NumGoroutines: %d\n", runtime.NumGoroutine())
	wch3 := worker(ch, done)

	fmt.Printf("Starting worker: NumGoroutines: %d\n", runtime.NumGoroutine())
	wch4 := worker(ch, done)

	count := 0
	for {
		select {
		case v, ok := <-wch1:
			if ok {
				count++
				fmt.Printf("(%d) wch1 received value: %d\n", count, v)
			}
		case v, ok := <-wch2:
			if ok {
				count++
				fmt.Printf("(%d) wch2 received value: %d\n", count, v)
			}
		case v, ok := <-wch3:
			if ok {
				count++
				fmt.Printf("(%d) wch3 received value: %d\n", count, v)
			}
		case v, ok := <-wch4:
			if ok {
				count++
				fmt.Printf("(%d) wch4 received value: %d\n", count, v)
			}
		}

		if count == 10 {
			break
		}
	}

	close(done)

	time.Sleep(1 * time.Second)
	fmt.Printf("End main(): NumGoroutines: %d\n", runtime.NumGoroutine())
}

//*********************************************************
