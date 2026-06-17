package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//*******************************************************************

type LogStats struct {
	TotalLines  int
	ErrorLines  int
	WarnLines   int
	LongestLine string
}

//*******************************************************************

// A pointer receiver is necessary because the receiver as a value
// would only operate on a copy and thus not mutate the underlying data
func (l *LogStats) ProcessLine(line string) {

	if strings.Contains(line, "ERROR") {
		l.ErrorLines += 1
	}

	if strings.Contains(line, "WARN") {
		l.WarnLines += 1
	}

	l.TotalLines += 1
	if len(line) > len(l.LongestLine) { // Ties go to first seen
		l.LongestLine = line
	}
}

//*******************************************************************

func ReadLog(path string) (*LogStats, error) {

	log := LogStats{TotalLines: 0, ErrorLines: 0, WarnLines: 0, LongestLine: ""}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		log.ProcessLine(line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &log, nil
}

func WriteReport(path string, stats *LogStats) error {

	// Create because it opens if it doesn't exist, truncates if it does
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("writereport: create - %w", err)
	}
	defer f.Close()

	reportLine := "TotalLines: " + strconv.Itoa(stats.TotalLines) + "\n"
	_, err = f.Write([]byte(reportLine))
	if err != nil {
		return fmt.Errorf("writereport: total lines - %w", err)
	}

	reportLine = "ErrorLines: " + strconv.Itoa(stats.ErrorLines) + "\n"
	_, err = f.Write([]byte(reportLine))
	if err != nil {
		return fmt.Errorf("writereport: error lines - %w", err)
	}

	reportLine = "WarnLines: " + strconv.Itoa(stats.WarnLines) + "\n"
	_, err = f.Write([]byte(reportLine))
	if err != nil {
		return fmt.Errorf("writereport: warn lines - %w", err)
	}

	reportLine = "LongestLine: " + stats.LongestLine
	_, err = f.Write([]byte(reportLine))
	if err != nil {
		return fmt.Errorf("writereport: longest line - %w", err)
	}
	return nil
}

//*******************************************************************

func main() {

	log, err := ReadLog("input.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("input file not found: ", err)
		} else {
			fmt.Println("failed to read log: ", err)
		}
		return
	}

	err = WriteReport("output.txt", log)
	if err != nil {
		fmt.Println(err)
	}
}

//*******************************************************************
