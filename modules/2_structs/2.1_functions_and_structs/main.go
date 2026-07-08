// Exercise 2:
// Write a function that accepts a variadic list of temperature readings, appends them to an internal nil slice,
// and returns a sub-slice of the last 3 readings and an error if fewer than 3 were provided

package main

import (
	"errors"
	"fmt"
)

// Struct with 2 fields - Criteria 1
type Rectangle struct {
	Length int
	Width  int
}

// Area function for Rectangle - Criteria 2
func (r Rectangle) Area() int {
	return r.Length * r.Width
}

// Function that returns largest rectangle and error if none passed in - Criteria 3
func largestArea(rectangles ...Rectangle) (int, int, error) {
	if len(rectangles) == 0 {
		return -1, -1, errors.New("No rectangles passed to function largestArea")
	}

	idx := 0
	area := 0
	temp := 0
	for i, n := range rectangles {
		temp = n.Area()
		if temp > area {
			area = temp
			idx = i
		}
	}
	return idx, area, nil
}

func main() {

	a := Rectangle{Length: 4, Width: 4}
	b := Rectangle{Length: 10, Width: 20}
	c := Rectangle{Length: 2, Width: 5}

	index, result, err := largestArea(a, b, c)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Largest area is rectangle at index %d with value %d", index, result)
	}
}
