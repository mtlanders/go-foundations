package main

import (
	"errors"
	"fmt"
	"strconv"
)

//*********************************************************

// Requirement 1
var ErrConfigNotFound = errors.New("config not found")
var ErrConfigInvalid = errors.New("config invalid")

//*********************************************************

// Requirement 2
func readRaw(path string) (string, error) {
	if path == "" {
		// Formatting with %w to preserve original sentinel
		return "", fmt.Errorf("readRaw: %w", ErrConfigNotFound)
	} else if path == "empty" {
		return "", nil
	}
	return "100", nil
}

//*********************************************************

// Requirement 3
func parseConfig(raw string) (int, error) {
	if raw == "" {
		// Formatting with %w to preserve original sentinel
		// Dev Note (Req 3): cannot return offending value if raw is empty
		return -1, fmt.Errorf("parseConfig: %w", ErrConfigInvalid)
	}

	// Converting string instead of parsing (req 3 is ambiguous with the 'may')
	val, err := strconv.Atoi(raw)
	if err != nil {
		// Formatting with %w to preserve original error type
		return -1, fmt.Errorf("parseConfig: %w", err)
	}
	return val, nil
}

//*********************************************************

// Requirement 4
func loadTimeout(path string) (int, error) {
	rawStr, err := readRaw(path)
	if err != nil {
		return -1, err
	}

	val, err := parseConfig(rawStr)
	if err != nil {
		return -1, err
	}

	return val, nil
}

//*********************************************************

func detectSentinelType(err error) string {
	if errors.Is(err, ErrConfigNotFound) {
		return "*** Is(): ErrConfigNotFound"
	} else if errors.Is(err, ErrConfigInvalid) {
		return "*** Is(): ErrConfigInvalid"
	}
	return "Unknown error type"
}

//*********************************************************

func main() {
	// Requirement 6 - no capitalized error strings or trailing punctuation (done)
	// Ignoring timeout value since it isn't actually used
	_, err := loadTimeout("/home/users/jsmith/config.txt")
	if err != nil {
		fmt.Println(detectSentinelType(err))
	}

	_, err = loadTimeout("empty")
	if err != nil {
		fmt.Println(detectSentinelType(err))
		fmt.Println(err)
	}

	fmt.Println("")

	_, err = loadTimeout("")
	if err != nil {
		fmt.Println(detectSentinelType(err))
		fmt.Println(err)
	}
}

//*********************************************************
