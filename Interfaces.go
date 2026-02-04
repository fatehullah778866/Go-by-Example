// package main identifies this file as an executable program
package main

import (
	"fmt"  // Used for printing text
	"math" // Used for mathematical constants like Pi
)

// Define an interface named 'geometry'.
// Any type that has these two methods is automatically a 'geometry' type.
type geometry interface {
	area() float64
	perim() float64
}

// Define a struct for a Rectangle
type rect struct {
	width, height float64
}

// Define a struct for a Circle
type circle struct {
	radius float64
}

// Method for Rectangle: calculate area
func (r rect) area() float64 {
	return r.width * r.height
}

// Method for Rectangle: calculate perimeter
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// Method for Circle: calculate area (Pi * r^2)
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Method for Circle: calculate perimeter (2 * Pi * r)
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// This function takes any 'geometry' as an argument.
// It doesn't care if it's a circle or a rectangle; it only cares that it can call area() and perim().
func measure(g geometry) {
	fmt.Println(g)         // Print the shape details
	fmt.Println(g.area())  // Print the calculated area
	fmt.Println(g.perim()) // Print the calculated perimeter
}

// This function uses a "Type Assertion" to check what the shape actually is.
func detectCircle(g geometry) {
	// Check if the geometry 'g' is actually a 'circle' type.
	// 'ok' will be true if it is a circle, and 'false' if it is something else.
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

func main() {
	// Create a rectangle and a circle
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// Both 'r' and 'c' can be used in the measure function
	// because they both follow the 'geometry' rules.
	measure(r)
	measure(c)

	// Check if these shapes are circles
	detectCircle(r) // This will do nothing because 'r' is a rectangle
	detectCircle(c) // This will print the radius
}
