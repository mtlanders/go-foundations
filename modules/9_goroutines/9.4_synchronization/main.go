package main

import (
	"fmt"
	"sync"
)

//*********************************************************

// Requirement 1
type Stats struct {
	mu    sync.RWMutex
	total int
	count int
	max   int
}

//*********************************************************

// Requirement 1
func (s *Stats) Record(v int) {
	// Using Lock here. The exclusivity model will block
	// ALL access types until the write is completed (once
	// the mutex is acquired). This will make any future
	// reads safe and keeps the write atomic
	s.mu.Lock()
	defer s.mu.Unlock()
	s.total += v
	s.count += 1
	if v > s.max {
		s.max = v
	}
}

//*********************************************************

// Requirement 2
func (s *Stats) Snapshot() (total, count, max int) {
	// Using RLock here. The exclusivity model for RLock
	// allows multiple reads without blocking those other
	// reads. Avoiding Lock because we don't want serial
	// access.
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.total, s.count, s.max
}

//*********************************************************

// Requirement 3
func runWorkers(n int, batchSize int, stats *Stats) {
	var wg sync.WaitGroup
	for i := range n {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range batchSize {
				stats.Record(i*1000 + j)
			}
		}()
	}

	wg.Wait()
}

//*********************************************************

func main() {
	st := Stats{total: 0, count: 0, max: 0}
	runWorkers(5, 100, &st)

	t, c, m := st.Snapshot()
	fmt.Printf("total: %d, count: %d, max: %d\n", t, c, m)
}

//*********************************************************
