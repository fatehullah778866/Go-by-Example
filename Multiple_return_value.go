package main // Defines the package.

import "fmt" // Imports a dependency.

func vals() (int, int) { // Defines the vals function.
	return 3, 7 // Returns a value from the function.
} // Ends the current block.

func main() { // Defines the main function.
	a, b := vals() // Initializes a variable.
	fmt.Println(a) // Writes output to the console.
	fmt.Println(b) // Writes output to the console.

	_, c := vals() // Initializes a variable.
	fmt.Println(c) // Writes output to the console.
} // Ends the current block.
