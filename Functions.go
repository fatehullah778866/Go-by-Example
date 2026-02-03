package main // Defines the package.

import "fmt" // Imports a dependency.

func plus(a int, b int) int { // Defines the plus function.
	return a + b // Returns a value from the function.
} // Ends the current block.

func plusPlus(a, b, c int) int { // Defines the plusPlus function.
	return a + b + c // Returns a value from the function.
} // Ends the current block.

func main() { // Defines the main function.
	res := plus(1, 2) // Initializes a variable.
	fmt.Println("1+2 =", res) // Writes output to the console.

	res = plusPlus(1, 2, 3) // Executes a statement.
	fmt.Println("1+2+3 =", res) // Writes output to the console.
} // Ends the current block.
