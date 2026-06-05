// Exercise 5 (comprehensive):
// The Rectangle sweep covering area method, scale method with pointer receiver, variadic function
// returning a sub-slice of the last 3 rectangles, and demonstration in main.

package main

import (
	"errors"
	"fmt"
)

// Rectangle struct with fields Width and Height - Requirement
type Rectangle struct {
	Width  int
	Height int
}

// Function which returns Rectangle area as float64 - Criteria 1
func (rect Rectangle) Area() float64 {
	return float64(rect.Height * rect.Width)
}

// Function which multiples dims by scalar with pointer receiver - Criteria 2
// Scalar receiver - Requirement 2
func (rect *Rectangle) scaleRectangle(scalar int) {
	rect.Width *= scalar  // Ptr to struct auto-deref
	rect.Height *= scalar // Ptr to struct auto-deref
}

func appendRectangleSlice(rects ...*Rectangle) ([]*Rectangle, error) {
	if len(rects) < 3 {
		return nil, errors.New("fewer than 3 rectangles provided") // Lower-case error string - Requirement 1
	}

	var retval []*Rectangle // Nil starting point - Requirement 4
	for _, n := range rects {
		retval = append(retval, n)
	}

	return retval[(len(rects) - 3):], nil // Derived from internal slice - Requirement 3
}

func main() {

	rect := Rectangle{5, 10}
	fmt.Printf("Rect area before scale - %f\n", rect.Area()) // Print before in main - Criteria 4

	rectPtr := &rect
	fmt.Printf("Rect dims before scale - Width: %d, Height: %d\n", rectPtr.Width, rectPtr.Height) // Print before in main - Criteria 4
	rectPtr.scaleRectangle(2)
	fmt.Printf("Rect area after scale  - %f\n", rect.Area())                                      // Print after in main - Criteria 4
	fmt.Printf("Rect dims after scale  - Width: %d, Height: %d\n", rectPtr.Width, rectPtr.Height) // Print after in main - Criteria 4

	rectPtr2 := &Rectangle{2, 4}
	rectPtr3 := &Rectangle{3, 7}
	rectPtr4 := &Rectangle{9, 6}

	result, err := appendRectangleSlice(rectPtr, rectPtr2, rectPtr3, rectPtr4)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		for i := 0; i < len(result); i++ { // Intentionally used indexed loop to show Rectangle number
			fmt.Printf("Rectangle %d - Width: %d, Height: %d\n", i, result[i].Width, result[i].Height)
		}
	}
}
