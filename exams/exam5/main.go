package main

import (
	"fmt"
	"runtime"
	"time"
)

//*********************************************************

// Requirement 1
func stage1(n int, done <-chan struct{}) <-chan int {
	retChan := make(chan int)
	go func() {
		for ii := range n {
			select {
			case retChan <- (ii + 1):
			case _, ok := <-done:
				if !ok {
					close(retChan)
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
func stage2(in <-chan int, done <-chan struct{}) <-chan int {
	retChan := make(chan int, 4)
	for range 3 {
		go func() {
			for {
				var val int
				select {
				case v, ok := <-in:
					if ok {
						val = v * v
					} else {
						return
					}
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
func stage3(in <-chan int, done <-chan struct{}) <-chan string {
	retChan := make(chan string)
	for range 2 {
		go func() {
			for {
				var str string
				select {
				case v, ok := <-in:
					if ok {
						str = fmt.Sprintf("result: %d", v)
					} else {
						return
					}
				case _, ok := <-done:
					if !ok {
						return
					}
				}

				select {
				case retChan <- str:
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

func main() {

	fmt.Printf("start main(): NumGoroutines: %d\n", runtime.NumGoroutine())

	done := make(chan struct{})
	ch1 := stage1(15, done)
	ch2 := stage2(ch1, done)
	ch3 := stage3(ch2, done)

	count := 0
	for {
		str, ok := <-ch3
		if ok {
			count++
			fmt.Printf("(%d) received: '%s'\n", count, str)
		}

		if count == 8 {
			break
		}
	}

	close(done)

	time.Sleep(1 * time.Second)
	fmt.Printf("end main(): NumGoroutines: %d\n", runtime.NumGoroutine())
}

//*********************************************************
