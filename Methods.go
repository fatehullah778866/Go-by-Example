// package main identifies this file as an executable program
package main

// import "fmt" allows us to print the results to the screen
import "fmt"

// Define a struct named 'rect' to represent a rectangle
// It has two integer fields: width and height
type rect struct {
	width, height int
}

// This is a method with a "pointer receiver" (*rect).
// It calculates the area: width multiplied by height.
// Using a pointer is efficient because it doesn't copy the whole struct.
func (r *rect) area() int {
	return r.width * r.height
}

// This is a method with a "value receiver" (rect).
// It calculates the perimeter: (2 * width) + (2 * height).
// This takes a copy of the struct instead of the original memory address.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	// Create a new rectangle 'r' with a width of 10 and height of 5
	r := rect{width: 10, height: 5}

	// Call the area and perimeter methods on the struct variable 'r'
	// Go automatically handles the conversion between value and pointer here
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Create a pointer 'rp' that points to the rectangle 'r'
	rp := &r

	// Call the methods using the pointer 'rp'
	// Again, Go is smart enough to use the pointer to call both types of methods
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
