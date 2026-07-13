package main

import (
	"fmt"
	"os"
	"strconv"
	"urlhealth/checker"
	"urlhealth/config"
	"urlhealth/reader"
)

//*********************************************************

func main() {

	// Usage request
	args := os.Args[1:]

	// Now process the command line
	cfg := config.AppConfig{Path: "", NumJobs: 0, Timeout: 0}
	err := cfg.ProcessCmdLine(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Now compile the list of URLs for splitting out into jobs
	urlList := reader.URLList{List: make([]string, 0), Size: 0}
	err = urlList.Read(cfg.Path)
	if err != nil {
		fmt.Println(err)
		return
	}

	results := checker.Check(cfg.NumJobs, cfg.Timeout, urlList.List)

	success := 0
	count := 0
	numWhitespace := 0
	numUrls := len(urlList.List)
	for r := range results {
		count++
		if r.Success {
			success++
		}
		fmt.Printf("[%d/%d] %s", count, numUrls, r.URL)

		urlStr := "[" + strconv.Itoa(count) + "/" + strconv.Itoa(numUrls) + "] " + r.URL
		numWhitespace = 40 - len(urlStr)

		successStr := "Success: " + strconv.FormatBool(r.Success) + "\n"
		fmt.Printf("%*c%s", numWhitespace, ' ', successStr)

	}

	fmt.Printf("*** Results: %d/%d URL checks succeeded (%d failed or could not be processed)", success, numUrls, (numUrls - success))
}

//*********************************************************
