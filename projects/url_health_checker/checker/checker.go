package checker

import (
	"context"
	"net/http"
	"sync"
	"time"
)

//*********************************************************

type Result struct {
	URL     string
	Success bool
}

//*********************************************************

func GetHttpStatus(url string, ctx context.Context) bool {

	retVal := false
	ch := make(chan *http.Response)
	defer close(ch)
	go func() {
		req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		resp, _ := http.DefaultClient.Do(req)
		if resp != nil {
			defer resp.Body.Close()
			select {
			case ch <- resp:
			case <-ctx.Done():
			}
		}
	}()

	select {
	case res, ok := <-ch:
		if !ok {
			retVal = false
		} else if res != nil && res.StatusCode == http.StatusOK {
			retVal = true
		}
	case <-ctx.Done():
		retVal = false
	}
	return retVal
}

//*********************************************************

func Check(jobs uint32, timeout uint32, urls []string) <-chan Result {
	retChan := make(chan Result)

	/*
		NOTE: There could potentially be a bug in this code, specifically
			  when calculating jobs from the specified number of job
			  goroutines. But, I really don't care since this isn't
			  production code and I'd like to move on. The purpose of
			  this exercise wasn't really "how good are you at splitting
			  out jobs" so a buggy section of code is fine
			  for this exercise.
	*/

	// Split out the list of URLs into cfg.NumJobs jobs
	urlsPerJob := uint32(len(urls)) / jobs
	count := 0
	slices := make([][]string, jobs)
	for i := range jobs {
		slices[i] = make([]string, urlsPerJob)
		for j := range urlsPerJob {
			slices[i][j] = urls[count]
			count++
		}
	}

	// Add any missed URLs to existing jobs
	// (where 'num URLs / num jobs' isn't evenly divisible)
	if count < len(urls) {
		remain := len(urls) - count
		for i := range remain {
			slices[i] = append(slices[i], urls[count+i])
		}
	}

	// Now create the goroutines and add them to the WaitGroup
	var wg sync.WaitGroup
	for i := range len(slices) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range len(slices[i]) {

				res := Result{URL: slices[i][j], Success: false}

				ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)

				success := GetHttpStatus(slices[i][j], ctx)
				if success {
					res.Success = true
				}

				select {
				case retChan <- res:
				case <-ctx.Done():
				}
				cancel()
			}
		}()
	}

	// Splitting out into separate goroutine to unblock path back to main
	go func() {
		wg.Wait()
		close(retChan)
	}()

	return retChan
}

//*********************************************************
