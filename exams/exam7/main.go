package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

//*********************************************************

// Requirement 1
type Job struct {
	ID    int
	Value int
}

// Requirement 1
type Result struct {
	JobID  int
	Output int
}

// Requirement 1
type WorkerStats struct {
	Mtx     sync.RWMutex
	Total   int
	Sum     int
	Highest int
}

// Requirement 1
func (w *WorkerStats) Record(result Result) {
	// Using Lock to block ALL accesses for atomic write
	// to WorkerStats fields. This will ensure that no
	// subsequent read catches the write in an partial state
	// and will make those reads safer
	w.Mtx.Lock()
	defer w.Mtx.Unlock()
	w.Total += 1
	w.Sum += result.Output
	if result.Output > w.Highest {
		w.Highest = result.Output
	}
}

// Requirement 1
func (w *WorkerStats) Snapshot() (total, sum, highest int) {
	// Using RLock for read safety and permissiveness
	w.Mtx.RLock()
	defer w.Mtx.RUnlock()
	return w.Total, w.Sum, w.Highest
}

//*********************************************************

// Requirement 2
func dispatcher(jobCount int, done <-chan struct{}) <-chan Job {
	retChan := make(chan Job)
	go func() {
		for i := range jobCount {
			job := Job{ID: (i + 1), Value: (i + 1) * 2}
			select {
			case retChan <- job:
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

// Requirement 3
func worker(jobs <-chan Job, done <-chan struct{}, stats *WorkerStats) <-chan Result {
	retChan := make(chan Result, 3)
	go func() {
		// Deferring close so that retChan closes when closure function returns
		defer close(retChan)
		for {
			var result Result
			select {
			case j, ok := <-jobs:
				if !ok {
					return
				}
				result = Result{JobID: j.ID, Output: j.Value * j.Value}
				stats.Record(result)
			case _, ok := <-done:
				if !ok {
					return
				}
			}

			select {
			case retChan <- result:
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

// Requirement 4
func fanIn(done <-chan struct{}, resultChans ...<-chan Result) <-chan Result {
	retChan := make(chan Result)
	var wg sync.WaitGroup
	for _, r := range resultChans {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				var result Result
				select {
				case v, ok := <-r:
					if !ok {
						return
					}
					result = v
				case _, ok := <-done:
					if !ok {
						return
					}
				}

				select {
				case retChan <- result:
				case _, ok := <-done:
					if !ok {
						return
					}
				}
			}
		}()
	}

	// Wasn't originally a goroutine, but was causing deadlocks.
	// I split it into a separate goroutine because it costs little,
	// don't affect the wait, and unblocks the critical path back
	// to main
	go func() {
		wg.Wait()
		close(retChan)
	}()

	return retChan
}

//*********************************************************

func main() {
	done := make(chan struct{})
	stat := WorkerStats{Total: 0, Sum: 0, Highest: 0}

	fmt.Printf("start main(): NumGoroutine: %d\n", runtime.NumGoroutine())
	dsp := dispatcher(30, done)

	wrk1 := worker(dsp, done, &stat)
	wrk2 := worker(dsp, done, &stat)
	wrk3 := worker(dsp, done, &stat)
	wrk4 := worker(dsp, done, &stat)

	fan := fanIn(done, wrk1, wrk2, wrk3, wrk4)

	for range fan {
		// Do nothing
	}

	t, s, h := stat.Snapshot()
	fmt.Printf("worker stats: Total: %d, Sum: %d, Highest: %d\n", t, s, h)

	time.Sleep(1 * time.Second)
	fmt.Printf("end main(): NumGoroutine: %d\n", runtime.NumGoroutine())
}

//*********************************************************
