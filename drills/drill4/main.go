package main

import "fmt"

const Pi = 3.14159

//*******************************************************************

// Interface - Req 1
type Shape interface {
	Area() float64
}

//*******************************************************************

// Structures - Req 2
type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

//*******************************************************************

// Receivers - Req 3

// Using value receivers - Area() does not mutate the struct internals
// and executes a trivial calculation, so copy overhead is of no concern
func (c Circle) Area() float64 {
	return Pi * (c.Radius * c.Radius)
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

//*******************************************************************

// TotalArea function - Req 4
func TotalArea(shapes []Shape) float64 {
	retval := 0.0
	for _, n := range shapes {
		retval += n.Area()
	}
	return retval
}

//*******************************************************************

func main() {
	rect1 := Rectangle{Width: 10.5, Height: 12.4}
	rect2 := Rectangle{Width: 9.7, Height: 5.2}
	circ1 := Circle{Radius: 4.3}
	circ2 := Circle{Radius: 22.8}

	// Shape slice - Req 5
	ss := []Shape{rect1, rect2, circ1, circ2}
	totalArea := TotalArea(ss)

	fmt.Printf("TotalArea: %0.2f\n", totalArea)
}

//*******************************************************************
